package optimizer

import (
	"github.com/arc-language/arc-lang/builder/ir"
)

// DCE implements Aggressive Dead Code Elimination.
// It assumes all instructions are dead until proven alive by side effects or dependencies.
type DCE struct {
	alive    map[ir.Instruction]bool
	worklist []ir.Instruction
}

func NewDCE() *DCE {
	return &DCE{
		alive:    make(map[ir.Instruction]bool),
		worklist: make([]ir.Instruction, 0),
	}
}

// Run executes the optimization pass on the entire module.
func (opt *DCE) Run(module *ir.Module) {
	for _, fn := range module.Functions {
		// Only optimize functions with bodies
		if len(fn.Blocks) > 0 {
			opt.optimizeFunction(fn)
		}
	}
}

func (opt *DCE) optimizeFunction(fn *ir.Function) {
	// 1. Reset state
	opt.alive = make(map[ir.Instruction]bool)
	opt.worklist = opt.worklist[:0]

	// 2. Cleanup Control Flow Graph (Remove unreachable blocks first)
	// This helps DCE by removing entire dead branches so we don't process them.
	opt.removeUnreachableBlocks(fn)

	// 3. Identification Phase: Find "Critical" instructions
	// These are the roots of our liveness graph.
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if opt.hasSideEffects(inst) {
				opt.markAlive(inst)
			}
		}
	}

	// 4. Propagation Phase: Follow Use-Def chains
	// If 'Inst' is alive, then its Operands (definitions) must also be alive.
	for len(opt.worklist) > 0 {
		inst := opt.worklist[len(opt.worklist)-1]
		opt.worklist = opt.worklist[:len(opt.worklist)-1]

		for _, op := range inst.Operands() {
			// If the operand is an instruction, mark it alive
			if opInst, ok := op.(ir.Instruction); ok {
				opt.markAlive(opInst)
			}
		}
	}

	// 5. Sweep Phase: Remove dead instructions
	for _, block := range fn.Blocks {
		var newInsts []ir.Instruction
		// Pre-allocate to avoid resize overhead
		newInsts = make([]ir.Instruction, 0, len(block.Instructions))

		for _, inst := range block.Instructions {
			if opt.alive[inst] {
				newInsts = append(newInsts, inst)
			} else {
				// Instruction is dead.
				// Unregister it from its operands' Users lists to keep the graph clean.
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
	// Control flow is critical
	case ir.OpRet, ir.OpBr, ir.OpCondBr, ir.OpSwitch, ir.OpUnreachable:
		return true

	// Writing to memory is critical
	// (Unless we later implement Dead Store Elimination, which assumes stack slots are local)
	case ir.OpStore:
		return true

	// Function calls are assumed to have side effects (I/O, globals)
	// (A smarter pass could check 'readnone' or 'readonly' attributes)
	case ir.OpCall, ir.OpSyscall, ir.OpRaise:
		return true

	// Volatile loads interact with hardware/concurrency
	case ir.OpLoad:
		if load, ok := inst.(*ir.LoadInst); ok && load.Volatile {
			return true
		}
		return false

	// Threading/Process primitives are side effects
	case ir.OpAsyncTaskCreate, ir.OpProcessCreate:
		return true

	// Everything else (Math, Casts, Allocas, GEPs) is dead if unused.
	// Note: Removing 'alloca' is safe because if no one stores/loads from it,
	// the stack slot is irrelevant.
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

	// BFS to find all reachable blocks
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

	// Filter the block list
	var activeBlocks []*ir.BasicBlock
	for _, block := range fn.Blocks {
		if reachable[block] {
			activeBlocks = append(activeBlocks, block)

			// Clean up predecessor lists of the blocks we keep
			// Remove dead predecessors that no longer exist
			var validPreds []*ir.BasicBlock
			for _, pred := range block.Predecessors {
				if reachable[pred] {
					validPreds = append(validPreds, pred)
				}
			}
			block.Predecessors = validPreds

		} else {
			// Block is dead. Remove it from its successors' Predecessors list.
			for _, succ := range block.Successors {
				opt.removePredecessor(succ, block)
			}
		}
	}
	fn.Blocks = activeBlocks
}

// removePredecessor removes a specific block from a target's predecessor list.
func (opt *DCE) removePredecessor(target *ir.BasicBlock, deadPred *ir.BasicBlock) {
	var newPreds []*ir.BasicBlock
	for _, pred := range target.Predecessors {
		if pred != deadPred {
			newPreds = append(newPreds, pred)
		}
	}
	target.Predecessors = newPreds
}