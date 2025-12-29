// --- START OF FILE codegen/arch/amd64/controlflow.go ---
package amd64

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

// Return instruction
func (c *compiler) retOp(inst *ir.RetInst) error {
	if inst.NumOperands() > 0 && inst.Operands()[0] != nil {
		retVal := inst.Operands()[0]

		// Check if it's a struct return
		if structType, ok := retVal.Type().(*types.StructType); ok {
			structSize := SizeOf(structType)
			
			if structSize <= 8 {
				// Small struct (≤8 bytes) - return in RAX
				c.loadToReg(RAX, retVal)
			} else if structSize <= 16 {
				// Medium struct (9-16 bytes) - return in RAX:RDX
				// Load the struct and split into RAX (low 8) and RDX (high 8)
				// For now, just load into RAX
				c.loadToReg(RAX, retVal)
			} else {
				// Large struct (>16 bytes) - should use hidden pointer parameter
				// This should be handled differently at the IR level
				c.loadToReg(RAX, retVal)
			}
		} else if types.IsFloat(retVal.Type()) {
			c.loadToFpReg(0, retVal) // Return in XMM0
		} else {
			c.loadToReg(RAX, retVal) // Return in RAX
		}
	}

	// Epilogue
	// leave (equivalent to: mov rsp, rbp; pop rbp)
	c.emitBytes(0xC9)
	// ret
	c.emitBytes(0xC3)

	return nil
}

// Unconditional branch
func (c *compiler) brOp(inst *ir.BrInst) error {
	// Handle phi nodes in target block before branching
	c.handlePhiForBranch(inst.Parent(), inst.Target)
	
	// jmp rel32
	c.emitBytes(0xE9)
	c.fixups = append(c.fixups, jumpFixup{
		offset: c.text.Len(),
		target: inst.Target,
	})
	c.emitUint32(0) // Placeholder

	return nil
}

// Conditional branch
func (c *compiler) condBrOp(inst *ir.CondBrInst) error {
	c.loadToReg(RAX, inst.Condition)

	// test rax, rax
	c.emitBytes(0x48, 0x85, 0xC0)

	// jz false_block (jump to false block if zero)
	c.emitBytes(0x0F, 0x84)
	c.fixups = append(c.fixups, jumpFixup{
		offset: c.text.Len(),
		target: inst.FalseBlock,
	})
	c.emitUint32(0) // Placeholder

	// True path falls through - handle phi and jump to true block
	c.handlePhiForBranch(inst.Parent(), inst.TrueBlock)
	c.emitBytes(0xE9)
	c.fixups = append(c.fixups, jumpFixup{
		offset: c.text.Len(),
		target: inst.TrueBlock,
	})
	c.emitUint32(0)

	// Note: No false path handling here - the jz above jumps directly to FalseBlock
	// If FalseBlock has phi nodes, they should be handled at the start of that block

	return nil
}

// Switch instruction
func (c *compiler) switchOp(inst *ir.SwitchInst) error {
	c.loadToReg(RAX, inst.Condition)

	// Generate comparison chain
	for _, switchCase := range inst.Cases {
		// cmp rax, case_value
		if switchCase.Value.Value >= -128 && switchCase.Value.Value <= 127 {
			c.emitBytes(0x48, 0x83, 0xF8, byte(switchCase.Value.Value))
		} else {
			c.emitBytes(0x48, 0x3D)
			c.emitInt32(int32(switchCase.Value.Value))
		}

		// je case_block
		c.emitBytes(0x0F, 0x84)
		c.fixups = append(c.fixups, jumpFixup{
			offset: c.text.Len(),
			target: switchCase.Block,
		})
		c.emitUint32(0)
	}

	// Jump to default block
	c.handlePhiForBranch(inst.Parent(), inst.DefaultBlock)
	c.emitBytes(0xE9)
	c.fixups = append(c.fixups, jumpFixup{
		offset: c.text.Len(),
		target: inst.DefaultBlock,
	})
	c.emitUint32(0)

	return nil
}

// Helper function to handle phi nodes before branching
func (c *compiler) handlePhiForBranch(fromBlock, toBlock *ir.BasicBlock) {
	// Find all phi nodes in the target block
	for _, inst := range toBlock.Instructions {
		phi, ok := inst.(*ir.PhiInst)
		if !ok {
			break // Phi nodes are always at the start of a block
		}
		
		// Find the incoming value from fromBlock
		for _, incoming := range phi.Incoming {
			if incoming.Block == fromBlock {
				// Copy the value to phi's location
				c.loadToReg(RAX, incoming.Value)
				c.storeFromReg(RAX, phi)
				break
			}
		}
	}
}

// Phi node - now properly handled before branches
func (c *compiler) phiOp(inst *ir.PhiInst) error {
	// Phi nodes are handled by the branch instructions
	// The value is already in place when we reach this instruction
	// So we don't need to do anything here
	return nil
}

// Select (ternary operator)
func (c *compiler) selectOp(inst *ir.SelectInst) error {
	ops := inst.Operands()
	cond := ops[0]
	trueVal := ops[1]
	falseVal := ops[2]

	c.loadToReg(RAX, cond)
	c.loadToReg(RCX, trueVal)
	c.loadToReg(RDX, falseVal)

	// test rax, rax
	c.emitBytes(0x48, 0x85, 0xC0)

	// cmovz rcx, rdx (move rdx to rcx if zero)
	c.emitBytes(0x48, 0x0F, 0x44, 0xCA)

	// Result in RCX
	c.storeFromReg(RCX, inst)
	return nil
}

// Function call
func (c *compiler) callOp(inst *ir.CallInst) error {
	ops := inst.Operands()

	// System V AMD64 ABI calling convention
	// Integer/pointer args: RDI, RSI, RDX, RCX, R8, R9, then stack
	// Float args: XMM0-XMM7, then stack
	// Return: RAX (integer), XMM0 (float), RAX:RDX (large struct up to 16 bytes)

	intArgRegs := []int{RDI, RSI, RDX, RCX, R8, R9}
	fpArgRegs := []int{0, 1, 2, 3, 4, 5, 6, 7} // XMM0-XMM7

	intArgIdx := 0
	fpArgIdx := 0
	stackArgs := []ir.Value{}

	// Classify and place arguments
	for _, arg := range ops {
		if types.IsFloat(arg.Type()) {
			if fpArgIdx < len(fpArgRegs) {
				c.loadToFpReg(fpArgRegs[fpArgIdx], arg)
				fpArgIdx++
			} else {
				stackArgs = append(stackArgs, arg)
			}
		} else {
			if intArgIdx < len(intArgRegs) {
				c.loadToReg(intArgRegs[intArgIdx], arg)
				intArgIdx++
			} else {
				stackArgs = append(stackArgs, arg)
			}
		}
	}

	// Push stack arguments in reverse order
	for i := len(stackArgs) - 1; i >= 0; i-- {
		c.loadToReg(RAX, stackArgs[i])
		// push rax
		c.emitBytes(0x50)
	}

	// Align stack to 16 bytes if needed (ABI requirement)
	stackAdjust := len(stackArgs) * 8
	if stackAdjust%16 != 0 {
		// sub rsp, 8
		c.emitBytes(0x48, 0x83, 0xEC, 0x08)
		stackAdjust += 8
	}

	// Emit call
	calleeName := inst.CalleeName
	if inst.Callee != nil {
		calleeName = inst.Callee.Name()
	}

	// call rel32
	c.emitBytes(0xE8)

	// Add relocation for the call
	c.relocations = append(c.relocations, Relocation{
		Offset:     uint64(c.text.Len()),
		SymbolName: calleeName,
		Type:       R_X86_64_PLT32,
		Addend:     -4,
	})
	c.emitUint32(0) // Placeholder

	// Clean up stack
	if stackAdjust > 0 {
		if stackAdjust <= 127 {
			c.emitBytes(0x48, 0x83, 0xC4, byte(stackAdjust))
		} else {
			c.emitBytes(0x48, 0x81, 0xC4)
			c.emitUint32(uint32(stackAdjust))
		}
	}

	// Store return value
	if inst.Type() != nil && inst.Type().Kind() != types.VoidKind {
		// Check if return type is a small struct (returned in registers)
		if structType, ok := inst.Type().(*types.StructType); ok {
			structSize := SizeOf(structType)
			
			if structSize <= 8 {
				// Small struct (≤8 bytes) returned directly in RAX
				// Store RAX as-is - it contains the struct value, not a pointer
				c.storeFromReg(RAX, inst)
			} else if structSize <= 16 {
				// Medium struct (9-16 bytes) returned in RAX (low 8) and RDX (high 8)
				// For now, just store RAX portion
				// TODO: Properly handle the RDX portion for 9-16 byte structs
				c.storeFromReg(RAX, inst)
			} else {
				// Large struct (>16 bytes) - caller allocates space and passes hidden pointer
				// The function returns the pointer in RAX
				c.storeFromReg(RAX, inst)
			}
		} else if types.IsFloat(inst.Type()) {
			c.storeFromFpReg(0, inst)
		} else {
			c.storeFromReg(RAX, inst)
		}
	}

	return nil
}