// --- START OF FILE codegen/arch/amd64/ops.go ---
package amd64

import (
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (c *compiler) compileInstruction(inst ir.Instruction) error {
	switch inst.Opcode() {
	// Arithmetic
	case ir.OpAdd:
		return c.addOp(inst)
	case ir.OpSub:
		return c.subOp(inst)
	case ir.OpMul:
		return c.mulOp(inst)
	case ir.OpUDiv, ir.OpSDiv:
		return c.divOp(inst, false)
	case ir.OpURem, ir.OpSRem:
		return c.divOp(inst, true)

	// Floating point
	case ir.OpFAdd:
		return c.fpBinOp(inst, 0x58)
	case ir.OpFSub:
		return c.fpBinOp(inst, 0x5C)
	case ir.OpFMul:
		return c.fpBinOp(inst, 0x59)
	case ir.OpFDiv:
		return c.fpBinOp(inst, 0x5E)

	// Bitwise
	case ir.OpAnd:
		return c.andOp(inst)
	case ir.OpOr:
		return c.orOp(inst)
	case ir.OpXor:
		return c.xorOp(inst)
	case ir.OpShl:
		return c.shiftOp(inst, 0x00) // shl uses /4 -> 0xE0
	case ir.OpLShr:
		return c.shiftOp(inst, 0x08) // shr uses /5 -> 0xE8
	case ir.OpAShr:
		return c.shiftOp(inst, 0x18) // sar uses /7 -> 0xF8

	// Memory
	case ir.OpAlloca:
		return c.allocaOp(inst.(*ir.AllocaInst))
	case ir.OpLoad:
		return c.loadOp(inst.(*ir.LoadInst))
	case ir.OpStore:
		return c.storeOp(inst.(*ir.StoreInst))
	case ir.OpGetElementPtr:
		return c.gepOp(inst.(*ir.GetElementPtrInst))

	// Comparison
	case ir.OpICmp:
		return c.icmpOp(inst.(*ir.ICmpInst))
	case ir.OpFCmp:
		return c.fcmpOp(inst.(*ir.FCmpInst))

	// Control flow
	case ir.OpRet:
		return c.retOp(inst.(*ir.RetInst))
	case ir.OpBr:
		return c.brOp(inst.(*ir.BrInst))
	case ir.OpCondBr:
		return c.condBrOp(inst.(*ir.CondBrInst))
	case ir.OpSwitch:
		return c.switchOp(inst.(*ir.SwitchInst))

	// Casts
	case ir.OpTrunc, ir.OpZExt, ir.OpSExt:
		return c.intCastOp(inst.(*ir.CastInst))
	case ir.OpFPTrunc, ir.OpFPExt:
		return c.fpCastOp(inst.(*ir.CastInst))
	case ir.OpFPToUI, ir.OpFPToSI:
		return c.fpToIntOp(inst.(*ir.CastInst))
	case ir.OpUIToFP, ir.OpSIToFP:
		return c.intToFpOp(inst.(*ir.CastInst))
	case ir.OpPtrToInt, ir.OpIntToPtr, ir.OpBitcast:
		return c.bitcastOp(inst.(*ir.CastInst))

	// Other
	case ir.OpPhi:
		return c.phiOp(inst.(*ir.PhiInst))
	case ir.OpSelect:
		return c.selectOp(inst.(*ir.SelectInst))
	case ir.OpCall:
		return c.callOp(inst.(*ir.CallInst))
	case ir.OpSyscall:
		return c.syscallOp(inst.(*ir.SyscallInst))
	case ir.OpExtractValue:
		return c.extractValueOp(inst.(*ir.ExtractValueInst))
	case ir.OpInsertValue:
		return c.insertValueOp(inst.(*ir.InsertValueInst))

	// Variadic operations
	case ir.OpVaStart:
		return c.vaStartOp(inst.(*ir.VaStartInst))
	case ir.OpVaArg:
		return c.vaArgOp(inst.(*ir.VaArgInst))
	case ir.OpVaEnd:
		return c.vaEndOp(inst.(*ir.VaEndInst))

	// Intrinsics
	case ir.OpSizeOf:
		return c.sizeOfOp(inst.(*ir.SizeOfInst))
	case ir.OpAlignOf:
		return c.alignOfOp(inst.(*ir.AlignOfInst))
	case ir.OpMemSet:
		return c.memSetOp(inst.(*ir.MemSetInst))
	case ir.OpMemCpy:
		return c.memCpyOp(inst.(*ir.MemCpyInst))
	case ir.OpMemMove:
		return c.memMoveOp(inst.(*ir.MemMoveInst))
	case ir.OpStrLen:
		return c.strLenOp(inst.(*ir.StrLenInst))
	case ir.OpMemChr:
		return c.memChrOp(inst.(*ir.MemChrInst))
	case ir.OpMemCmp:
		return c.memCmpOp(inst.(*ir.MemCmpInst))
	case ir.OpRaise:
		return c.raiseOp(inst.(*ir.RaiseInst))

	// Coroutine operations
	case ir.OpCoroId:
		return c.coroIdOp(inst.(*ir.CoroIdInst))
	case ir.OpCoroBegin:
		return c.coroBeginOp(inst.(*ir.CoroBeginInst))
	case ir.OpCoroSuspend:
		return c.coroSuspendOp(inst.(*ir.CoroSuspendInst))
	case ir.OpCoroEnd:
		return c.coroEndOp(inst.(*ir.CoroEndInst))
	case ir.OpCoroFree:
		return c.coroFreeOp(inst.(*ir.CoroFreeInst))

	default:
		return fmt.Errorf("unsupported opcode: %s", inst.Opcode())
	}
}

// Addition
func (c *compiler) addOp(inst ir.Instruction) error {
	ops := inst.Operands()
	lhs := ops[0]
	rhs := ops[1]

	c.loadToReg(RAX, lhs)

	// Check if rhs is a constant
	if constInt, ok := rhs.(*ir.ConstantInt); ok {
		if constInt.Value >= -128 && constInt.Value <= 127 {
			// 8-bit immediate: add rax, imm8 (48 83 C0 ib)
			c.emitBytes(0x48, 0x83, 0xC0, byte(constInt.Value))
		} else {
			// 32-bit immediate: add rax, imm32 (48 81 C0 id)
			c.emitBytes(0x48, 0x81, 0xC0)
			c.emitInt32(int32(constInt.Value))
		}
	} else {
		// Register form: add rax, rcx
		c.loadToReg(RCX, rhs)
		c.emitBytes(0x48, 0x01, 0xC8)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Subtraction
func (c *compiler) subOp(inst ir.Instruction) error {
	ops := inst.Operands()
	lhs := ops[0]
	rhs := ops[1]

	c.loadToReg(RAX, lhs)

	// Check if rhs is a constant
	if constInt, ok := rhs.(*ir.ConstantInt); ok {
		if constInt.Value >= -128 && constInt.Value <= 127 {
			// 8-bit immediate: sub rax, imm8 (48 83 E8 ib)
			c.emitBytes(0x48, 0x83, 0xE8, byte(constInt.Value))
		} else {
			// 32-bit immediate: sub rax, imm32 (48 81 E8 id)
			c.emitBytes(0x48, 0x81, 0xE8)
			c.emitInt32(int32(constInt.Value))
		}
	} else {
		// Register form: sub rax, rcx
		c.loadToReg(RCX, rhs)
		c.emitBytes(0x48, 0x29, 0xC8)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// AND operation
func (c *compiler) andOp(inst ir.Instruction) error {
	ops := inst.Operands()
	lhs := ops[0]
	rhs := ops[1]

	c.loadToReg(RAX, lhs)

	// Check if rhs is a constant
	if constInt, ok := rhs.(*ir.ConstantInt); ok {
		if constInt.Value >= -128 && constInt.Value <= 127 {
			// 8-bit immediate: and rax, imm8 (48 83 E0 ib)
			c.emitBytes(0x48, 0x83, 0xE0, byte(constInt.Value))
		} else {
			// 32-bit immediate: and rax, imm32 (48 81 E0 id)
			c.emitBytes(0x48, 0x81, 0xE0)
			c.emitInt32(int32(constInt.Value))
		}
	} else {
		// Register form: and rax, rcx
		c.loadToReg(RCX, rhs)
		c.emitBytes(0x48, 0x21, 0xC8)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// OR operation
func (c *compiler) orOp(inst ir.Instruction) error {
	ops := inst.Operands()
	lhs := ops[0]
	rhs := ops[1]

	c.loadToReg(RAX, lhs)

	// Check if rhs is a constant
	if constInt, ok := rhs.(*ir.ConstantInt); ok {
		if constInt.Value >= -128 && constInt.Value <= 127 {
			// 8-bit immediate: or rax, imm8 (48 83 C8 ib)
			c.emitBytes(0x48, 0x83, 0xC8, byte(constInt.Value))
		} else {
			// 32-bit immediate: or rax, imm32 (48 81 C8 id)
			c.emitBytes(0x48, 0x81, 0xC8)
			c.emitInt32(int32(constInt.Value))
		}
	} else {
		// Register form: or rax, rcx
		c.loadToReg(RCX, rhs)
		c.emitBytes(0x48, 0x09, 0xC8)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// XOR operation
func (c *compiler) xorOp(inst ir.Instruction) error {
	ops := inst.Operands()
	lhs := ops[0]
	rhs := ops[1]

	c.loadToReg(RAX, lhs)

	// Check if rhs is a constant
	if constInt, ok := rhs.(*ir.ConstantInt); ok {
		if constInt.Value >= -128 && constInt.Value <= 127 {
			// 8-bit immediate: xor rax, imm8 (48 83 F0 ib)
			c.emitBytes(0x48, 0x83, 0xF0, byte(constInt.Value))
		} else {
			// 32-bit immediate: xor rax, imm32 (48 81 F0 id)
			c.emitBytes(0x48, 0x81, 0xF0)
			c.emitInt32(int32(constInt.Value))
		}
	} else {
		// Register form: xor rax, rcx
		c.loadToReg(RCX, rhs)
		c.emitBytes(0x48, 0x31, 0xC8)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Multiplication
func (c *compiler) mulOp(inst ir.Instruction) error {
	ops := inst.Operands()
	c.loadToReg(RAX, ops[0])
	c.loadToReg(RCX, ops[1])

	// imul rax, rcx
	c.emitBytes(0x48, 0x0F, 0xAF, 0xC1)

	c.storeFromReg(RAX, inst)
	return nil
}

// Division and remainder
func (c *compiler) divOp(inst ir.Instruction, remainder bool) error {
	ops := inst.Operands()
	signed := inst.Opcode() == ir.OpSDiv || inst.Opcode() == ir.OpSRem

	c.loadToReg(RAX, ops[0]) // Dividend in RAX
	c.loadToReg(RCX, ops[1]) // Divisor in RCX

	if signed {
		// cqo - sign extend RAX into RDX:RAX
		c.emitBytes(0x48, 0x99)
		// idiv rcx
		c.emitBytes(0x48, 0xF7, 0xF9)
	} else {
		// xor rdx, rdx - zero out RDX
		c.emitBytes(0x48, 0x31, 0xD2)
		// div rcx
		c.emitBytes(0x48, 0xF7, 0xF1)
	}

	// Quotient in RAX, remainder in RDX
	if remainder {
		c.storeFromReg(RDX, inst)
	} else {
		c.storeFromReg(RAX, inst)
	}
	return nil
}

// Floating point binary operations
func (c *compiler) fpBinOp(inst ir.Instruction, opcode byte) error {
	ops := inst.Operands()

	// Load operands to XMM registers
	c.loadToFpReg(0, ops[0]) // XMM0
	c.loadToFpReg(1, ops[1]) // XMM1

	// Determine if single or double precision
	fpType := inst.Type().(*types.FloatType)
	prefix := byte(0xF2) // Default to double (sd)
	if fpType.BitWidth == 32 {
		prefix = 0xF3 // Single precision (ss)
	}

	// Execute operation: XMM0 = XMM0 op XMM1
	c.emitBytes(prefix, 0x0F, opcode, 0xC1)

	c.storeFromFpReg(0, inst)
	return nil
}

// Shift operations
func (c *compiler) shiftOp(inst ir.Instruction, opext byte) error {
	ops := inst.Operands()
	value := ops[0]
	amount := ops[1]

	c.loadToReg(RAX, value)

	if constInt, ok := amount.(*ir.ConstantInt); ok {
		// Immediate shift
		if constInt.Value == 1 {
			// Special encoding for shift by 1: 48 D1 E0+opext
			c.emitBytes(0x48, 0xD1, 0xE0|opext)
		} else {
			// Shift by immediate: 48 C1 E0+opext imm8
			c.emitBytes(0x48, 0xC1, 0xE0|opext, byte(constInt.Value))
		}
	} else {
		// Variable shift (amount in CL): 48 D3 E0+opext
		c.loadToReg(RCX, amount)
		c.emitBytes(0x48, 0xD3, 0xE0|opext)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Alloca - stack allocation
func (c *compiler) allocaOp(inst *ir.AllocaInst) error {
	// Retrieve pre-calculated offset
	allocOffset, ok := c.allocaOffsets[inst]
	if !ok {
		return fmt.Errorf("unknown alloca instruction")
	}

	// lea rax, [rbp + allocOffset] (allocOffset is negative)
	c.emitBytes(0x48, 0x8D, 0x85)
	c.emitInt32(int32(allocOffset))

	// Store the address
	c.storeFromReg(RAX, inst)
	return nil
}

// Load from memory
func (c *compiler) loadOp(inst *ir.LoadInst) error {
	ptr := inst.Operands()[0]
	c.loadToReg(RAX, ptr) // Load pointer address

	// Determine size
	size := SizeOf(inst.Type())

	// For aggregate types, we typically return a pointer to the data on stack
	// EXCEPT if it's small enough to fit in a register (≤ 8 bytes), 
	// in which case we load the value directly to match ABI register passing.
	if types.IsAggregate(inst.Type()) && size > 8 {
		// For large aggregates:
		// Just store the pointer (address of the struct)
		c.storeFromReg(RAX, inst)
		return nil
	}

	// Load value (scalar or small aggregate)
	switch size {
	case 1:
		// movzx rax, byte ptr [rax]
		c.emitBytes(0x48, 0x0F, 0xB6, 0x00)
	case 2:
		// movzx rax, word ptr [rax]
		c.emitBytes(0x48, 0x0F, 0xB7, 0x00)
	case 4:
		// mov eax, [rax] (zero-extends to 64-bit)
		c.emitBytes(0x8B, 0x00)
	case 8:
		// mov rax, [rax]
		c.emitBytes(0x48, 0x8B, 0x00)
	default:
		return fmt.Errorf("unsupported load size: %d", size)
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Store to memory
func (c *compiler) storeOp(inst *ir.StoreInst) error {
	ops := inst.Operands()
	value := ops[0]
	ptr := ops[1]

	size := SizeOf(value.Type())

	// For large structs, we need memcpy
	if size > 8 {
		// This is a struct copy - use memcpy
		// For now, skip it (the value is already in the right place)
		// A proper implementation would:
		// 1. Load source address
		// 2. Load dest address  
		// 3. Call memcpy or use rep movsb
		
		// Simple implementation: assume the struct is already where it needs to be
		return nil
	}

	c.loadToReg(RAX, value) // Value to store
	c.loadToReg(RCX, ptr)   // Pointer

	// mov [rcx], rax (with appropriate size)
	switch size {
	case 1:
		// mov byte ptr [rcx], al
		c.emitBytes(0x88, 0x01)
	case 2:
		// mov word ptr [rcx], ax
		c.emitBytes(0x66, 0x89, 0x01)
	case 4:
		// mov dword ptr [rcx], eax
		c.emitBytes(0x89, 0x01)
	case 8:
		// mov qword ptr [rcx], rax
		c.emitBytes(0x48, 0x89, 0x01)
	default:
		return fmt.Errorf("unsupported store size: %d", size)
	}

	return nil
}

// GetElementPtr - pointer arithmetic
func (c *compiler) gepOp(inst *ir.GetElementPtrInst) error {
	ops := inst.Operands()
	c.loadToReg(RAX, ops[0]) // Base pointer

	currentType := inst.SourceElementType

	for i, idx := range ops[1:] {
		// Calculate offset for this index
		var elemSize int

		if i == 0 {
			// First index: scale by the size of the base type
			elemSize = SizeOf(currentType)
		} else {
			// Subsequent indices: navigate through the type
			switch ty := currentType.(type) {
			case *types.ArrayType:
				elemSize = SizeOf(ty.ElementType)
				currentType = ty.ElementType
			case *types.StructType:
				// For structs, index must be constant
				if constIdx, ok := idx.(*ir.ConstantInt); ok {
					fieldIdx := int(constIdx.Value)
					offset := GetStructFieldOffset(ty, fieldIdx)

					// add rax, offset
					if offset <= 127 {
						c.emitBytes(0x48, 0x83, 0xC0, byte(offset))
					} else {
						c.emitBytes(0x48, 0x05)
						c.emitInt32(int32(offset))
					}

					currentType = ty.Fields[fieldIdx]
					continue
				}
				return fmt.Errorf("struct GEP requires constant index")
			case *types.PointerType:
				elemSize = SizeOf(ty.ElementType)
				currentType = ty.ElementType
			default:
				return fmt.Errorf("invalid GEP type: %T", ty)
			}
		}

		// Load index and multiply by element size
		if constIdx, ok := idx.(*ir.ConstantInt); ok {
			// Constant offset
			offset := int(constIdx.Value) * elemSize
			if offset != 0 {
				if offset >= -128 && offset <= 127 {
					c.emitBytes(0x48, 0x83, 0xC0, byte(offset))
				} else {
					c.emitBytes(0x48, 0x05)
					c.emitInt32(int32(offset))
				}
			}
		} else {
			// Variable offset
			c.loadToReg(RCX, idx)

			// imul rcx, elemSize
			if elemSize == 1 {
				// No scaling needed
			} else if elemSize <= 127 {
				c.emitBytes(0x48, 0x6B, 0xC9, byte(elemSize))
			} else {
				c.emitBytes(0x48, 0x69, 0xC9)
				c.emitInt32(int32(elemSize))
			}

			// add rax, rcx
			c.emitBytes(0x48, 0x01, 0xC8)
		}
	}

	c.storeFromReg(RAX, inst)
	return nil
}

// Integer comparison
func (c *compiler) icmpOp(inst *ir.ICmpInst) error {
	ops := inst.Operands()
	c.loadToReg(RAX, ops[0])
	c.loadToReg(RCX, ops[1])

	// cmp rax, rcx
	c.emitBytes(0x48, 0x39, 0xC8)

	// SETcc al
	var setcc byte
	switch inst.Predicate {
	case ir.ICmpEQ:
		setcc = 0x94 // sete
	case ir.ICmpNE:
		setcc = 0x95 // setne
	case ir.ICmpSLT:
		setcc = 0x9C // setl
	case ir.ICmpSLE:
		setcc = 0x9E // setle
	case ir.ICmpSGT:
		setcc = 0x9F // setg
	case ir.ICmpSGE:
		setcc = 0x9D // setge
	case ir.ICmpULT:
		setcc = 0x92 // setb
	case ir.ICmpULE:
		setcc = 0x96 // setbe
	case ir.ICmpUGT:
		setcc = 0x97 // seta
	case ir.ICmpUGE:
		setcc = 0x93 // setae
	default:
		return fmt.Errorf("unsupported icmp predicate: %v", inst.Predicate)
	}

	c.emitBytes(0x0F, setcc, 0xC0)

	// movzx rax, al
	c.emitBytes(0x48, 0x0F, 0xB6, 0xC0)

	c.storeFromReg(RAX, inst)
	return nil
}

// Floating point comparison
func (c *compiler) fcmpOp(inst *ir.FCmpInst) error {
	ops := inst.Operands()

	c.loadToFpReg(0, ops[0]) // XMM0
	c.loadToFpReg(1, ops[1]) // XMM1

	fpType := ops[0].Type().(*types.FloatType)
	prefix := byte(0xF2)
	if fpType.BitWidth == 32 {
		prefix = 0xF3
	}

	// ucomiss/ucomisd xmm0, xmm1
	c.emitBytes(prefix, 0x0F, 0x2E, 0xC1)

	// Map FCmp predicates to x86 condition codes
	var setcc byte
	switch inst.Predicate {
	case ir.FCmpOEQ:
		setcc = 0x94 // sete (equal, no parity)
	case ir.FCmpONE:
		setcc = 0x95 // setne
	case ir.FCmpOLT:
		setcc = 0x92 // setb (below)
	case ir.FCmpOLE:
		setcc = 0x96 // setbe
	case ir.FCmpOGT:
		setcc = 0x97 // seta (above)
	case ir.FCmpOGE:
		setcc = 0x93 // setae
	default:
		return fmt.Errorf("unsupported fcmp predicate: %v", inst.Predicate)
	}

	c.emitBytes(0x0F, setcc, 0xC0)
	c.emitBytes(0x48, 0x0F, 0xB6, 0xC0) // movzx rax, al

	c.storeFromReg(RAX, inst)
	return nil
}

// System Call (Linux x86_64)
func (c *compiler) syscallOp(inst *ir.SyscallInst) error {
	ops := inst.Operands()
	if len(ops) == 0 {
		return fmt.Errorf("syscall requires at least a syscall number")
	}

	// Linux x86_64 Syscall Calling Convention
	// Syscall Number: RAX
	// Args: RDI, RSI, RDX, R10, R8, R9
	// Return: RAX
	
	// Registers in order for arguments 1..6
	argRegs := []int{RDI, RSI, RDX, R10, R8, R9}

	// 1. Load Syscall Number into RAX (ops[0])
	c.loadToReg(RAX, ops[0])

	// 2. Load Arguments into specific registers
	// Note: args start at ops[1]
	for i, arg := range ops[1:] {
		if i >= len(argRegs) {
			return fmt.Errorf("too many arguments for syscall (max 6 supported)")
		}
		c.loadToReg(argRegs[i], arg)
	}

	// 3. Emit 'syscall' instruction
	// Opcode: 0F 05
	c.emitBytes(0x0F, 0x05)

	// 4. Store result (RAX) to stack slot allocated for this instruction
	// This captures the return value of the syscall
	c.storeFromReg(RAX, inst)

	return nil
}

// Extract value from aggregate
func (c *compiler) extractValueOp(inst *ir.ExtractValueInst) error {
	agg := inst.Operands()[0]
	// Do not load aggregate value to register (as done before), because it might be
	// a large struct or we need to access fields by offset in stack.
	
	// Calculate offset from the base of the aggregate in the stack
	baseOffset, ok := c.stackMap[agg]
	if !ok {
		return fmt.Errorf("extractvalue operand not in stack map: %s", agg.Name())
	}

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

	// Load the value from stack location
	// The address is RBP + baseOffset + offset
	size := SizeOf(inst.Type())
	c.emitLoadFromStack(RAX, baseOffset+offset, size)

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

// ============================================================================
// Cast Operations (Integer/Float Conversions)
// ============================================================================

// Integer cast operations
func (c *compiler) intCastOp(inst *ir.CastInst) error {
	src := inst.Operands()[0]
	c.loadToReg(RAX, src)

	srcSize := SizeOf(src.Type())

	switch inst.Opcode() {
	case ir.OpTrunc:
		// Truncation - just take lower bits (already in RAX)
		// No operation needed, storing will handle it by taking correct size

	case ir.OpZExt:
		// Zero extension
		switch srcSize {
		case 1:
			c.emitBytes(0x48, 0x0F, 0xB6, 0xC0) // movzx rax, al
		case 2:
			c.emitBytes(0x48, 0x0F, 0xB7, 0xC0) // movzx rax, ax
		case 4:
			c.emitBytes(0x89, 0xC0) // mov eax, eax (zero-extends to 64-bit)
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

// Floating point cast operations (float <-> double)
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

	// For bitcast, pointer cast, etc., just copy the bits
	c.loadToReg(RAX, src)
	c.storeFromReg(RAX, inst)

	return nil
}

// ============================================================================
// Variadic Operations
// ============================================================================

// va_start - initialize va_list
func (c *compiler) vaStartOp(inst *ir.VaStartInst) error {
	// System V AMD64 ABI va_list structure:
	// struct {
	//   uint32 gp_offset;     // offset to next GP register arg (0-48)
	//   uint32 fp_offset;     // offset to next FP register arg (48-176)
	//   void*  overflow_arg_area;  // pointer to stack args
	//   void*  reg_save_area;      // pointer to register save area
	// }
	
	vaList := inst.Operands()[0]
	c.loadToReg(RAX, vaList) // Get va_list pointer

	// Initialize gp_offset to 0 (no GP registers consumed yet)
	// mov dword ptr [rax], 0
	c.emitBytes(0xC7, 0x00, 0x00, 0x00, 0x00, 0x00)

	// Initialize fp_offset to 48 (FP registers start after 6 GP regs * 8 bytes)
	// mov dword ptr [rax + 4], 48
	c.emitBytes(0xC7, 0x40, 0x04, 0x30, 0x00, 0x00, 0x00)

	// Set overflow_arg_area to point to stack args (after return address and saved RBP)
	// lea rcx, [rbp + 16]
	c.emitBytes(0x48, 0x8D, 0x8D, 0x10, 0x00, 0x00, 0x00)
	// mov [rax + 8], rcx
	c.emitBytes(0x48, 0x89, 0x48, 0x08)

	// Set reg_save_area to point to the register save area (typically allocated by caller)
	// For simplicity, we'll point to a stack location where we saved registers
	// This would need proper implementation based on calling convention
	// For now, set to NULL as a placeholder
	// mov qword ptr [rax + 16], 0
	c.emitBytes(0x48, 0xC7, 0x40, 0x10, 0x00, 0x00, 0x00, 0x00)

	return nil
}

// va_arg - retrieve next argument
func (c *compiler) vaArgOp(inst *ir.VaArgInst) error {
	vaList := inst.Operands()[0]
	argType := inst.ArgType

	c.loadToReg(RAX, vaList) // Get va_list pointer

	// Simplified implementation: assume all args come from overflow area
	// A full implementation would check gp_offset/fp_offset and use reg_save_area

	// Load overflow_arg_area pointer: mov rcx, [rax + 8]
	c.emitBytes(0x48, 0x8B, 0x48, 0x08)

	// Load the argument from [rcx]
	size := SizeOf(argType)
	switch size {
	case 1:
		c.emitBytes(0x48, 0x0F, 0xB6, 0x09) // movzx rax, byte ptr [rcx]
	case 2:
		c.emitBytes(0x48, 0x0F, 0xB7, 0x09) // movzx rax, word ptr [rcx]
	case 4:
		c.emitBytes(0x8B, 0x01) // mov eax, [rcx]
		// Move to rax for consistency
		c.emitBytes(0x48, 0x89, 0xC0)
	case 8:
		c.emitBytes(0x48, 0x8B, 0x09) // mov rax, [rcx]
	}

	// Advance overflow_arg_area pointer
	// add rcx, 8 (arguments are 8-byte aligned)
	c.emitBytes(0x48, 0x83, 0xC1, 0x08)
	
	// Store updated pointer back: mov [rax + 8], rcx (rax still has va_list ptr)
	c.loadToReg(RDX, vaList) // Reload va_list pointer to RDX
	c.emitBytes(0x48, 0x89, 0x4A, 0x08)

	// Store result
	c.storeFromReg(RAX, inst)
	return nil
}

// va_end - cleanup va_list
func (c *compiler) vaEndOp(inst *ir.VaEndInst) error {
	// On x86_64, va_end is typically a no-op
	// No cleanup needed for the va_list structure
	return nil
}

// ============================================================================
// Intrinsic Operations
// ============================================================================

// sizeof - compile-time size calculation
func (c *compiler) sizeOfOp(inst *ir.SizeOfInst) error {
	size := SizeOf(inst.QueryType)
	
	// Load size as immediate constant
	c.loadConstInt(RAX, int64(size))
	c.storeFromReg(RAX, inst)
	
	return nil
}

// alignof - compile-time alignment calculation
func (c *compiler) alignOfOp(inst *ir.AlignOfInst) error {
	align := AlignOf(inst.QueryType)
	
	// Load alignment as immediate constant
	c.loadConstInt(RAX, int64(align))
	c.storeFromReg(RAX, inst)
	
	return nil
}

// memset - fill memory with a byte value
func (c *compiler) memSetOp(inst *ir.MemSetInst) error {
	ops := inst.Operands()
	dest := ops[0]    // *void
	val := ops[1]     // byte
	count := ops[2]   // usize

	// Load arguments into registers for rep stosb
	c.loadToReg(RDI, dest)   // Destination
	c.loadToReg(RAX, val)    // Value (in AL)
	c.loadToReg(RCX, count)  // Count

	// rep stosb - repeat store byte from AL to [RDI], RCX times
	c.emitBytes(0xF3, 0xAA)

	return nil
}

// memcpy - copy memory
func (c *compiler) memCpyOp(inst *ir.MemCpyInst) error {
	ops := inst.Operands()
	dest := ops[0]    // *void
	src := ops[1]     // *void
	count := ops[2]   // usize

	// Load arguments for rep movsb
	c.loadToReg(RDI, dest)   // Destination
	c.loadToReg(RSI, src)    // Source
	c.loadToReg(RCX, count)  // Count

	// rep movsb - repeat move byte from [RSI] to [RDI], RCX times
	c.emitBytes(0xF3, 0xA4)

	return nil
}

// memmove - copy memory (handles overlapping regions)
func (c *compiler) memMoveOp(inst *ir.MemMoveInst) error {
	ops := inst.Operands()
	dest := ops[0]
	src := ops[1]
	count := ops[2]

	// For simplicity, use same implementation as memcpy
	// A proper implementation would check for overlap and copy backwards if needed
	// Or call the libc memmove function
	
	c.loadToReg(RDI, dest)
	c.loadToReg(RSI, src)
	c.loadToReg(RCX, count)

	// rep movsb
	c.emitBytes(0xF3, 0xA4)

	return nil
}

// strlen - calculate C-string length
func (c *compiler) strLenOp(inst *ir.StrLenInst) error {
	str := inst.Operands()[0]
	
	c.loadToReg(RDI, str)    // String pointer
	c.emitXorReg(RAX, RAX)   // RAX = 0 (length counter)
	c.emitXorReg(RCX, RCX)   // RCX = 0

	// Loop to find null terminator
	loopStart := c.text.Len()
	
	// cmp byte ptr [rdi + rax], 0
	c.emitBytes(0x80, 0x3C, 0x07, 0x00)
	
	// je end (jump if zero found)
	c.emitBytes(0x74, 0x04) // Short jump forward 4 bytes
	
	// inc rax
	c.emitBytes(0x48, 0xFF, 0xC0)
	
	// jmp loop_start
	rel := loopStart - (c.text.Len() + 2)
	c.emitBytes(0xEB, byte(rel))
	
	// end: result in RAX
	c.storeFromReg(RAX, inst)
	return nil
}

// memchr - find byte in memory
func (c *compiler) memChrOp(inst *ir.MemChrInst) error {
	ops := inst.Operands()
	ptr := ops[0]
	val := ops[1]
	count := ops[2]

	c.loadToReg(RDI, ptr)    // Buffer pointer
	c.loadToReg(RAX, val)    // Byte to find
	c.loadToReg(RCX, count)  // Count

	// repne scasb - repeat while not equal, scan byte
	// Compares AL with [RDI], increments RDI, decrements RCX
	c.emitBytes(0xF2, 0xAE)

	// If found, RDI points one past the found byte
	// We need to return RDI-1, or NULL if not found

	// Check if found (ZF=1)
	// jne not_found
	c.emitBytes(0x75, 0x07) // Jump 7 bytes if not equal

	// Found: sub rdi, 1 (point to the actual byte)
	c.emitBytes(0x48, 0x83, 0xEF, 0x01)
	// mov rax, rdi
	c.emitBytes(0x48, 0x89, 0xF8)
	// jmp end
	c.emitBytes(0xEB, 0x03) // Jump 3 bytes

	// not_found: xor rax, rax (return NULL)
	c.emitXorReg(RAX, RAX)

	// end:
	c.storeFromReg(RAX, inst)
	return nil
}

// memcmp - compare memory regions
func (c *compiler) memCmpOp(inst *ir.MemCmpInst) error {
	ops := inst.Operands()
	ptr1 := ops[0]
	ptr2 := ops[1]
	count := ops[2]

	c.loadToReg(RSI, ptr1)   // First buffer
	c.loadToReg(RDI, ptr2)   // Second buffer
	c.loadToReg(RCX, count)  // Count

	// repe cmpsb - repeat while equal, compare bytes
	c.emitBytes(0xF3, 0xA6)

	// After repe cmpsb:
	// - If all equal: ZF=1, RSI and RDI point past end
	// - If different: ZF=0, RSI and RDI point past the differing byte

	// We need to return: 0 if equal, <0 if ptr1 < ptr2, >0 if ptr1 > ptr2
	
	// Load the last compared bytes
	// movzx rax, byte ptr [rsi - 1]
	c.emitBytes(0x48, 0x0F, 0xB6, 0x46, 0xFF)
	
	// movzx rcx, byte ptr [rdi - 1]
	c.emitBytes(0x48, 0x0F, 0xB6, 0x4F, 0xFF)
	
	// sub rax, rcx (compute difference)
	c.emitBytes(0x48, 0x29, 0xC8)

	c.storeFromReg(RAX, inst)
	return nil
}

// raise - abort execution with message
func (c *compiler) raiseOp(inst *ir.RaiseInst) error {
	// Implementation: call exit(1) or trigger SIGABRT
	// For simplicity, we'll use exit syscall
	
	// Load exit code 1 into RDI
	c.loadConstInt(RDI, 1)
	
	// Load syscall number for exit (60 on Linux x86_64)
	c.loadConstInt(RAX, 60)
	
	// syscall
	c.emitBytes(0x0F, 0x05)
	
	// This never returns
	return nil
}