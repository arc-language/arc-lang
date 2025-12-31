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
		// but to be safe, we can just store it.
		// If src was 32-bit, MOV EAX, ... zeroes high bits.
		c.store(RAX, inst)

	case ir.OpSExt:
		// We need to sign extend explicitly if the load didn't do it (our load does MOVZX/MOV usually).
		// A proper implementation checks source type.
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
		// We don't "store" the result of alloca to stack because alloca *is* the stack slot.
		// But in IR, %1 = alloca is a value.
		// So we actually need to store the POINTER to %1's slot?
		// Wait, in my `compileFunction`, `stackMap[inst]` points to the DATA.
		// So if any other instruction uses `%1`, `c.load` performs `LEA`.
		// So we don't need to generate code here! The slot is static.
		// EXCEPT if the IR expects `%1` to be a pointer stored in memory?
		// For optimization, we treat alloca values as their addresses.
		// No-op here.

	case ir.OpLoad:
		c.load(RCX, inst.Operands()[0]) // Load Pointer Address into RCX
		// Now load value FROM [RCX]
		// Determine destination size
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

		// Process indices
		currentType := gep.SourceElementType
		
		// Skip the first index if it's 0 (pointer arithmetic base)
		// Usually GEP %ptr, 0, 1 -> Access field 1 of struct at %ptr
		// GEP %ptr, 1 -> Access element 1 of array starting at %ptr
		
		for i, idxVal := range gep.Operands()[1:] {
			// Calculate offset
			if st, ok := currentType.(*types.StructType); ok {
				// Struct access - index must be constant
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
				// Array access
				elemSize := SizeOf(at.ElementType)
				if cIdx, ok := idxVal.(*ir.ConstantInt); ok {
					// Constant index
					offset := int(cIdx.Value) * elemSize
					if offset != 0 {
						c.asm.Add(RegOp(RAX), ImmOp(int64(offset)))
					}
				} else {
					// Dynamic index
					c.load(RCX, idxVal) // Index
					c.asm.ImulImm(RCX, int32(elemSize)) // RCX = Index * Size
					c.asm.Add(RegOp(RAX), RegOp(RCX))
				}
				currentType = at.ElementType
			} else if pt, ok := currentType.(*types.PointerType); ok {
				// Pointer arithmetic (usually first index)
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
			
			// Note: The logic above simplifies multi-level GEPs. 
			// A robust implementation tracks type at every step.
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
		c.asm.Test(RAX, RAX) // TEST RAX, RAX
		
		// JZ (Zero) -> False
		offFalse := c.asm.JccRel(CondEq, 0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offFalse, target: cbr.FalseBlock})
		
		c.handlePhi(inst.Parent(), cbr.TrueBlock)
		offTrue := c.asm.JmpRel(0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offTrue, target: cbr.TrueBlock})

	case ir.OpICmp:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Cmp(RAX, RCX)
		
		icmp := inst.(*ir.ICmpInst)
		var cc CondCode
		switch icmp.Predicate {
		case ir.ICmpEQ: cc = CondEq
		case ir.ICmpNE: cc = CondNe
		case ir.ICmpSLT: cc = CondLt
		case ir.ICmpSLE: cc = CondLe
		case ir.ICmpSGT: cc = CondGt
		case ir.ICmpSGE: cc = CondGe
		// Unsigned
		case ir.ICmpULT: cc = CondBlo
		case ir.ICmpULE: cc = CondBle
		case ir.ICmpUGT: cc = CondA
		case ir.ICmpUGE: cc = CondAe
		default: cc = CondEq
		}
		
		// SETcc AL
		c.asm.Setcc(cc, RAX)
		// MOVZX RAX, AL (Clear high bits)
		c.asm.MovZX(RAX, RegOp(RAX), 8)
		c.store(RAX, inst)

	case ir.OpPhi:
		// No-op
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
					// Store to the stack slot allocated for the PHI instruction
					c.store(RAX, phi)
					break
				}
			}
		}
	}
}