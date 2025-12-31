package amd64

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
)

// compileInst is the main dispatcher for all IR instructions
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
		c.asm.Cqo() // Sign extend RAX->RDX
		c.asm.Div(RCX, true)
		c.store(RAX, inst)

	// --- Bitwise ---
	case ir.OpXor:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		c.asm.Xor(RegOp(RAX), RegOp(RCX))
		c.store(RAX, inst)

	// --- Control Flow (Formerly controlflow.go) ---
	case ir.OpRet:
		if len(inst.Operands()) > 0 {
			c.load(RAX, inst.Operands()[0])
		}
		// Epilogue: mov rsp, rbp; pop rbp; ret
		c.asm.Mov(RegOp(RSP), RegOp(RBP), 64)
		c.asm.Pop(RBP)
		c.asm.Ret()

	case ir.OpCall:
		call := inst.(*ir.CallInst)
		// System V ABI: RDI, RSI, RDX, RCX, R8, R9
		regs := []Register{RDI, RSI, RDX, RCX, R8, R9}
		for i, arg := range call.Operands() {
			if i < len(regs) {
				c.load(regs[i], arg)
			} else {
				// Stack args (simplification: push in reverse order)
				// Real implementation needs to align stack to 16 bytes
			}
		}
		
		if call.Callee != nil {
			c.asm.CallRelative(call.Callee.Name())
		} else {
			c.asm.CallRelative(call.CalleeName)
		}
		
		// If returns value, store it
		if call.Type() != nil {
			c.store(RAX, inst)
		}

	case ir.OpBr:
		br := inst.(*ir.BrInst)
		// Handle Phi nodes for the target block
		c.handlePhi(inst.Parent(), br.Target)
		// JMP rel32
		off := c.asm.JmpRel(0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: off, target: br.Target})

	case ir.OpCondBr:
		cbr := inst.(*ir.CondBrInst)
		c.load(RAX, cbr.Condition)
		// TEST RAX, RAX
		c.asm.buf.Write([]byte{0x48, 0x85, 0xC0})
		
		// JZ (Equal to 0) -> FalseBlock
		offFalse := c.asm.JccRel(CondEq, 0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offFalse, target: cbr.FalseBlock})

		// True Path
		c.handlePhi(inst.Parent(), cbr.TrueBlock)
		offTrue := c.asm.JmpRel(0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offTrue, target: cbr.TrueBlock})

	case ir.OpSwitch:
		sw := inst.(*ir.SwitchInst)
		c.load(RAX, sw.Condition)
		
		// Simple comparison chain
		for _, cse := range sw.Cases {
			// CMP RAX, val
			c.asm.buf.Write([]byte{0x48, 0x3D}) // CMP RAX, imm32
			c.asm.emitInt32(int32(cse.Value.Value))
			
			// JE target
			off := c.asm.JccRel(CondEq, 0)
			c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: off, target: cse.Block})
		}
		// Default
		c.handlePhi(inst.Parent(), sw.DefaultBlock)
		offDef := c.asm.JmpRel(0)
		c.jumpsToFix = append(c.jumpsToFix, jumpFixup{asmOffset: offDef, target: sw.DefaultBlock})

	case ir.OpPhi:
		// No-op: handled by the predecessor branch instruction
		return nil

	case ir.OpSelect:
		// Select cond, trueVal, falseVal
		// Becomes: CMOV
		c.load(RAX, inst.Operands()[0]) // Cond
		c.load(RCX, inst.Operands()[1]) // True
		c.load(RDX, inst.Operands()[2]) // False
		
		// TEST RAX, RAX
		c.asm.buf.Write([]byte{0x48, 0x85, 0xC0})
		
		// CMOVZ RCX, RDX (Move False to True if Z=1 aka Cond=0)
		// 0x48 0x0F 0x44 0xCA (CMOVZ RCX, RDX)
		c.asm.buf.Write([]byte{0x48, 0x0F, 0x44, 0xCA})
		
		// Result is in RCX
		c.store(RCX, inst)

	// --- Memory ---
	case ir.OpAlloca:
		offset := c.stackMap[inst.(*ir.AllocaInst)]
		c.asm.Lea(RAX, NewMem(RBP, offset))
		c.store(RAX, inst)

	case ir.OpLoad:
		c.load(RCX, inst.Operands()[0]) // Ptr
		c.asm.Mov(RegOp(RAX), NewMem(RCX, 0), SizeOf(inst.Type()))
		c.store(RAX, inst)

	case ir.OpStore:
		c.load(RAX, inst.Operands()[0]) // Val
		c.load(RCX, inst.Operands()[1]) // Ptr
		c.asm.Mov(NewMem(RCX, 0), RegOp(RAX), SizeOf(inst.Operands()[0].Type()))

	case ir.OpICmp:
		c.load(RAX, inst.Operands()[0])
		c.load(RCX, inst.Operands()[1])
		
		// CMP RAX, RCX
		c.asm.buf.Write([]byte{0x48, 0x39, 0xC8})

		// SETcc AL
		icmp := inst.(*ir.ICmpInst)
		var cc CondCode
		switch icmp.Predicate {
		case ir.ICmpEQ: cc = 0x94 // SETE
		case ir.ICmpNE: cc = 0x95 // SETNE
		case ir.ICmpSLT: cc = 0x9C // SETL
		case ir.ICmpSGT: cc = 0x9F // SETG
		// ... etc
		default: cc = 0x94
		}
		
		c.asm.emitByte(0x0F); c.asm.emitByte(byte(cc)); c.asm.emitByte(0xC0) // SETcc AL
		c.asm.emitByte(0x48); c.asm.emitByte(0x0F); c.asm.emitByte(0xB6); c.asm.emitByte(0xC0) // MOVZX
		c.store(RAX, inst)
	
	// --- Coroutines (Dispatch) ---
	case ir.OpCoroId, ir.OpCoroBegin, ir.OpCoroSuspend, ir.OpCoroEnd, ir.OpCoroFree:
		return c.compileCoroInst(inst)

	default:
		return fmt.Errorf("unknown opcode: %s", inst.Opcode())
	}
	return nil
}

// handlePhi copies values for Phi nodes in the target block
func (c *compiler) handlePhi(from, to *ir.BasicBlock) {
	for _, inst := range to.Instructions {
		if phi, ok := inst.(*ir.PhiInst); ok {
			for _, incoming := range phi.Incoming {
				if incoming.Block == from {
					c.load(RAX, incoming.Value)
					c.store(RAX, phi) // Store to Phi's stack slot
					break
				}
			}
		} else {
			break // Phis are always first
		}
	}
}