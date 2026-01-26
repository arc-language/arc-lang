package optimizer

import (
	"github.com/arc-language/arc-lang/builder/ir"
)

// DCE implements Aggressive Dead Code Elimination.
type DCE struct {
	// Local DCE state
	alive    map[ir.Instruction]bool
	worklist []ir.Instruction

	// Global DCE state
	reachableGlobals   map[*ir.Global]bool
	reachableFunctions map[*ir.Function]bool
	globalWorklist     []ir.Value
	
	// Track functions called by async/process to protect their bodies
	asyncCalledFunctions map[*ir.Function]bool
}

func NewDCE() *DCE {
	return &DCE{
		alive:                make(map[ir.Instruction]bool),
		worklist:             make([]ir.Instruction, 0),
		reachableGlobals:     make(map[*ir.Global]bool),
		reachableFunctions:   make(map[*ir.Function]bool),
		globalWorklist:       make([]ir.Value, 0),
		asyncCalledFunctions: make(map[*ir.Function]bool),
	}
}

func (opt *DCE) Run(module *ir.Module) {
	opt.runGlobalDCE(module)
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
	opt.reachableGlobals = make(map[*ir.Global]bool)
	opt.reachableFunctions = make(map[*ir.Function]bool)
	opt.globalWorklist = make([]ir.Value, 0)
	opt.asyncCalledFunctions = make(map[*ir.Function]bool)

	// 1. Find Roots
	for _, fn := range m.Functions {
		if fn.Name() == "main" || fn.Name() == "_start" || fn.Linkage == ir.ExternalLinkage {
			opt.markFunctionReachable(fn)
		}
	}
	
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
		}
	}

	// 3. Prune Module
	var activeFunctions []*ir.Function
	for _, fn := range m.Functions {
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
			// Extract and track function references
			opt.extractFunctionRefs(inst)
			
			// Scan all operands
			for _, op := range inst.Operands() {
				opt.checkValue(op)
			}

			// Handle Phi nodes (operands in Incoming, not Ops)
			if phi, ok := inst.(*ir.PhiInst); ok {
				for _, inc := range phi.Incoming {
					opt.checkValue(inc.Value)
				}
			}
		}
	}
}

// extractFunctionRefs extracts and marks reachable any function references
func (opt *DCE) extractFunctionRefs(inst ir.Instruction) {
	switch inst.Opcode() {
	case ir.OpCall:
		if call, ok := inst.(*ir.CallInst); ok {
			if call.Callee != nil {
				opt.markFunctionReachable(call.Callee)
			}
			if call.CalleeVal != nil {
				opt.checkValue(call.CalleeVal)
			}
		}
		
	case ir.OpAsyncTaskCreate:
		// CRITICAL: Functions called by async must have their bodies fully preserved
		if async, ok := inst.(*ir.AsyncTaskCreateInst); ok {
			if async.Callee != nil {
				opt.markFunctionReachable(async.Callee)
				opt.asyncCalledFunctions[async.Callee] = true
			}
			if async.CalleeVal != nil {
				opt.checkValue(async.CalleeVal)
			}
		}
		
	case ir.OpProcessCreate:
		// CRITICAL: Functions called by process must have their bodies fully preserved
		if proc, ok := inst.(*ir.ProcessCreateInst); ok {
			if proc.Callee != nil {
				opt.markFunctionReachable(proc.Callee)
				opt.asyncCalledFunctions[proc.Callee] = true
			}
		}
	}
}

func (opt *DCE) scanGlobalDeps(g *ir.Global) {
	if g.Initializer != nil {
		opt.scanConstantDeps(g.Initializer)
	}
}

func (opt *DCE) scanConstantDeps(c ir.Constant) {
	if c == nil {
		return
	}
	
	switch t := c.(type) {
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
	if v == nil {
		return
	}
	
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
	// CRITICAL: Skip local DCE for functions called by async/process operations
	// Their bodies must be preserved exactly as written
	if opt.asyncCalledFunctions[fn] {
		return
	}
	
	opt.alive = make(map[ir.Instruction]bool)
	opt.worklist = make([]ir.Instruction, 0)

	opt.removeUnreachableBlocks(fn)

	// Mark critical instructions
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if opt.isCritical(inst) {
				opt.markAlive(inst)
			}
		}
	}

	// Propagate liveness
	for len(opt.worklist) > 0 {
		inst := opt.worklist[len(opt.worklist)-1]
		opt.worklist = opt.worklist[:len(opt.worklist)-1]

		// Mark operands as alive
		for _, op := range inst.Operands() {
			if opInst, ok := op.(ir.Instruction); ok {
				opt.markAlive(opInst)
			}
		}

		// Handle Phi nodes
		if phi, ok := inst.(*ir.PhiInst); ok {
			for _, inc := range phi.Incoming {
				if def, ok := inc.Value.(ir.Instruction); ok {
					opt.markAlive(def)
				}
			}
		}
	}

	// Sweep dead instructions
	for _, block := range fn.Blocks {
		var newInsts []ir.Instruction
		for _, inst := range block.Instructions {
			if opt.alive[inst] {
				newInsts = append(newInsts, inst)
			} else {
				// Unregister from Use-Def chains
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

// isCritical determines if an instruction must be kept
func (opt *DCE) isCritical(inst ir.Instruction) bool {
	// All terminators are critical
	if inst.IsTerminator() {
		return true
	}

	switch inst.Opcode() {
	case ir.OpStore:
		return true
		
	case ir.OpCall:
		// All calls are critical
		return true
		
	case ir.OpSyscall, ir.OpRaise:
		return true
		
	case ir.OpVaStart, ir.OpVaEnd:
		return true
		
	case ir.OpLoad:
		if load, ok := inst.(*ir.LoadInst); ok {
			return load.Volatile
		}
		return false
		
	// Async/Process operations have side effects
	case ir.OpAsyncTaskCreate, ir.OpAsyncTaskAwait, ir.OpProcessCreate:
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

	var activeBlocks []*ir.BasicBlock
	for _, block := range fn.Blocks {
		if reachable[block] {
			activeBlocks = append(activeBlocks, block)
			
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