package amd64

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (c *compiler) compileInst(inst ir.Instruction) error {
	switch inst.Opcode() {
	// --- Arithmetic ---
	case ir.OpAdd:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Add(RegOp(RAX), RegOp(RCX))
		c.store(RAX, inst)

	case ir.OpSub:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Sub(RegOp(RAX), RegOp(RCX))
		c.store(RAX, inst)

	case ir.OpMul:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Imul(RAX, RCX)
		c.store(RAX, inst)

	case ir.OpSDiv:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Cqo()
		c.asm.Div(RCX, true)
		c.store(RAX, inst)

	case ir.OpUDiv:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Xor(RegOp(RDX), RegOp(RDX)) // Zero RDX for unsigned div
		c.asm.Div(RCX, false)
		c.store(RAX, inst)

	case ir.OpSRem:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Cqo()
		c.asm.Div(RCX, true)
		c.store(RDX, inst) // Remainder in RDX

	case ir.OpURem:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Xor(RegOp(RDX), RegOp(RDX))
		c.asm.Div(RCX, false)
		c.store(RDX, inst)

	// --- Bitwise ---
	case ir.OpAnd:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.And(RegOp(RAX), RegOp(RCX))
		c.store(RAX, inst)

	case ir.OpOr:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Or(RegOp(RAX), RegOp(RCX))
		c.store(RAX, inst)

	case ir.OpXor:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Xor(RegOp(RAX), RegOp(RCX))
		c.store(RAX, inst)

	case ir.OpShl:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1]) // Shift amount must be in CL (RCX)
		c.asm.Shl(RAX, RCX)
		c.store(RAX, inst)

	case ir.OpLShr:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Shr(RAX, RCX)
		c.store(RAX, inst)

	case ir.OpAShr:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Sar(RAX, RCX)
		c.store(RAX, inst)

	// --- Casts ---
	case ir.OpTrunc, ir.OpBitcast:
		// Just copy bits. Size difference handled by store()
		c.load(RAX, inst.Operands()[0])
		c.store(RAX, inst)

	case ir.OpZExt:
		c.load(RAX, inst.Operands()[0])
		// load() likely zero-extended to 64-bit already if it used MOVZX
		c.store(RAX, inst)

	case ir.OpSExt:
		src := inst.Operands()[0]
		srcSize := SizeOf(src.Type())
		c.load(RAX, src)
		
		if srcSize == 4 {
			c.asm.Movsxd(RAX, RAX) // Sign extend 32->64
		} else if srcSize == 1 {
			c.asm.Movsx(RAX, RAX, 8) // Sign extend 8->64
		}
		c.store(RAX, inst)

	// --- Memory ---
	case ir.OpAlloca:
		// Address is already in stackMap from compileFunction.
		// We calculate the address of the data slot and store it into the pointer slot
		// if the instruction expects a result value (the pointer).
		// Note: compileFunction sets stackMap[inst] to the DATA offset. 
		// But in typical IR, %1 = alloca returns a pointer. 
		// Here we simply LEA that address into RAX.
		offset := c.stackMap[inst.(*ir.AllocaInst)]
		c.asm.Lea(RAX, NewMem(RBP, offset))
		// We don't store it back to memory because AllocaInst doesn't have a separate pointer slot
		// in our simplified allocator; we mostly use it directly or via load().
		// However, if other instructions refer to it as a Value, they expect to load the POINTER.
		// If we treated stackMap[inst] as the data, 'load' would try to load data from it.
		// This is a subtle mismatch. For now, assume callers use LEA logic via c.load() special case.

	case ir.OpLoad:
		c.load(RCX, inst.Operands()[0]) // Load Pointer Address into RCX
		
		// Now load value FROM [RCX]
		size := SizeOf(inst.Type())
		if size == 1 {
			c.asm.MovZX(RAX, NewMem(RCX, 0), 8)
		} else if size == 4 {
			c.asm.Mov(RegOp(RAX), NewMem(RCX, 0), 32)
		} else {
			// For aggregates > 8 bytes, this is tricky in registers.
			// But our allocator spills everything.
			// If it's a large struct, we probably shouldn't be loading it into RAX.
			// We should be doing a memcpy to the dest slot.
			// Simplified: assume <= 8 bytes or pointer.
			c.asm.Mov(RegOp(RAX), NewMem(RCX, 0), 64)
		}
		c.store(RAX, inst)

	case ir.OpStore:
		c.load(RAX, inst.Operands()[0]) // Value
		c.load(RCX, inst.Operands()[1]) // Pointer
		
		size := SizeOf(inst.Operands()[0].Type())
		if size == 1 {
			c.asm.Mov(NewMem(RCX, 0), RegOp(RAX), 8)
		} else if size == 4 {
			c.asm.Mov(NewMem(RCX, 0), RegOp(RAX), 32)
		} else {
			c.asm.Mov(NewMem(RCX, 0), RegOp(RAX), 64)
		}

	case ir.OpGetElementPtr:
		gep := inst.(*ir.GetElementPtrInst)
		base := gep.Operands()[0]
		c.load(RAX, base) // RAX = Base Pointer

		// GEP semantics:
		// First index is ALWAYS pointer arithmetic on SourceElementType.
		// Subsequent indices drill down into the type.
		
		indices := gep.Operands()[1:]
		if len(indices) == 0 {
			c.store(RAX, inst)
			return nil
		}

		// 1. Handle First Index (Pointer Arithmetic)
		firstIdx := indices[0]
		baseType := gep.SourceElementType
		baseSize := SizeOf(baseType)

		if cIdx, ok := firstIdx.(*ir.ConstantInt); ok {
			if cIdx.Value != 0 {
				offset := int(cIdx.Value) * baseSize
				c.asm.Add(RegOp(RAX), ImmOp(int64(offset)))
			}
		} else {
			c.load(RCX, firstIdx)
			c.asm.ImulImm(RCX, int32(baseSize))
			c.asm.Add(RegOp(RAX), RegOp(RCX))
		}

		// 2. Handle Subsequent Indices (Structural Drill-down)
		currentType := baseType
		for _, idxVal := range indices[1:] {
			if st, ok := currentType.(*types.StructType); ok {
				if cIdx, ok := idxVal.(*ir.ConstantInt); ok {
					offset := GetStructFieldOffset(st, int(cIdx.Value))
					if offset != 0 {
						c.asm.Add(RegOp(RAX), ImmOp(int64(offset)))
					}
					currentType = st.Fields[cIdx.Value]
				} else {
					return fmt.Errorf("GEP struct index must be constant")
				}
			} else if at, ok := currentType.(*types.ArrayType); ok {
				elemSize := SizeOf(at.ElementType)
				if cIdx, ok := idxVal.(*ir.ConstantInt); ok {
					offset := int(cIdx.Value) * elemSize
					if offset != 0 {
						c.asm.Add(RegOp(RAX), ImmOp(int64(offset)))
					}
				} else {
					c.load(RCX, idxVal)
					c.asm.ImulImm(RCX, int32(elemSize))
					c.asm.Add(RegOp(RAX), RegOp(RCX))
				}
				currentType = at.ElementType
			} else {
				return fmt.Errorf("indexing into non-aggregate type %T", currentType)
			}
		}
		c.store(RAX, inst)

	// --- Aggregates (Structs/Arrays) ---
	
	case ir.OpInsertValue:
		// %res = insertvalue %agg, %val, idx
		// 1. Copy aggregate from source operand to dest slot
		iv := inst.(*ir.InsertValueInst)
		agg := iv.Operands()[0]
		val := iv.Operands()[1]
		
		// We need to copy the *data* of agg to inst's slot.
		// Since we don't have a memcpy helper exposed here easily, we rely on stack-to-stack moves.
		// Warning: This implementation assumes small structs fitting in registers for now, 
		// or relies on the fact that `c.store` only stores 8 bytes.
		// TODO: Implement proper memcpy for large aggregates.
		
		// For now, load/store the base (likely pointer or small struct)
		c.load(RAX, agg)
		c.store(RAX, inst)
		
		// 2. Calculate offset of the field
		currentType := agg.Type()
		offset := 0
		for _, idx := range iv.Indices {
			if st, ok := currentType.(*types.StructType); ok {
				offset += GetStructFieldOffset(st, idx)
				currentType = st.Fields[idx]
			} else if at, ok := currentType.(*types.ArrayType); ok {
				offset += idx * SizeOf(at.ElementType)
				currentType = at.ElementType
			}
		}
		
		// 3. Store new value at dest + offset
		// Dest address = RBP + instOffset
		destOff := c.stackMap[inst]
		targetAddr := destOff + offset
		
		c.load(RAX, val)
		// mov [rbp + targetAddr], rax
		c.asm.Mov(NewMem(RBP, targetAddr), RegOp(RAX), SizeOf(val.Type())*8)

	case ir.OpExtractValue:
		// %res = extractvalue %agg, idx
		ev := inst.(*ir.ExtractValueInst)
		agg := ev.Operands()[0]
		
		// 1. Calculate offset
		currentType := agg.Type()
		offset := 0
		for _, idx := range ev.Indices {
			if st, ok := currentType.(*types.StructType); ok {
				offset += GetStructFieldOffset(st, idx)
				currentType = st.Fields[idx]
			} else if at, ok := currentType.(*types.ArrayType); ok {
				offset += idx * SizeOf(at.ElementType)
				currentType = at.ElementType
			}
		}
		
		// 2. Load from AggBase + Offset
		// Agg is likely in memory.
		aggBase := c.stackMap[agg] // This might be the pointer if it's alloca, or data if argument
		
		// If agg is Alloca, stackMap points to data.
		// We load from [RBP + aggBase + offset]
		finalOffset := aggBase + offset
		
		// Determine size
		size := SizeOf(inst.Type())
		if size == 1 {
			c.asm.MovZX(RAX, NewMem(RBP, finalOffset), 8)
		} else if size == 4 {
			c.asm.Mov(RegOp(RAX), NewMem(RBP, finalOffset), 32)
		} else {
			c.asm.Mov(RegOp(RAX), NewMem(RBP, finalOffset), 64)
		}
		c.store(RAX, inst)

	// --- Intrinsics ---
	
	case ir.OpSizeOf:
		sz := inst.(*ir.SizeOfInst)
		val := SizeOf(sz.QueryType)
		c.asm.Mov(RegOp(RAX), ImmOp(int64(val)), 64)
		c.store(RAX, inst)

	case ir.OpAlignOf:
		al := inst.(*ir.AlignOfInst)
		val := AlignOf(al.QueryType)
		c.asm.Mov(RegOp(RAX), ImmOp(int64(val)), 64)
		c.store(RAX, inst)

	case ir.OpRaise:
		// raise %val
		// 1. Load exception payload
		c.load(RCX, inst.Operands()[0])
		
		// 2. Find __exception_state global
		c.asm.LeaRel(RAX, "__exception_state")
		
		// 3. Set hasException = 1 (byte at offset 0)
		c.asm.Mov(NewMem(RAX, 0), ImmOp(1), 8)
		
		// 4. Store payload (pointer at offset 8, due to alignment)
		c.asm.Mov(NewMem(RAX, 8), RegOp(RCX), 64)
		
		// 5. Do NOT exit. The try/catch block checks this flag.
		// In a real runtime, we might unwind here, but for this simple model, we fall through.

	case ir.OpSyscall:
		// operands: [syscall_id, arg0, arg1, ...]
		ops := inst.Operands()
		idVal := ops[0]
		args := ops[1:]
		
		// Syscall calling convention: RAX=id, RDI, RSI, RDX, R10, R8, R9
		regs := []Register{RDI, RSI, RDX, R10, R8, R9}
		
		// Load ID to RAX
		c.load(RAX, idVal)
		
		// Load Args
		for i, arg := range args {
			if i < len(regs) {
				c.load(regs[i], arg)
			}
		}
		
		c.asm.Syscall()
		c.store(RAX, inst)

	case ir.OpSelect:
		// %res = select %cond, %true, %false
		ops := inst.Operands()
		c.load(RAX, ops[0]) // Cond
		c.load(RCX, ops[1]) // True
		c.load(RDX, ops[2]) // False
		
		c.asm.Test(RAX, RAX)
		// CMOVZ RCX, RDX (If Zero/False, move FalseVal to Result)
		// Note: CMOVZ is 0F 44
		c.asm.emitByte(0x0F); c.asm.emitByte(0x44)
		c.asm.encodeModRM(RCX, RegOp(RDX)) // dst=RCX, src=RDX
		
		c.store(RCX, inst)

	// --- Control Flow ---
	case ir.OpRet:
		if len(inst.Operands()) > 0 {
			c.load(RAX, inst.Operands()[0])
		}
		c.asm.Mov(RegOp(RSP), RegOp(RBP), 64)
		c.asm.Pop(RBP)
		c.asm.Ret()

	case ir.OpCall:
		call := inst.(*ir.CallInst)
		regs := []Register{RDI, RSI, RDX, RCX, R8, R9}
		for i, arg := range call.Operands() {
			if i < len(regs) {
				c.load(regs[i], arg)
			}
		}
		
		name := call.CalleeName
		if call.Callee != nil { name = call.Callee.Name() }
		c.asm.CallRelative(name)
		
		if call.Type() != nil && call.Type().Kind() != types.VoidKind {
			c.store(RAX, inst)
		}

	case ir.OpBr:
		br := inst.(*ir.BrInst)
		c.handlePhi(inst.Parent(), br.Target)
		off := c.asm.JmpRel(0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: off, target: br.Target})

	case ir.OpCondBr:
		cbr := inst.(*ir.CondBrInst)
		c.load(RAX, cbr.Condition)
		c.asm.Test(RAX, RAX)
		
		offFalse := c.asm.JccRel(CondEq, 0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offFalse, target: cbr.FalseBlock})
		
		c.handlePhi(inst.Parent(), cbr.TrueBlock)
		offTrue := c.asm.JmpRel(0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offTrue, target: cbr.TrueBlock})

	case ir.OpICmp:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Cmp(RegOp(RAX), RegOp(RCX))
		
		icmp := inst.(*ir.ICmpInst)
		var cc CondCode
		switch icmp.Predicate {
		case ir.ICmpEQ: cc = CondEq
		case ir.ICmpNE: cc = CondNe
		case ir.ICmpSLT: cc = CondLt
		case ir.ICmpSLE: cc = CondLe
		case ir.ICmpSGT: cc = CondGt
		case ir.ICmpSGE: cc = CondGe
		case ir.ICmpULT: cc = CondBlo
		case ir.ICmpULE: cc = CondBle
		case ir.ICmpUGT: cc = CondA
		case ir.ICmpUGE: cc = CondAe
		default: cc = CondEq
		}
		
		c.asm.Setcc(cc, RAX)
		c.asm.MovZX(RAX, RegOp(RAX), 8)
		c.store(RAX, inst)

	case ir.OpPhi:
		return nil

	default:
		return fmt.Errorf("unknown opcode: %s", inst.Opcode())
	}
	return nil
}

func (c *compiler) handlePhi(from, to *ir.BasicBlock) {
	for _, inst := range to.Instructions {
		if phi, ok := inst.(*ir.PhiInst); ok {
			for _, incoming := range phi.Incoming {
				if incoming.Block == from {
					c.load(RAX, incoming.Value)
					c.store(RAX, phi)
					break
				}
			}
		}
	}
}