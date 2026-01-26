package optimizer

import (
	"github.com/arc-language/arc-lang/builder/ir"
)

// DCE implements Aggressive Dead Code Elimination.
// It performs Global DCE (removing unused functions/globals)
// and Local DCE (removing unused instructions within functions).
type DCE struct {
	// Local DCE state
	alive    map[ir.Instruction]bool
	worklist []ir.Instruction

	// Global DCE state
	reachableGlobals   map[*ir.Global]bool
	reachableFunctions map[*ir.Function]bool
	globalWorklist     []ir.Value
}

func NewDCE() *DCE {
	return &DCE{
		alive:              make(map[ir.Instruction]bool),
		worklist:           make([]ir.Instruction, 0),
		reachableGlobals:   make(map[*ir.Global]bool),
		reachableFunctions: make(map[*ir.Function]bool),
		globalWorklist:     make([]ir.Value, 0),
	}
}

// Run executes the optimization pass on the entire module.
func (opt *DCE) Run(module *ir.Module) {
	// 1. Global DCE: Remove unused functions and globals
	opt.runGlobalDCE(module)

	// 2. Local DCE: Remove unused instructions inside the remaining functions
	for _, fn := range module.Functions {
		if len(fn.Blocks) > 0 {
			opt.runLocalDCE(fn)
		}
	}
}

// ============================================================================
// Global DCE (Inter-procedural)
// ============================================================================

func (opt *DCE) runGlobalDCE(m *ir.Module) {
	// Reset State
	opt.reachableGlobals = make(map[*ir.Global]bool)
	opt.reachableFunctions = make(map[*ir.Function]bool)
	opt.globalWorklist = make([]ir.Value, 0)

	// 1. Find Roots (Entry points and Exports)
	for _, fn := range m.Functions {
		// Keep main, _start, and any externally visible function (libraries)
		if fn.Name() == "main" || fn.Name() == "_start" || fn.Linkage == ir.ExternalLinkage {
			opt.markFunctionReachable(fn)
		}
	}
	
	// Keep externally visible globals
	for _, g := range m.Globals {
		if g.Linkage == ir.ExternalLinkage {
			opt.markGlobalReachable(g)
		}
	}

	// 2. Trace Reachability
	for len(opt.globalWorklist) > 0 {
		curr := opt.globalWorklist[len(opt.globalWorklist)-1]
		opt.globalWorklist = opt.globalWorklist[:len(opt.globalWorklist)-1]

		switch v := curr.(type) {
		case *ir.Function:
			opt.scanFunctionDeps(v)
		case *ir.Global:
			opt.scanGlobalDeps(v)
		case ir.Constant:
			opt.scanConstantDeps(v)
		}
	}

	// 3. Prune Module
	var activeFunctions []*ir.Function
	for _, fn := range m.Functions {
		// Keep definitions if reachable.
		// Note: We also keep declarations (functions with no blocks) if they are referenced.
		if opt.reachableFunctions[fn] {
			activeFunctions = append(activeFunctions, fn)
		}
	}
	m.Functions = activeFunctions

	var activeGlobals []*ir.Global
	for _, glob := range m.Globals {
		if opt.reachableGlobals[glob] {
			activeGlobals = append(activeGlobals, glob)
		}
	}
	m.Globals = activeGlobals
}

func (opt *DCE) markFunctionReachable(fn *ir.Function) {
	if fn == nil || opt.reachableFunctions[fn] {
		return
	}
	opt.reachableFunctions[fn] = true
	opt.globalWorklist = append(opt.globalWorklist, fn)
}

func (opt *DCE) markGlobalReachable(g *ir.Global) {
	if g == nil || opt.reachableGlobals[g] {
		return
	}
	opt.reachableGlobals[g] = true
	opt.globalWorklist = append(opt.globalWorklist, g)
}

func (opt *DCE) scanFunctionDeps(fn *ir.Function) {
	if fn == nil {
		return
	}
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			// 1. Scan Operands (handles standard values, args)
			for _, op := range inst.Operands() {
				opt.checkValue(op)
			}

			// 2. Scan Special Fields (Callee fields are not in Operands list)
			// This is CRITICAL for Async/Process instructions to keep their targets alive.
			switch t := inst.(type) {
			case *ir.CallInst:
				if t.Callee != nil {
					opt.markFunctionReachable(t.Callee)
				}
				if t.CalleeVal != nil {
					opt.checkValue(t.CalleeVal)
				}
			case *ir.AsyncTaskCreateInst:
				if t.Callee != nil {
					opt.markFunctionReachable(t.Callee)
				}
				if t.CalleeVal != nil {
					opt.checkValue(t.CalleeVal)
				}
			case *ir.ProcessCreateInst:
				if t.Callee != nil {
					opt.markFunctionReachable(t.Callee)
				}
			case *ir.PhiInst:
				// Phi nodes store values in 'Incoming', not 'Ops'
				for _, inc := range t.Incoming {
					opt.checkValue(inc.Value)
				}
			}
		}
	}
}

func (opt *DCE) scanGlobalDeps(g *ir.Global) {
	if g.Initializer != nil {
		opt.scanConstantDeps(g.Initializer)
	}
}

// scanConstantDeps recursively finds function/global pointers inside complex constants
func (opt *DCE) scanConstantDeps(c ir.Constant) {
	if c == nil { return }
	
	switch t := c.(type) {
	case *ir.Function:
		opt.markFunctionReachable(t)
	case *ir.Global:
		opt.markGlobalReachable(t)
	case *ir.ConstantArray:
		for _, elem := range t.Elements {
			opt.scanConstantDeps(elem)
		}
	case *ir.ConstantStruct:
		for _, field := range t.Fields {
			opt.scanConstantDeps(field)
		}
	}
}

func (opt *DCE) checkValue(v ir.Value) {
	if v == nil { return }
	
	switch t := v.(type) {
	case *ir.Function:
		opt.markFunctionReachable(t)
	case *ir.Global:
		opt.markGlobalReachable(t)
	case ir.Constant:
		opt.scanConstantDeps(t)
	}
}

// ============================================================================
// Local DCE (Intra-procedural)
// ============================================================================

func (opt *DCE) runLocalDCE(fn *ir.Function) {
	// 1. Reset state
	opt.alive = make(map[ir.Instruction]bool)
	opt.worklist = make([]ir.Instruction, 0)

	// 2. Cleanup Control Flow Graph (Remove unreachable blocks)
	opt.removeUnreachableBlocks(fn)

	// 3. Identification Phase: Find "Critical" instructions (Roots)
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if opt.hasSideEffects(inst) {
				opt.markAlive(inst)
			}
		}
	}

	// 4. Propagation Phase
	for len(opt.worklist) > 0 {
		inst := opt.worklist[len(opt.worklist)-1]
		opt.worklist = opt.worklist[:len(opt.worklist)-1]

		// Mark operands as alive
		for _, op := range inst.Operands() {
			if opInst, ok := op.(ir.Instruction); ok {
				opt.markAlive(opInst)
			}
		}

		// Handle Phi nodes specifically (operands are in Incoming structs)
		if phi, ok := inst.(*ir.PhiInst); ok {
			for _, inc := range phi.Incoming {
				if def, ok := inc.Value.(ir.Instruction); ok {
					opt.markAlive(def)
				}
			}
		}
	}

	// 5. Sweep Phase
	for _, block := range fn.Blocks {
		var newInsts []ir.Instruction
		newInsts = make([]ir.Instruction, 0, len(block.Instructions))

		for _, inst := range block.Instructions {
			if opt.alive[inst] {
				newInsts = append(newInsts, inst)
			} else {
				// Dead instruction cleanup: Unregister from tracking to keep Use-Def clean
				for _, op := range inst.Operands() {
					if tracker, ok := op.(ir.TrackableValue); ok {
						tracker.RemoveUser(inst)
					}
				}
			}
		}
		block.Instructions = newInsts
	}
}

func (opt *DCE) markAlive(inst ir.Instruction) {
	if !opt.alive[inst] {
		opt.alive[inst] = true
		opt.worklist = append(opt.worklist, inst)
	}
}

// hasSideEffects determines if an instruction affects state or control flow
func (opt *DCE) hasSideEffects(inst ir.Instruction) bool {
	if inst.IsTerminator() {
		return true
	}

	switch inst.Opcode() {
	case ir.OpStore, ir.OpCall, ir.OpSyscall, ir.OpRaise:
		return true
	case ir.OpVaStart, ir.OpVaEnd:
		return true
	case ir.OpLoad:
		if load, ok := inst.(*ir.LoadInst); ok {
			return load.Volatile
		}
		return false
	// Important: Async/Process creation implies side effects (spawning threads/processes)
	case ir.OpAsyncTaskCreate, ir.OpProcessCreate, ir.OpAsyncTaskAwait:
		return true
	// Memory intrinsics
	case ir.OpMemCpy, ir.OpMemMove, ir.OpMemSet:
		return true
	default:
		return false
	}
}

func (opt *DCE) removeUnreachableBlocks(fn *ir.Function) {
	if fn.EntryBlock() == nil {
		return
	}

	reachable := make(map[*ir.BasicBlock]bool)
	queue := []*ir.BasicBlock{fn.EntryBlock()}
	reachable[fn.EntryBlock()] = true

	// BFS to find reachable blocks
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, succ := range curr.Successors {
			if !reachable[succ] {
				reachable[succ] = true
				queue = append(queue, succ)
			}
		}
	}

	// Filter
	var activeBlocks []*ir.BasicBlock
	for _, block := range fn.Blocks {
		if reachable[block] {
			activeBlocks = append(activeBlocks, block)
			
			// Clean up predecessor lists for reachable blocks
			var validPreds []*ir.BasicBlock
			for _, pred := range block.Predecessors {
				if reachable[pred] {
					validPreds = append(validPreds, pred)
				}
			}
			block.Predecessors = validPreds
		}
	}
	fn.Blocks = activeBlocks
}