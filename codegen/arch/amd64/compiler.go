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
			// Allocate space for result if it has a type (and is not void)
			if inst.Type() != nil && inst.Type().Kind() != types.VoidKind {
				size := SizeOf(inst.Type())
				if size == 0 { continue }
				
				offset += size
				// Align to 8 bytes for simple access
				if offset % 8 != 0 { offset += 8 - (offset%8) }
				c.stackMap[inst] = -offset
			}
			
			// Allocate extra space for AllocaInst underlying memory
			if alloca, ok := inst.(*ir.AllocaInst); ok {
				allocSize := SizeOf(alloca.AllocatedType)
				if count, ok := alloca.NumElements.(*ir.ConstantInt); ok {
					allocSize *= int(count.Value)
				}
				// Align alloca memory to 16 bytes
				if offset % 16 != 0 { offset += 16 - (offset%16) }
				offset += allocSize
				c.stackMap[ir.Value(alloca)] = -offset
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

	// 3. Save Register Arguments
	regs := []Register{RDI, RSI, RDX, RCX, R8, R9}
	for i, arg := range fn.Arguments {
		if i < len(regs) {
			slot := c.getStackSlot(arg)
			// Move registers to stack using appropriate size, or 64-bit for simplicity
			// For structs passed by value in regs, this ABI is simplified.
			c.asm.Mov(slot, RegOp(regs[i]), 64)
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
	case *ir.Function:
		c.asm.LeaRel(dst, v.Name())
	case *ir.AllocaInst:
		// Load the address of the alloca memory, not the stack slot of the instruction
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
			// Default fallback for pointers/etc
			c.asm.Mov(RegOp(dst), slot, 64)
		}
	}
}

func (c *compiler) store(src Register, dst ir.Value) {
	slot := c.getStackSlot(dst)
	size := SizeOf(dst.Type())
	if size == 8 {
		c.asm.Mov(slot, RegOp(src), 64)
	} else if size == 4 {
		c.asm.Mov(slot, RegOp(src), 32)
	} else if size == 1 {
		c.asm.Mov(slot, RegOp(src), 8)
	} else {
		c.asm.Mov(slot, RegOp(src), 64)
	}
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
		val := uint64(v.Value)
		for i := 0; i < size; i++ {
			c.data.WriteByte(byte(val))
			val >>= 8
		}
	case *ir.ConstantFloat:
		// Simplified float emission
		c.data.Write(make([]byte, SizeOf(v.Type())))
	case *ir.ConstantArray:
		for _, elem := range v.Elements {
			if err := c.emitConstant(elem); err != nil { return err }
		}
	case *ir.ConstantStruct:
		st, ok := v.Type().(*types.StructType)
		if !ok { return fmt.Errorf("ConstantStruct has non-struct type") }
		
		currentOffset := 0
		for i, field := range v.Fields {
			// Handle Padding
			targetOffset := GetStructFieldOffset(st, i)
			if targetOffset > currentOffset {
				padding := targetOffset - currentOffset
				c.data.Write(make([]byte, padding))
				currentOffset += padding
			}
			
			if err := c.emitConstant(field); err != nil { return err }
			currentOffset += SizeOf(field.Type())
		}
		// Tail Padding
		totalSize := SizeOf(v.Type())
		if currentOffset < totalSize {
			c.data.Write(make([]byte, totalSize - currentOffset))
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