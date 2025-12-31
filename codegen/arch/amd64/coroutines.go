package amd64

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
)

func (c *compiler) compileCoroInst(inst ir.Instruction) error {
	switch inst.Opcode() {
	case ir.OpCoroId:
		// Placeholder: Return instruction address as ID
		c.asm.Mov(RegOp(RAX), ImmOp(c.asm.Len()), 64)
		c.store(RAX, inst)

	case ir.OpCoroBegin:
		// 1. Allocate Frame (simplified malloc/mmap via syscall)
		// MMAP: rax=9, rdi=0, rsi=size, rdx=3(RW), r10=34(Anon|Priv), r8=-1, r9=0
		frameSize := 256
		
		c.asm.Xor(RegOp(RDI), RegOp(RDI))              // addr = 0
		c.asm.Mov(RegOp(RSI), ImmOp(frameSize), 64)    // size
		c.asm.Mov(RegOp(RDX), ImmOp(3), 64)            // prot = RW
		c.asm.Mov(RegOp(R10), ImmOp(34), 64)           // flags
		c.asm.Mov(RegOp(R8),  ImmOp(-1), 64)           // fd = -1
		c.asm.Xor(RegOp(R9),  RegOp(R9))               // off = 0
		c.asm.Mov(RegOp(RAX), ImmOp(9), 64)            // SYS_mmap
		c.asm.Syscall()
		
		// RAX has pointer, store it
		c.store(RAX, inst)

	case ir.OpCoroSuspend:
		// Return 0 (Suspended) for now
		c.asm.Xor(RegOp(RAX), RegOp(RAX))
		c.store(RAX, inst)
		
		// Real implementation needs to save registers to the frame pointer 
		// generated in CoroBegin, then RET to caller.

	case ir.OpCoroEnd:
		// Just return "should destroy" = 1
		c.asm.Mov(RegOp(RAX), ImmOp(1), 64)
		// Logic to jump to cleanup would go here

	case ir.OpCoroFree:
		// Return the frame pointer handle
		handle := inst.Operands()[1]
		c.load(RAX, handle)
		c.store(RAX, inst)

	default:
		return fmt.Errorf("unknown coro opcode: %s", inst.Opcode())
	}
	return nil
}