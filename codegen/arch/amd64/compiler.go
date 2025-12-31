package amd64

import (
	"bytes"
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
)

type Artifact struct {
	Text   []byte
	Data   []byte
	Relocs []RelocationRecord
	Symbols []SymbolDef
}

type SymbolDef struct {
	Name   string
	Offset uint64
	Size   uint64
	IsFunc bool
}

type compiler struct {
	asm          *Assembler
	data         *bytes.Buffer
	stackMap     map[ir.Value]int // Value -> RBP offset (negative)
	frameSize    int
	blockOffsets map[*ir.BasicBlock]int
	jumpsToFix   []jumpFixup
	
	// Compilation State
	currentFunc *ir.Function
}

type jumpFixup struct {
	asmOffset int
	target    *ir.BasicBlock
}

func Compile(m *ir.Module) (*Artifact, error) {
	c := &compiler{
		asm:      NewAssembler(),
		data:     new(bytes.Buffer),
		stackMap: make(map[ir.Value]int),
	}

	var syms []SymbolDef

	// 1. Compile Globals (Data Section)
	for _, g := range m.Globals {
		// Align
		for c.data.Len()%8 != 0 { c.data.WriteByte(0) }
		
		offset := c.data.Len()
		c.emitGlobal(g)
		syms = append(syms, SymbolDef{
			Name: g.Name(), Offset: uint64(offset), Size: uint64(c.data.Len() - offset), IsFunc: false,
		})
	}

	// 2. Compile Functions (Text Section)
	for _, fn := range m.Functions {
		if len(fn.Blocks) == 0 { continue } // External
		
		start := c.asm.Len()
		if err := c.compileFunction(fn); err != nil {
			return nil, err
		}
		end := c.asm.Len()

		syms = append(syms, SymbolDef{
			Name: fn.Name(), Offset: uint64(start), Size: uint64(end - start), IsFunc: true,
		})
	}

	return &Artifact{
		Text:    c.asm.Bytes(),
		Data:    c.data.Bytes(),
		Relocs:  c.asm.Relocs,
		Symbols: syms,
	}, nil
}

func (c *compiler) compileFunction(fn *ir.Function) error {
	c.currentFunc = fn
	c.stackMap = make(map[ir.Value]int)
	c.blockOffsets = make(map[*ir.BasicBlock]int)
	c.jumpsToFix = nil
	
	// 1. Stack Layout
	// RBP points to saved RBP. Arguments at RBP+16... Locals at RBP-XXX
	offset := 0
	
	// Map arguments
	for _, arg := range fn.Arguments {
		// Simple implementation: All args go to stack slots for spill/fill
		// Ideally: Register allocation
		size := SizeOf(arg.Type())
		offset += size
		if offset % 8 != 0 { offset += 8 - (offset%8) }
		c.stackMap[arg] = -offset
	}

	// Map Instructions
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if inst.Type() != nil {
				size := SizeOf(inst.Type())
				if size == 0 { continue }
				offset += size
				if offset % 8 != 0 { offset += 8 - (offset%8) }
				c.stackMap[inst] = -offset
			}
		}
	}
	
	// Align frame to 16 bytes
	if offset % 16 != 0 { offset += 16 - (offset%16) }
	c.frameSize = offset

	// 2. Prologue
	c.asm.Push(RBP)
	c.asm.Mov(RegOp(RBP), RegOp(RSP), 64)
	if c.frameSize > 0 {
		c.asm.Sub(RegOp(RSP), ImmOp(c.frameSize))
	}

	// 3. Save Register Args (Spill to stack immediately for simplicity)
	// System V: RDI, RSI, RDX, RCX, R8, R9
	regs := []Register{RDI, RSI, RDX, RCX, R8, R9}
	for i, arg := range fn.Arguments {
		if i < len(regs) {
			// Move reg to stack slot
			slot := c.getStackSlot(arg)
			c.asm.Mov(slot, RegOp(regs[i]), SizeOf(arg.Type()))
		}
	}

	// 4. Body
	for _, block := range fn.Blocks {
		c.blockOffsets[block] = c.asm.Len()
		for _, inst := range block.Instructions {
			if err := c.compileInst(inst); err != nil {
				return err
			}
		}
	}

	// 5. Fixup Jumps
	for _, fix := range c.jumpsToFix {
		targetOff, ok := c.blockOffsets[fix.target]
		if !ok { return fmt.Errorf("jump target not found") }
		// Relative jump: Target - (InstAddr + 4)
		rel := int32(targetOff - (fix.asmOffset + 4))
		c.asm.PatchInt32(fix.asmOffset, rel)
	}

	return nil
}

// Helpers

func (c *compiler) getStackSlot(v ir.Value) MemOp {
	off, ok := c.stackMap[v]
	if !ok {
		// Should be constant or global if not in stackMap
		panic(fmt.Sprintf("Value %v not allocated", v))
	}
	return NewMem(RBP, off)
}

func (c *compiler) load(dst Register, src ir.Value) {
	// If constant
	if cst, ok := src.(*ir.ConstantInt); ok {
		c.asm.Mov(RegOp(dst), ImmOp(cst.Value), 64)
		return
	}
	// If stack
	slot := c.getStackSlot(src)
	c.asm.Mov(RegOp(dst), slot, 64) // Assuming 64-bit ops for now
}

func (c *compiler) store(src Register, dst ir.Value) {
	slot := c.getStackSlot(dst)
	c.asm.Mov(slot, RegOp(src), 64)
}

func (c *compiler) emitGlobal(g *ir.Global) {
	// Simple zero init or raw data
	size := SizeOf(g.Type())
	c.data.Write(make([]byte, size))
}