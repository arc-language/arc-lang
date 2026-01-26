package optimizer

import (
	"github.com/arc-language/arc-lang/builder/ir"
)

// DCE implements Aggressive Dead Code Elimination.
// It performs both Global DCE (removing unused functions/globals)
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
	// 1. Remove unused functions and globals first
	opt.runGlobalDCE(module)

	// 2. Optimize the bodies of the remaining functions
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
	opt.globalWorklist = opt.globalWorklist[:0]

	// 1. Find Roots (Entry points)
	for _, fn := range m.Functions {
		if fn.Name() == "main" || fn.Name() == "_start" {
			opt.markFunctionReachable(fn)
		}
	}

	// 2. Trace Reachability
	for len(opt.globalWorklist) > 0 {
		curr := opt.globalWorklist[len(opt.globalWorklist)-1]
		opt.globalWorklist = opt.globalWorklist[:len(opt.globalWorklist)-1]

		if fn, ok := curr.(*ir.Function); ok {
			opt.scanFunctionDeps(fn)
		} else if glob, ok := curr.(*ir.Global); ok {
			opt.scanGlobalDeps(glob)
		}
	}

	// 3. Prune Module
	var activeFunctions []*ir.Function
	for _, fn := range m.Functions {
		if len(fn.Blocks) == 0 {
			// Keep declarations if they are reachable
			if opt.reachableFunctions[fn] {
				activeFunctions = append(activeFunctions, fn)
			}
		} else {
			// Keep definitions if they are reachable
			if opt.reachableFunctions[fn] {
				activeFunctions = append(activeFunctions, fn)
			}
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
	// FIX: Guard against nil pointers (typed nils from interfaces)
	if fn == nil {
		return
	}
	if !opt.reachableFunctions[fn] {
		opt.reachableFunctions[fn] = true
		opt.globalWorklist = append(opt.globalWorklist, fn)
	}
}

func (opt *DCE) markGlobalReachable(g *ir.Global) {
	// FIX: Guard against nil pointers
	if g == nil {
		return
	}
	if !opt.reachableGlobals[g] {
		opt.reachableGlobals[g] = true
		opt.globalWorklist = append(opt.globalWorklist, g)
	}
}

func (opt *DCE) scanFunctionDeps(fn *ir.Function) {
	if fn == nil {
		return
	}
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			// Scan Operands (handles Loaded Globals, Function pointers passed as args)
			for _, op := range inst.Operands() {
				opt.checkValue(op)
			}

			// Scan Special Fields (CallInst Targets)
			if call, ok := inst.(*ir.CallInst); ok {
				if call.Callee != nil {
					opt.markFunctionReachable(call.Callee)
				}
				// Also check CalleeVal if it's a value (could be global/function pointer)
				if call.CalleeVal != nil {
					opt.checkValue(call.CalleeVal)
				}
			}
			
			// FIX: Trace Async Task Targets
			if async, ok := inst.(*ir.AsyncTaskCreateInst); ok {
				if async.Callee != nil {
					opt.markFunctionReachable(async.Callee)
				}
				if async.CalleeVal != nil {
					opt.checkValue(async.CalleeVal)
				}
			}
			
			// FIX: Trace Process Targets (This fixes the segfault)
			if proc, ok := inst.(*ir.ProcessCreateInst); ok {
				if proc.Callee != nil {
					opt.markFunctionReachable(proc.Callee)
				}
			}
		}
	}
}

func (opt *DCE) scanGlobalDeps(g *ir.Global) {
	if g.Initializer != nil {
		opt.checkValue(g.Initializer)
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
	case *ir.ConstantStruct:
		for _, f := range t.Fields {
			opt.checkValue(f)
		}
	case *ir.ConstantArray:
		for _, e := range t.Elements {
			opt.checkValue(e)
		}
	}
}

// ============================================================================
// Local DCE (Intra-procedural)
// ============================================================================

func (opt *DCE) runLocalDCE(fn *ir.Function) {
	// 1. Reset state
	opt.alive = make(map[ir.Instruction]bool)
	opt.worklist = opt.worklist[:0]

	// 2. Cleanup Control Flow Graph
	opt.removeUnreachableBlocks(fn)

	// 3. Identification Phase: Find "Critical" instructions
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

		for _, op := range inst.Operands() {
			if opInst, ok := op.(ir.Instruction); ok {
				opt.markAlive(opInst)
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
				// Dead instruction cleanup
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

// markAlive marks an instruction as live and adds it to the worklist.
func (opt *DCE) markAlive(inst ir.Instruction) {
	if !opt.alive[inst] {
		opt.alive[inst] = true
		opt.worklist = append(opt.worklist, inst)
	}
}

// hasSideEffects determines if an instruction MUST be kept.
func (opt *DCE) hasSideEffects(inst ir.Instruction) bool {
	switch inst.Opcode() {
	case ir.OpRet, ir.OpBr, ir.OpCondBr, ir.OpSwitch, ir.OpUnreachable:
		return true
	case ir.OpStore:
		return true
	case ir.OpCall, ir.OpSyscall, ir.OpRaise:
		return true
	case ir.OpLoad:
		if load, ok := inst.(*ir.LoadInst); ok && load.Volatile {
			return true
		}
		return false
	// FIX: Added Async/Process ops to side-effects.
	// Even if the return value (Handle/PID) is unused, the task/process must still start.
	// Await creates a synchronization point, so it must also be kept.
	case ir.OpAsyncTaskCreate, ir.OpProcessCreate, ir.OpAsyncTaskAwait:
		return true
	default:
		return false
	}
}

// removeUnreachableBlocks performs simple CFG cleanup.
func (opt *DCE) removeUnreachableBlocks(fn *ir.Function) {
	if fn.EntryBlock() == nil {
		return
	}

	reachable := make(map[*ir.BasicBlock]bool)
	queue := []*ir.BasicBlock{fn.EntryBlock()}
	reachable[fn.EntryBlock()] = true

	// BFS
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
			var validPreds []*ir.BasicBlock
			for _, pred := range block.Predecessors {
				if reachable[pred] {
					validPreds = append(validPreds, pred)
				}
			}
			block.Predecessors = validPreds
		} else {
			for _, succ := range block.Successors {
				opt.removePredecessor(succ, block)
			}
		}
	}
	fn.Blocks = activeBlocks
}

func (opt *DCE) removePredecessor(target *ir.BasicBlock, deadPred *ir.BasicBlock) {
	var newPreds []*ir.BasicBlock
	for _, pred := range target.Predecessors {
		if pred != deadPred {
			newPreds = append(newPreds, pred)
		}
	}
	target.Predecessors = newPreds
}