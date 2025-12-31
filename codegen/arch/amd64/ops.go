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
		// Address is already in stackMap from compileFunction
		offset := c.stackMap[inst.(*ir.AllocaInst)]
		c.asm.Lea(RAX, NewMem(RBP, offset))

	case ir.OpLoad:
		c.load(RCX, inst.Operands()[0]) // Load Pointer Address into RCX
		// Now load value FROM [RCX]
		size := SizeOf(inst.Type())
		if size == 1 {
			c.asm.MovZX(RAX, NewMem(RCX, 0), 8)
		} else if size == 4 {
			c.asm.Mov(RegOp(RAX), NewMem(RCX, 0), 32)
		} else {
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

		currentType := gep.SourceElementType
		
		// Use _ instead of i to fix unused variable error
		for _, idxVal := range gep.Operands()[1:] {
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
					c.load(RCX, idxVal) // Index
					c.asm.ImulImm(RCX, int32(elemSize))
					c.asm.Add(RegOp(RAX), RegOp(RCX))
				}
				currentType = at.ElementType
			} else if pt, ok := currentType.(*types.PointerType); ok {
				elemSize := SizeOf(pt.ElementType)
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
				currentType = pt.ElementType
			}
		}
		c.store(RAX, inst)

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
		// Cast Registers to RegOp to satisfy Operand interface
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