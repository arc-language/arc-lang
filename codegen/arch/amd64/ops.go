package amd64

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (c *compiler) compileInst(inst ir.Instruction) error {
	// Dispatch Coroutine intrinsics
	if inst.Opcode() >= ir.OpCoroId && inst.Opcode() <= ir.OpCoroDone {
		return c.compileCoroInst(inst)
	}

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
		op0 := inst.Operands()[0]
		op1 := inst.Operands()[1]
		c.load(RAX, op0)
		if op0.Type().BitSize() == 32 { c.asm.Movsxd(RAX, RAX) }
		
		c.load(RCX, op1)
		if op1.Type().BitSize() == 32 { c.asm.Movsxd(RCX, RCX) }

		c.asm.Cqo()
		c.asm.Div(RCX, true)
		c.store(RAX, inst)

	case ir.OpUDiv:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Xor(RegOp(RDX), RegOp(RDX))
		c.asm.Div(RCX, false)
		c.store(RAX, inst)

	case ir.OpSRem:
		op0 := inst.Operands()[0]
		op1 := inst.Operands()[1]
		c.load(RAX, op0)
		if op0.Type().BitSize() == 32 { c.asm.Movsxd(RAX, RAX) }

		c.load(RCX, op1)
		if op1.Type().BitSize() == 32 { c.asm.Movsxd(RCX, RCX) }

		c.asm.Cqo()
		c.asm.Div(RCX, true)
		c.store(RDX, inst)

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
		c.load(RCX, inst.Operands()[1])
		c.asm.Shl(RAX, RCX)
		c.store(RAX, inst)

	case ir.OpLShr:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Shr(RAX, RCX)
		c.store(RAX, inst)

	case ir.OpAShr:
		op0 := inst.Operands()[0]
		c.load(RAX, op0)
		if op0.Type().BitSize() == 32 { c.asm.Movsxd(RAX, RAX) }
		c.load(RCX, inst.Operands()[1])
		c.asm.Sar(RAX, RCX)
		c.store(RAX, inst)

	// --- Casts ---
	case ir.OpTrunc:
		// Truncate by loading full size and storing small size
		c.load(RAX, inst.Operands()[0])
		c.store(RAX, inst)

	case ir.OpBitcast:
		c.moveValue(RBP, c.stackMap[inst], inst.Operands()[0])

	case ir.OpZExt:
		c.load(RAX, inst.Operands()[0])
		c.store(RAX, inst)

	case ir.OpSExt:
		src := inst.Operands()[0]
		srcSize := SizeOf(src.Type())
		c.load(RAX, src)
		if srcSize == 4 { c.asm.Movsxd(RAX, RAX) } else if srcSize == 1 { c.asm.Movsx(RAX, RAX, 8) }
		c.store(RAX, inst)

	// --- Memory & Aggregates ---
	
	case ir.OpAlloca:
		offset := c.stackMap[inst.(*ir.AllocaInst)]
		c.asm.Lea(RAX, NewMem(RBP, offset))

	case ir.OpLoad:
		ptr := inst.Operands()[0]
		c.load(RCX, ptr) // RCX = Source Address
		c.moveFromMem(RBP, c.stackMap[inst], RCX, 0, SizeOf(inst.Type()))

	case ir.OpStore:
		val := inst.Operands()[0]
		ptr := inst.Operands()[1]
		c.load(RCX, ptr) // RCX = Dest Address
		c.moveValue(RCX, 0, val)

	case ir.OpGetElementPtr:
		gep := inst.(*ir.GetElementPtrInst)
		base := gep.Operands()[0]
		c.load(RAX, base) // RAX = Base Pointer

		indices := gep.Operands()[1:]
		if len(indices) == 0 {
			c.store(RAX, inst)
			return nil
		}

		firstIdx := indices[0]
		baseType := gep.SourceElementType
		baseSize := SizeOf(baseType)

		if cIdx, ok := firstIdx.(*ir.ConstantInt); ok {
			if cIdx.Value != 0 {
				c.asm.Add(RegOp(RAX), ImmOp(int64(int(cIdx.Value)*baseSize)))
			}
		} else {
			c.load(RCX, firstIdx)
			c.asm.ImulImm(RCX, int32(baseSize))
			c.asm.Add(RegOp(RAX), RegOp(RCX))
		}

		currentType := baseType
		for _, idxVal := range indices[1:] {
			if st, ok := currentType.(*types.StructType); ok {
				if cIdx, ok := idxVal.(*ir.ConstantInt); ok {
					offset := GetStructFieldOffset(st, int(cIdx.Value))
					if offset != 0 {
						c.asm.Add(RegOp(RAX), ImmOp(int64(offset)))
					}
					currentType = st.Fields[cIdx.Value]
				}
			} else if at, ok := currentType.(*types.ArrayType); ok {
				elemSize := SizeOf(at.ElementType)
				if cIdx, ok := idxVal.(*ir.ConstantInt); ok {
					c.asm.Add(RegOp(RAX), ImmOp(int64(int(cIdx.Value)*elemSize)))
				} else {
					c.load(RCX, idxVal)
					c.asm.ImulImm(RCX, int32(elemSize))
					c.asm.Add(RegOp(RAX), RegOp(RCX))
				}
				currentType = at.ElementType
			}
		}
		c.store(RAX, inst)

	case ir.OpInsertValue:
		iv := inst.(*ir.InsertValueInst)
		agg := iv.Operands()[0]
		val := iv.Operands()[1]
		destOff := c.stackMap[inst]

		// 1. Copy entire Aggregate
		c.moveValue(RBP, destOff, agg)

		// 2. Calculate offset
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

		// 3. Write new value
		c.moveValue(RBP, destOff+offset, val)

	case ir.OpExtractValue:
		ev := inst.(*ir.ExtractValueInst)
		agg := ev.Operands()[0]
		
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
		
		// Copy from Agg[Offset] to Dest
		if srcSlot, ok := c.stackMap[agg]; ok {
			c.moveFromMem(RBP, c.stackMap[inst], RBP, srcSlot+offset, SizeOf(inst.Type()))
		}

	// --- Control Flow & Intrinsics ---
	
	case ir.OpSizeOf:
		val := SizeOf(inst.(*ir.SizeOfInst).QueryType)
		c.asm.Mov(RegOp(RAX), ImmOp(int64(val)), 64)
		c.store(RAX, inst)

	case ir.OpAlignOf:
		val := AlignOf(inst.(*ir.AlignOfInst).QueryType)
		c.asm.Mov(RegOp(RAX), ImmOp(int64(val)), 64)
		c.store(RAX, inst)

	case ir.OpSyscall:
		ops := inst.Operands()
		regs := []Register{RDI, RSI, RDX, R10, R8, R9}
		c.load(RAX, ops[0])
		for i, arg := range ops[1:] {
			if i < len(regs) { c.load(regs[i], arg) }
		}
		c.asm.Syscall()
		c.store(RAX, inst)

	case ir.OpCall:
		call := inst.(*ir.CallInst)
		regs := []Register{RDI, RSI, RDX, RCX, R8, R9}
		for i, arg := range call.Operands() {
			if i < len(regs) { c.load(regs[i], arg) }
		}
		name := call.CalleeName
		if call.Callee != nil { name = call.Callee.Name() }
		c.asm.CallRelative(name)
		if call.Type() != nil && call.Type().Kind() != types.VoidKind {
			c.store(RAX, inst)
		}

	case ir.OpRet:
		if len(inst.Operands()) > 0 { c.load(RAX, inst.Operands()[0]) }
		c.asm.Mov(RegOp(RSP), RegOp(RBP), 64)
		c.asm.Pop(RBP)
		c.asm.Ret()

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
		
		icmp := inst.(*ir.ICmpInst)
		isSigned := false
		switch icmp.Predicate {
		case ir.ICmpSLT, ir.ICmpSLE, ir.ICmpSGT, ir.ICmpSGE: isSigned = true
		}
		if isSigned && inst.Operands()[0].Type().BitSize() == 32 { c.asm.Movsxd(RAX, RAX) }
		if isSigned && inst.Operands()[1].Type().BitSize() == 32 { c.asm.Movsxd(RCX, RCX) }

		c.asm.Cmp(RegOp(RAX), RegOp(RCX))
		
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

// moveValue copies 'src' Value to [dstBase + dstDisp] handling any size
func (c *compiler) moveValue(dstBase Register, dstDisp int, src ir.Value) {
	// Case: AllocaInst (Pointer R-Value)
	if alloca, ok := src.(*ir.AllocaInst); ok {
		off := c.stackMap[alloca]
		c.asm.Lea(RAX, NewMem(RBP, off))
		c.asm.Mov(NewMem(dstBase, dstDisp), RegOp(RAX), 64)
		return
	}

	size := SizeOf(src.Type())

	if cInt, ok := src.(*ir.ConstantInt); ok {
		// Only supports up to 64-bit constants
		if size <= 4 || (cInt.Value >= -2147483648 && cInt.Value <= 2147483647) {
			c.asm.Mov(NewMem(dstBase, dstDisp), ImmOp(cInt.Value), size*8)
		} else {
			c.asm.Mov(RegOp(RAX), ImmOp(cInt.Value), 64)
			c.asm.Mov(NewMem(dstBase, dstDisp), RegOp(RAX), 64)
		}
		return
	}

	if _, ok := src.(*ir.ConstantZero); ok {
		for i := 0; i < size; i++ {
			c.asm.Mov(NewMem(dstBase, dstDisp+i), ImmOp(0), 8)
		}
		return
	}

	// Memory to Memory copy
	if srcSlot, ok := c.stackMap[src]; ok {
		c.moveFromMem(dstBase, dstDisp, RBP, srcSlot, size)
		return
	}

	// Fallback to load/store for scalars or globals
	if size <= 8 {
		c.load(RAX, src)
		c.asm.Mov(NewMem(dstBase, dstDisp), RegOp(RAX), size*8)
	} else {
		if g, ok := src.(*ir.Global); ok {
			c.asm.LeaRel(RCX, g.Name())
			c.moveFromMem(dstBase, dstDisp, RCX, 0, size)
		} else {
			panic(fmt.Sprintf("Unsupported large move from %T", src))
		}
	}
}

// moveFromMem copies 'size' bytes from [srcBase+srcDisp] to [dstBase+dstDisp] using RAX
func (c *compiler) moveFromMem(dstBase Register, dstDisp int, srcBase Register, srcDisp int, size int) {
	offset := 0
	for offset+8 <= size {
		c.asm.Mov(RegOp(RAX), NewMem(srcBase, srcDisp+offset), 64)
		c.asm.Mov(NewMem(dstBase, dstDisp+offset), RegOp(RAX), 64)
		offset += 8
	}
	if offset+4 <= size {
		c.asm.Mov(RegOp(RAX), NewMem(srcBase, srcDisp+offset), 32)
		c.asm.Mov(NewMem(dstBase, dstDisp+offset), RegOp(RAX), 32)
		offset += 4
	}
	for offset < size {
		c.asm.MovZX(RAX, NewMem(srcBase, srcDisp+offset), 8)
		c.asm.Mov(NewMem(dstBase, dstDisp+offset), RegOp(RAX), 8)
		offset++
	}
}

func (c *compiler) handlePhi(from, to *ir.BasicBlock) {
	for _, inst := range to.Instructions {
		if phi, ok := inst.(*ir.PhiInst); ok {
			for _, incoming := range phi.Incoming {
				if incoming.Block == from {
					c.moveValue(RBP, c.stackMap[phi], incoming.Value)
					break
				}
			}
		}
	}
}