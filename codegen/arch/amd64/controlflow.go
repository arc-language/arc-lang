package amd64

import (
	"fmt"

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
				// TODO: Properly handle 9-16 byte structs
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

// Extract value from aggregate
func (c *compiler) extractValueOp(inst *ir.ExtractValueInst) error {
	agg := inst.Operands()[0]
	c.loadToReg(RAX, agg)

	// Calculate offset based on indices
	currentType := agg.Type()
	offset := 0

	for _, idx := range inst.Indices {
		switch ty := currentType.(type) {
		case *types.StructType:
			offset += GetStructFieldOffset(ty, idx)
			currentType = ty.Fields[idx]
		case *types.ArrayType:
			elemSize := SizeOf(ty.ElementType)
			offset += idx * elemSize
			currentType = ty.ElementType
		default:
			return fmt.Errorf("extractvalue on non-aggregate type: %T", ty)
		}
	}

	// Load from aggregate + offset
	if offset > 0 {
		if offset <= 127 {
			c.emitBytes(0x48, 0x83, 0xC0, byte(offset))
		} else {
			c.emitBytes(0x48, 0x05)
			c.emitInt32(int32(offset))
		}
	}

	// Load the value
	size := SizeOf(inst.Type())
	switch size {
	case 1:
		c.emitBytes(0x48, 0x0F, 0xB6, 0x00) // movzx rax, byte ptr [rax]
	case 2:
		c.emitBytes(0x48, 0x0F, 0xB7, 0x00) // movzx rax, word ptr [rax]
	case 4:
		c.emitBytes(0x8B, 0x00) // mov eax, [rax]
	case 8:
		c.emitBytes(0x48, 0x8B, 0x00) // mov rax, [rax]
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Insert value into aggregate
func (c *compiler) insertValueOp(inst *ir.InsertValueInst) error {
	ops := inst.Operands()
	agg := ops[0]
	value := ops[1]

	// Calculate the size of the aggregate
	aggSize := SizeOf(inst.Type())
	
	// Allocate temporary space for the result struct on the stack
	// We need to store it at the instruction's stack location
	destOffset, ok := c.stackMap[inst]
	if !ok {
		return fmt.Errorf("no stack location for insertvalue result")
	}

	// Step 1: Copy the source aggregate to the destination
	// Handle ConstantZero specially
	if _, isZero := agg.(*ir.ConstantZero); isZero {
		// Zero out the destination
		// lea rax, [rbp + destOffset]
		c.emitBytes(0x48, 0x8D, 0x85)
		c.emitInt32(int32(destOffset))
		
		// xor ecx, ecx
		c.emitXorReg(RCX, RCX)
		
		// mov rcx, aggSize/8 (number of qwords)
		c.loadConstInt(RCX, int64((aggSize+7)/8))
		
		// Move RAX to RDI for stosq
		c.emitBytes(0x48, 0x89, 0xC7)
		
		// xor rax, rax (value to store)
		c.emitXorReg(RAX, RAX)
		
		// rep stosq - fill with zeros
		c.emitBytes(0xF3, 0x48, 0xAB)
	} else {
		// Copy from source aggregate location
		srcOffset, ok := c.stackMap[agg]
		if !ok {
			// Try to handle it as a constant or other value
			c.loadToReg(RAX, agg)
			c.emitStoreToStack(RAX, destOffset, aggSize)
		} else {
			// Copy struct from srcOffset to destOffset
			if aggSize <= 8 {
				// Small struct - simple copy
				c.emitLoadFromStack(RAX, srcOffset, aggSize)
				c.emitStoreToStack(RAX, destOffset, aggSize)
			} else {
				// Larger struct - copy multiple words
				for off := 0; off < aggSize; off += 8 {
					size := 8
					if off+8 > aggSize {
						size = aggSize - off
					}
					c.emitLoadFromStack(RAX, srcOffset+off, size)
					c.emitStoreToStack(RAX, destOffset+off, size)
				}
			}
		}
	}

	// Step 2: Calculate the offset for the field to modify
	currentType := agg.Type()
	offset := 0

	for _, idx := range inst.Indices {
		switch ty := currentType.(type) {
		case *types.StructType:
			offset += GetStructFieldOffset(ty, idx)
			currentType = ty.Fields[idx]
		case *types.ArrayType:
			elemSize := SizeOf(ty.ElementType)
			offset += idx * elemSize
			currentType = ty.ElementType
		}
	}

	// Step 3: Store the new value at destOffset + offset
	c.loadToReg(RAX, value)
	
	// lea rcx, [rbp + destOffset + offset]
	c.emitBytes(0x48, 0x8D, 0x8D)
	c.emitInt32(int32(destOffset + offset))
	
	// Store value to [rcx]
	size := SizeOf(value.Type())
	switch size {
	case 1:
		c.emitBytes(0x88, 0x01) // mov byte ptr [rcx], al
	case 2:
		c.emitBytes(0x66, 0x89, 0x01) // mov word ptr [rcx], ax
	case 4:
		c.emitBytes(0x89, 0x01) // mov dword ptr [rcx], eax
	case 8:
		c.emitBytes(0x48, 0x89, 0x01) // mov qword ptr [rcx], rax
	}

	// The result is already in place at destOffset
	// No need to call storeFromReg since the value is already where it should be
	return nil
}

// Integer cast operations
func (c *compiler) intCastOp(inst *ir.CastInst) error {
	src := inst.Operands()[0]
	c.loadToReg(RAX, src)

	srcSize := SizeOf(src.Type())

	switch inst.Opcode() {
	case ir.OpTrunc:
		// Truncation - just take lower bits (already in RAX)
		// No operation needed, storing will handle it

	case ir.OpZExt:
		// Zero extension
		switch srcSize {
		case 1:
			c.emitBytes(0x48, 0x0F, 0xB6, 0xC0) // movzx rax, al
		case 2:
			c.emitBytes(0x48, 0x0F, 0xB7, 0xC0) // movzx rax, ax
		case 4:
			c.emitBytes(0x89, 0xC0) // mov eax, eax (zero-extends)
		}

	case ir.OpSExt:
		// Sign extension
		switch srcSize {
		case 1:
			c.emitBytes(0x48, 0x0F, 0xBE, 0xC0) // movsx rax, al
		case 2:
			c.emitBytes(0x48, 0x0F, 0xBF, 0xC0) // movsx rax, ax
		case 4:
			c.emitBytes(0x48, 0x63, 0xC0) // movsxd rax, eax
		}
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Floating point cast operations
func (c *compiler) fpCastOp(inst *ir.CastInst) error {
	src := inst.Operands()[0]
	srcType := src.Type().(*types.FloatType)
	dstType := inst.Type().(*types.FloatType)

	c.loadToFpReg(0, src)

	if srcType.BitWidth == 32 && dstType.BitWidth == 64 {
		// cvtss2sd xmm0, xmm0
		c.emitBytes(0xF3, 0x0F, 0x5A, 0xC0)
	} else if srcType.BitWidth == 64 && dstType.BitWidth == 32 {
		// cvtsd2ss xmm0, xmm0
		c.emitBytes(0xF2, 0x0F, 0x5A, 0xC0)
	}

	c.storeFromFpReg(0, inst)
	return nil
}

// Float to integer conversion
func (c *compiler) fpToIntOp(inst *ir.CastInst) error {
	src := inst.Operands()[0]
	srcType := src.Type().(*types.FloatType)

	c.loadToFpReg(0, src)

	if srcType.BitWidth == 32 {
		// cvttss2si rax, xmm0
		c.emitBytes(0xF3, 0x48, 0x0F, 0x2C, 0xC0)
	} else {
		// cvttsd2si rax, xmm0
		c.emitBytes(0xF2, 0x48, 0x0F, 0x2C, 0xC0)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Integer to float conversion
func (c *compiler) intToFpOp(inst *ir.CastInst) error {
	src := inst.Operands()[0]
	dstType := inst.Type().(*types.FloatType)

	c.loadToReg(RAX, src)

	if dstType.BitWidth == 32 {
		// cvtsi2ss xmm0, rax
		c.emitBytes(0xF3, 0x48, 0x0F, 0x2A, 0xC0)
	} else {
		// cvtsi2sd xmm0, rax
		c.emitBytes(0xF2, 0x48, 0x0F, 0x2A, 0xC0)
	}

	c.storeFromFpReg(0, inst)
	return nil
}

// Bitcast and pointer casts
func (c *compiler) bitcastOp(inst *ir.CastInst) error {
	src := inst.Operands()[0]

	// For bitcast, just copy the bits
	// For pointer/int conversions, also just copy
	c.loadToReg(RAX, src)
	c.storeFromReg(RAX, inst)

	return nil
}
