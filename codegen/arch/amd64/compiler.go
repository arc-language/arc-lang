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
	// Reserve space for args and locals
	offset := 0
	
	// Arguments (System V ABI: first 6 in regs, rest on stack)
	// For simplicity, we spill register args to shadow stack slots immediately.
	for _, arg := range fn.Arguments {
		size := SizeOf(arg.Type())
		offset += size
		if offset % 8 != 0 { offset += 8 - (offset%8) } // Align to 8
		c.stackMap[arg] = -offset
	}

	// Instructions
	for _, block := range fn.Blocks {
		for _, inst := range block.Instructions {
			// Void types don't need storage (e.g. call void, store)
			if inst.Type() != nil && inst.Type().Kind() != types.VoidKind {
				size := SizeOf(inst.Type())
				// Alloca instruction returns a pointer, so we need 8 bytes for the pointer itself.
				// However, `ir.AllocaInst` logic in `ops.go` handles the actual allocation size differently.
				// Here we just allocate the slot for the *result* of the instruction.
				// For alloca, the result is the pointer.
				if size == 0 { continue }
				
				offset += size
				if offset % 8 != 0 { offset += 8 - (offset%8) }
				c.stackMap[inst] = -offset
			}
			
			// If it is an AllocaInst, we ALSO need space for the allocated data itself.
			if alloca, ok := inst.(*ir.AllocaInst); ok {
				allocSize := SizeOf(alloca.AllocatedType)
				// Handle array counts if constant
				if count, ok := alloca.NumElements.(*ir.ConstantInt); ok {
					allocSize *= int(count.Value)
				}
				offset += allocSize
				if offset % 16 != 0 { offset += 16 - (offset%16) }
				
				// We store the frame-relative offset in a special way or just rely on LEA calculation.
				// To keep it simple: we just expanded 'offset'. 
				// We need a way to map the AllocaInst to this data area.
				// Let's use a secondary map or encode it.
				// Current hack: The `stackMap[inst]` holds the POINTER. 
				// The data is at `RBP - current_offset`.
				c.stackMap[ir.Value(alloca)] = -offset // Point directly to data area? 
				// Actually, `OpAlloca` implementation usually LEAs the address into a register.
				// So `stackMap[alloca]` should ideally point to the DATA area.
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
			// Move register to stack
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

// load puts the value of `src` into `dst` register.
func (c *compiler) load(dst Register, src ir.Value) {
	switch v := src.(type) {
	case *ir.ConstantInt:
		c.asm.Mov(RegOp(dst), ImmOp(v.Value), 64)
	case *ir.ConstantNull:
		c.asm.Xor(RegOp(dst), RegOp(dst))
	case *ir.ConstantZero:
		c.asm.Xor(RegOp(dst), RegOp(dst))
	case *ir.Global:
		// Globals are pointers. We load their ADDRESS (LEA).
		// Accessing the value requires a subsequent Load instruction in IR.
		c.asm.LeaRel(dst, v.Name())
	case *ir.AllocaInst:
		// Alloca is an instruction that produces a pointer.
		// If it's in the stack map, we usually LEA it.
		// NOTE: In `compileFunction`, we mapped AllocaInst to its data offset.
		off := c.stackMap[v]
		c.asm.Lea(dst, NewMem(RBP, off))
	default:
		// Load from stack slot
		slot := c.getStackSlot(v)
		
		// Determine size for MOV/MOVZX
		size := SizeOf(v.Type())
		
		if size == 8 {
			c.asm.Mov(RegOp(dst), slot, 64)
		} else if size == 4 {
			// MOV r32, r/m32 (zero extends to r64 automatically)
			c.asm.Mov(RegOp(dst), slot, 32)
		} else if size == 1 {
			// MOVZX
			c.asm.MovZX(dst, slot, 8)
		} else {
			// Fallback
			c.asm.Mov(RegOp(dst), slot, 64)
		}
	}
}

// store puts the value in `src` register into `dst`'s stack slot.
func (c *compiler) store(src Register, dst ir.Value) {
	slot := c.getStackSlot(dst)
	size := SizeOf(dst.Type())
	c.asm.Mov(slot, RegOp(src), size*8)
}

func (c *compiler) emitGlobal(g *ir.Global) error {
	if g.Initializer != nil {
		return c.emitConstant(g.Initializer)
	}
	// Zero init
	size := SizeOf(g.Type())
	c.data.Write(make([]byte, size))
	return nil
}

func (c *compiler) emitConstant(k ir.Constant) error {
	switch v := k.(type) {
	case *ir.ConstantInt:
		size := SizeOf(v.Type())
		buf := make([]byte, 8) // Max size
		// We can use binary.PutVarint but we need fixed size
		// Simple implementation
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