package amd64

import (
	"bytes"
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

type Artifact struct {
	Text    []byte
	Data    []byte
	Relocs  []RelocationRecord
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

	// 1. Compile Globals
	for _, g := range m.Globals {
		// Align data to 8 bytes
		for c.data.Len()%8 != 0 { c.data.WriteByte(0) }
		
		offset := c.data.Len()
		if err := c.emitGlobal(g); err != nil {
			return nil, err
		}
		
		syms = append(syms, SymbolDef{
			Name: g.Name(), Offset: uint64(offset), Size: uint64(c.data.Len() - offset), IsFunc: false,
		})
	}

	// 2. Compile Functions
	for _, fn := range m.Functions {
		if len(fn.Blocks) == 0 { continue }
		
		// Align text to 16 bytes (NOP padding)
		for c.asm.Len()%16 != 0 { c.asm.emitByte(0x90) }

		start := c.asm.Len()
		if err := c.compileFunction(fn); err != nil {
			return nil, fmt.Errorf("in function %s: %w", fn.Name(), err)
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
	
	// 1. Stack Allocation
	offset := 0
	
	// Arguments
	for _, arg := range fn.Arguments {
		size := SizeOf(arg.Type())
		offset += size
		if offset % 8 != 0 { offset += 8 - (offset%8) }
		c.stackMap[arg] = -offset
	}

	// Instructions
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			if inst.Type() != nil && inst.Type().Kind() != types.VoidKind {
				size := SizeOf(inst.Type())
				if size == 0 { continue }
				
				offset += size
				if offset % 8 != 0 { offset += 8 - (offset%8) }
				c.stackMap[inst] = -offset
			}
			
			if alloca, ok := inst.(*ir.AllocaInst); ok {
				allocSize := SizeOf(alloca.AllocatedType)
				if count, ok := alloca.NumElements.(*ir.ConstantInt); ok {
					allocSize *= int(count.Value)
				}
				offset += allocSize
				if offset % 16 != 0 { offset += 16 - (offset%16) }
				
				// Map the alloca instruction to the offset of the allocated data
				// Note: stackMap entry for the inst itself (the pointer) was handled above
				// We need a separate way to track the data block, but for now 
				// we assume the 'alloca' instruction *value* resolves to the pointer slot,
				// and we calculate the data address relative to RBP using this offset
				// when handling OpAlloca.
				
				// Actually, `OpAlloca` in ops.go uses `c.stackMap[inst]`.
				// If we overwrite it here, we lose the pointer slot (if we allocated one).
				// Optimization: Don't allocate a pointer slot for AllocaInst. Just use the calculated offset.
				c.stackMap[ir.Value(alloca)] = -offset
			}
		}
	}
	
	// Align frame
	if offset % 16 != 0 { offset += 16 - (offset%16) }
	c.frameSize = offset

	// 2. Prologue
	c.asm.Push(RBP)
	c.asm.Mov(RegOp(RBP), RegOp(RSP), 64)
	if c.frameSize > 0 {
		c.asm.Sub(RegOp(RSP), ImmOp(c.frameSize))
	}

	// 3. Save Register Arguments
	regs := []Register{RDI, RSI, RDX, RCX, R8, R9}
	for i, arg := range fn.Arguments {
		if i < len(regs) {
			slot := c.getStackSlot(arg)
			c.asm.Mov(slot, RegOp(regs[i]), SizeOf(arg.Type())*8)
		}
	}

	// 4. Compile Body
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
		if !ok { return fmt.Errorf("jump target block not found") }
		rel := int32(targetOff - (fix.asmOffset + 4))
		c.asm.PatchInt32(fix.asmOffset, rel)
	}

	return nil
}

// Helpers

func (c *compiler) getStackSlot(v ir.Value) MemOp {
	off, ok := c.stackMap[v]
	if !ok {
		panic(fmt.Sprintf("Value %v (%s) not allocated in stack map", v, v.Name()))
	}
	return NewMem(RBP, off)
}

func (c *compiler) load(dst Register, src ir.Value) {
	switch v := src.(type) {
	case *ir.ConstantInt:
		c.asm.Mov(RegOp(dst), ImmOp(v.Value), 64)
	case *ir.ConstantNull:
		c.asm.Xor(RegOp(dst), RegOp(dst))
	case *ir.ConstantZero:
		c.asm.Xor(RegOp(dst), RegOp(dst))
	case *ir.Global:
		c.asm.LeaRel(dst, v.Name())
	case *ir.AllocaInst:
		off := c.stackMap[v]
		c.asm.Lea(dst, NewMem(RBP, off))
	default:
		slot := c.getStackSlot(v)
		size := SizeOf(v.Type())
		
		if size == 8 {
			c.asm.Mov(RegOp(dst), slot, 64)
		} else if size == 4 {
			c.asm.Mov(RegOp(dst), slot, 32)
		} else if size == 1 {
			c.asm.MovZX(dst, slot, 8)
		} else {
			c.asm.Mov(RegOp(dst), slot, 64)
		}
	}
}

func (c *compiler) store(src Register, dst ir.Value) {
	slot := c.getStackSlot(dst)
	size := SizeOf(dst.Type())
	c.asm.Mov(slot, RegOp(src), size*8)
}

func (c *compiler) emitGlobal(g *ir.Global) error {
	if g.Initializer != nil {
		return c.emitConstant(g.Initializer)
	}
	size := SizeOf(g.Type())
	c.data.Write(make([]byte, size))
	return nil
}

func (c *compiler) emitConstant(k ir.Constant) error {
	switch v := k.(type) {
	case *ir.ConstantInt:
		size := SizeOf(v.Type())
		// Removed unused 'buf' variable here
		val := uint64(v.Value)
		for i := 0; i < size; i++ {
			c.data.WriteByte(byte(val))
			val >>= 8
		}
	case *ir.ConstantArray:
		for _, elem := range v.Elements {
			if err := c.emitConstant(elem); err != nil { return err }
		}
	case *ir.ConstantStruct:
		for _, field := range v.Fields {
			if err := c.emitConstant(field); err != nil { return err }
		}
	case *ir.ConstantZero:
		size := SizeOf(v.Type())
		c.data.Write(make([]byte, size))
	case *ir.ConstantNull:
		c.data.Write(make([]byte, 8))
	default:
		return fmt.Errorf("unsupported constant type: %T", k)
	}
	return nil
}