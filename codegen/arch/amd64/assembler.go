package amd64

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// RelocationType defines x86-64 relocation types
type RelocationType int

const (
	RelocPC32  RelocationType = 2 // R_X86_64_PC32
	RelocPLT32 RelocationType = 4 // R_X86_64_PLT32
)

// RelocationRecord records where patching is needed
type RelocationRecord struct {
	Offset int
	Symbol string
	Type   RelocationType
	Addend int64
}

// Assembler handles low-level instruction encoding
type Assembler struct {
	buf    *bytes.Buffer
	Relocs []RelocationRecord
}

func NewAssembler() *Assembler {
	return &Assembler{
		buf: new(bytes.Buffer),
	}
}

func (a *Assembler) Bytes() []byte {
	return a.buf.Bytes()
}

func (a *Assembler) Len() int {
	return a.buf.Len()
}

// --- Encoding Primitives ---

// emitByte writes a single byte
func (a *Assembler) emitByte(b byte) {
	a.buf.WriteByte(b)
}

// emitInt32 writes a 32-bit little-endian integer
func (a *Assembler) emitInt32(v int32) {
	binary.Write(a.buf, binary.LittleEndian, v)
}

// emitInt64 writes a 64-bit little-endian integer
func (a *Assembler) emitInt64(v int64) {
	binary.Write(a.buf, binary.LittleEndian, v)
}

// encodeRex emits the REX prefix if needed
func (a *Assembler) encodeRex(w bool, r, x, b Register) {
	// REX Prefix: 0100 WRXB
	// W: 64-bit operand size
	// R: Extension of ModR/M reg field
	// X: Extension of SIB index field
	// B: Extension of ModR/M r/m field or SIB base
	
	var rex byte = 0x40
	needsRex := false

	if w {
		rex |= 0x08
		needsRex = true
	}
	if r >= 8 {
		rex |= 0x04
		needsRex = true
	}
	if x >= 8 {
		rex |= 0x02
		needsRex = true
	}
	if b >= 8 {
		rex |= 0x01
		needsRex = true
	}

	// Spl/Bpl/Sil/Dil access also requires REX on some encoders, 
	// but standard registers R8-R15 strictly require it.
	if needsRex {
		a.emitByte(rex)
	}
}

// encodeModRM emits the ModR/M byte and optional SIB/Disp
func (a *Assembler) encodeModRM(reg Register, rm Operand) {
	regCode := byte(reg) & 7

	switch op := rm.(type) {
	case RegOp:
		// Mode 11: Register-direct addressing
		rmCode := byte(op) & 7
		a.emitByte(0xC0 | (regCode << 3) | rmCode)

	case MemOp:
		// Logic for memory operand encoding is complex (SIB, Disp8 vs Disp32)
		// Simplified for this implementation: Always use Disp32 for safety
		
		baseCode := byte(op.Base) & 7
		
		if op.Index == NoReg {
			// [Base + Disp32]
			// Mode 10 (Disp32)
			if op.Base == RSP || op.Base == R12 {
				// RSP/R12 requires SIB byte even if no index
				// Mod 10, Reg, RM=100 (SIB present)
				a.emitByte(0x80 | (regCode << 3) | 0x04)
				// SIB: Scale=0, Index=100 (None), Base=RSP/R12
				a.emitByte(0x24) 
			} else {
				// Standard [Base + Disp]
				a.emitByte(0x80 | (regCode << 3) | baseCode)
			}
			a.emitInt32(op.Disp)
		} else {
			// SIB Addressing not fully implemented for brevity, 
			// assuming simple Base+Disp for stack access mostly.
			panic("Complex SIB addressing not implemented in this simplified assembler")
		}
	
	default:
		panic(fmt.Sprintf("unsupported operand for ModRM: %T", rm))
	}
}

// --- Instructions ---

// Mov emits: MOV dst, src
func (a *Assembler) Mov(dst, src Operand, size int) {
	// Case 1: MOV Reg, Reg
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(size == 64, Register(s), NoReg, Register(d))
			if size == 8 {
				a.emitByte(0x88)
			} else if size == 16 {
				a.emitByte(0x66)
				a.emitByte(0x89)
			} else {
				a.emitByte(0x89)
			}
			a.encodeModRM(Register(s), d) // MOV r/m, reg
			return
		}
	}

	// Case 2: MOV Reg, Mem
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(MemOp); ok {
			a.encodeRex(size == 64, Register(d), NoReg, s.Base)
			if size == 8 {
				a.emitByte(0x0F); a.emitByte(0xB6) // MOVZX for bytes usually preferred
			} else if size == 16 {
				a.emitByte(0x0F); a.emitByte(0xB7) // MOVZX
			} else {
				a.emitByte(0x8B) // MOV r32/r64, r/m
			}
			a.encodeModRM(Register(d), s)
			return
		}
	}

	// Case 3: MOV Mem, Reg
	if d, ok := dst.(MemOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(size == 64, Register(s), NoReg, d.Base)
			if size == 8 {
				a.emitByte(0x88)
			} else if size == 16 {
				a.emitByte(0x66); a.emitByte(0x89)
			} else {
				a.emitByte(0x89)
			}
			a.encodeModRM(Register(s), d)
			return
		}
	}

	// Case 4: MOV Reg, Imm
	if d, ok := dst.(RegOp); ok {
		if imm, ok := src.(ImmOp); ok {
			// Optimizations (XOR reg, reg) handled by caller or here
			if imm == 0 && size == 64 {
				a.Xor(d, d)
				return
			}
			
			// MOV reg, imm64
			reg := Register(d)
			a.encodeRex(true, NoReg, NoReg, reg)
			a.emitByte(0xB8 | (byte(reg) & 7))
			a.emitInt64(int64(imm))
			return
		}
	}

	panic(fmt.Sprintf("Unsupported MOV combination: %T -> %T", src, dst))
}

// Add emits: ADD dst, src
func (a *Assembler) Add(dst, src Operand) {
	// Only implementing Reg, Reg for now
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(true, Register(s), NoReg, Register(d))
			a.emitByte(0x01)
			a.encodeModRM(Register(s), d)
		} else if imm, ok := src.(ImmOp); ok {
			// ADD reg, imm32
			a.encodeRex(true, 0, NoReg, Register(d))
			if imm >= -128 && imm <= 127 {
				a.emitByte(0x83)
				a.encodeModRM(0, d)
				a.emitByte(byte(imm))
			} else {
				a.emitByte(0x81)
				a.encodeModRM(0, d)
				a.emitInt32(int32(imm))
			}
		}
	}
}

// Sub emits: SUB dst, src
func (a *Assembler) Sub(dst, src Operand) {
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(true, Register(s), NoReg, Register(d))
			a.emitByte(0x29)
			a.encodeModRM(Register(s), d)
		} else if imm, ok := src.(ImmOp); ok {
			a.encodeRex(true, 0, NoReg, Register(d))
			if imm >= -128 && imm <= 127 {
				a.emitByte(0x83)
				a.encodeModRM(5, d) // /5
				a.emitByte(byte(imm))
			} else {
				a.emitByte(0x81)
				a.encodeModRM(5, d)
				a.emitInt32(int32(imm))
			}
		}
	}
}

// Imul emits: IMUL dst, src
func (a *Assembler) Imul(dst, src Register) {
	a.encodeRex(true, dst, NoReg, src)
	a.emitByte(0x0F)
	a.emitByte(0xAF)
	a.encodeModRM(dst, RegOp(src))
}

// Div emits: DIV/IDIV src (Implicit RAX/RDX)
func (a *Assembler) Div(src Register, signed bool) {
	a.encodeRex(true, 0, NoReg, src)
	a.emitByte(0xF7)
	subOp := 6 // DIV
	if signed {
		subOp = 7 // IDIV
	}
	a.encodeModRM(Register(subOp), RegOp(src))
}

// Xor emits: XOR dst, src
func (a *Assembler) Xor(dst, src Operand) {
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(true, Register(s), NoReg, Register(d))
			a.emitByte(0x31)
			a.encodeModRM(Register(s), d)
		}
	}
}

// Push emits: PUSH src
func (a *Assembler) Push(src Register) {
	if src >= 8 {
		a.emitByte(0x41) // REX.B
	}
	a.emitByte(0x50 | (byte(src) & 7))
}

// Pop emits: POP dst
func (a *Assembler) Pop(dst Register) {
	if dst >= 8 {
		a.emitByte(0x41) // REX.B
	}
	a.emitByte(0x58 | (byte(dst) & 7))
}

// Ret emits: RET
func (a *Assembler) Ret() {
	a.emitByte(0xC3)
}

// Syscall emits: SYSCALL
func (a *Assembler) Syscall() {
	a.emitByte(0x0F)
	a.emitByte(0x05)
}

// Lea emits: LEA dst, mem
func (a *Assembler) Lea(dst Register, mem MemOp) {
	a.encodeRex(true, dst, NoReg, mem.Base)
	a.emitByte(0x8D)
	a.encodeModRM(dst, mem)
}

// CallRelative emits: CALL rel32
func (a *Assembler) CallRelative(symbol string) {
	a.emitByte(0xE8)
	// Record relocation for linker
	a.Relocs = append(a.Relocs, RelocationRecord{
		Offset: a.Len(),
		Symbol: symbol,
		Type:   RelocPLT32,
		Addend: -4,
	})
	a.emitInt32(0) // Placeholder
}

// CallReg emits: CALL reg
func (a *Assembler) CallReg(reg Register) {
	a.encodeRex(false, 0, NoReg, reg)
	a.emitByte(0xFF)
	a.encodeModRM(2, RegOp(reg)) // /2
}

// JmpRel emits: JMP rel32 (Unconditional)
func (a *Assembler) JmpRel(offset int32) int {
	a.emitByte(0xE9)
	pos := a.Len()
	a.emitInt32(offset)
	return pos
}

// JccRel emits: Jcc rel32 (Conditional)
func (a *Assembler) JccRel(cond CondCode, offset int32) int {
	a.emitByte(0x0F)
	a.emitByte(byte(cond))
	pos := a.Len()
	a.emitInt32(offset)
	return pos
}

// PatchInt32 allows fixing up jump targets later
func (a *Assembler) PatchInt32(offset int, value int32) {
	// Overwrite the bytes at offset
	old := a.buf.Bytes()
	binary.LittleEndian.PutUint32(old[offset:], uint32(value))
}

// Cqo emits: CQO (Sign extend RAX -> RDX:RAX)
func (a *Assembler) Cqo() {
	a.emitByte(0x48)
	a.emitByte(0x99)
}