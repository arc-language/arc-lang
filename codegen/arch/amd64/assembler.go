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

	// Case 5: MOV Mem, Imm
	if d, ok := dst.(MemOp); ok {
		if imm, ok := src.(ImmOp); ok {
			// MOV r/m, imm
			
			if size == 8 {
				// C6 /0 ib
				a.encodeRex(false, 0, NoReg, d.Base)
				a.emitByte(0xC6)
				a.encodeModRM(0, d)
				a.emitByte(byte(imm))
				return
			}
			
			if size == 16 {
				// 66 C7 /0 iw
				a.emitByte(0x66)
				a.encodeRex(false, 0, NoReg, d.Base)
				a.emitByte(0xC7)
				a.encodeModRM(0, d)
				// Emit 16-bit imm
				v := uint16(imm)
				a.emitByte(byte(v))
				a.emitByte(byte(v >> 8))
				return
			}
			
			if size == 32 {
				// C7 /0 id
				a.encodeRex(false, 0, NoReg, d.Base)
				a.emitByte(0xC7)
				a.encodeModRM(0, d)
				a.emitInt32(int32(imm))
				return
			}
			
			if size == 64 {
				// REX.W + C7 /0 id
				// Immediate must fit in 32 bits signed
				if int64(int32(imm)) != int64(imm) {
					panic(fmt.Sprintf("MOV Mem64, Imm requires 32-bit signed immediate, got %d", imm))
				}
				a.encodeRex(true, 0, NoReg, d.Base)
				a.emitByte(0xC7)
				a.encodeModRM(0, d)
				a.emitInt32(int32(imm))
				return
			}
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


func (a *Assembler) And(dst, src Operand) {
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(true, Register(s), NoReg, Register(d))
			a.emitByte(0x21)
			a.encodeModRM(Register(s), d)
		}
	}
}

func (a *Assembler) Or(dst, src Operand) {
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(true, Register(s), NoReg, Register(d))
			a.emitByte(0x09)
			a.encodeModRM(Register(s), d)
		}
	}
}

func (a *Assembler) Shl(dst, src Register) {
	// D1 /4 for 1, D3 /4 for CL
	// Assuming CL
	a.encodeRex(true, 0, NoReg, dst)
	a.emitByte(0xD3)
	a.encodeModRM(4, RegOp(dst))
}

func (a *Assembler) Shr(dst, src Register) {
	// D3 /5
	a.encodeRex(true, 0, NoReg, dst)
	a.emitByte(0xD3)
	a.encodeModRM(5, RegOp(dst))
}

func (a *Assembler) Sar(dst, src Register) {
	// D3 /7
	a.encodeRex(true, 0, NoReg, dst)
	a.emitByte(0xD3)
	a.encodeModRM(7, RegOp(dst))
}

func (a *Assembler) Cmp(dst, src Operand) {
	if d, ok := dst.(RegOp); ok {
		if s, ok := src.(RegOp); ok {
			a.encodeRex(true, Register(s), NoReg, Register(d))
			a.emitByte(0x39)
			a.encodeModRM(Register(s), d)
		}
	}
}

func (a *Assembler) Test(dst, src Register) {
	a.encodeRex(true, src, NoReg, dst)
	a.emitByte(0x85)
	a.encodeModRM(src, RegOp(dst))
}

// MovZX emits: MOVZX dst, src
func (a *Assembler) MovZX(dst Register, src Operand, srcSize int) {
	// 0F B6 /r (byte), 0F B7 /r (word)
	if srcSize == 8 {
		if m, ok := src.(MemOp); ok {
			a.encodeRex(true, dst, NoReg, m.Base)
			a.emitByte(0x0F); a.emitByte(0xB6)
			a.encodeModRM(dst, src)
		} else if s, ok := src.(RegOp); ok {
			// MOVZX r64, r8
			a.encodeRex(true, dst, NoReg, Register(s))
			a.emitByte(0x0F); a.emitByte(0xB6)
			a.encodeModRM(dst, src)
		}
	}
}

func (a *Assembler) Movsxd(dst, src Register) {
	a.encodeRex(true, dst, NoReg, src)
	a.emitByte(0x63)
	a.encodeModRM(dst, RegOp(src))
}

// Movsx emits: MOVSX dst, src
func (a *Assembler) Movsx(dst Register, src Operand, srcSize int) {
	// 0F BE /r (byte)
	if srcSize == 8 {
		if m, ok := src.(MemOp); ok {
			a.encodeRex(true, dst, NoReg, m.Base)
			a.emitByte(0x0F); a.emitByte(0xBE)
			a.encodeModRM(dst, src)
		} else if s, ok := src.(RegOp); ok {
			a.encodeRex(true, dst, NoReg, Register(s))
			a.emitByte(0x0F); a.emitByte(0xBE)
			a.encodeModRM(dst, src)
		}
	}
}

func (a *Assembler) Setcc(cc CondCode, dst Register) {
	// 0F 9x /0 (SETcc r/m8)
	// Note: REX not strictly needed if targeting AL, but needed for SIL/DIL
	a.emitByte(0x0F)
	a.emitByte(byte(cc) + 0x10) // 0x8x -> 0x9x
	a.encodeModRM(0, RegOp(dst))
}

func (a *Assembler) ImulImm(dst Register, imm int32) {
	// 69 /r id
	a.encodeRex(true, dst, NoReg, dst)
	a.emitByte(0x69)
	a.encodeModRM(dst, RegOp(dst))
	a.emitInt32(imm)
}

func (a *Assembler) LeaRel(dst Register, symbol string) {
	// LEA dst, [RIP + disp32]
	// REX + 8D /r
	a.encodeRex(true, dst, NoReg, 0)
	a.emitByte(0x8D)
	// ModRM: Mod=00, Reg=dst, RM=101 (RIP-relative)
	reg := byte(dst) & 7
	a.emitByte(0x05 | (reg << 3))
	
	// Record Relocation
	a.Relocs = append(a.Relocs, RelocationRecord{
		Offset: a.Len(),
		Symbol: symbol,
		Type:   RelocPC32,
		Addend: -4,
	})
	a.emitInt32(0)
}

func (a *Assembler) Cvttss2si(dst Register, src MemOp) {
	// F3 0F 2C /r (CVTTSS2SI r32, xmm2/m32)
	// Note: destination is GPR (r32/r64), source is Memory (float)
	a.emitByte(0xF3)
	a.encodeRex(false, dst, NoReg, src.Base) // W=0 for 32-bit dest int
	a.emitByte(0x0F)
	a.emitByte(0x2C)
	a.encodeModRM(dst, src)
}

func (a *Assembler) Cvtsi2ss(dst Register, src MemOp) {
	// F3 0F 2A /r (CVTSI2SS xmm1, r32/m32)
	// Destination is XMM, but we use 'dst' register index as placeholder for XMM0-15
	// Source is Memory (int)
	a.emitByte(0xF3)
	a.encodeRex(false, dst, NoReg, src.Base)
	a.emitByte(0x0F)
	a.emitByte(0x2A)
	a.encodeModRM(dst, src)
}

func (a *Assembler) Movss(dst MemOp, src Register) {
	// F3 0F 11 /r (MOVSS xmm2/m32, xmm1)
	// Store XMM (src) to Memory (dst)
	a.emitByte(0xF3)
	a.encodeRex(false, src, NoReg, dst.Base)
	a.emitByte(0x0F)
	a.emitByte(0x11)
	a.encodeModRM(src, dst)
}

func (a *Assembler) Movd(dst Register, src Register) {
	// 66 0F 6E /r (MOVD xmm, r32/m32) - GPR to XMM
	// We assume dst is XMM index, src is GPR
	a.emitByte(0x66)
	a.encodeRex(false, dst, NoReg, src)
	a.emitByte(0x0F)
	a.emitByte(0x6E)
	a.encodeModRM(dst, RegOp(src))
}

func (a *Assembler) Movq(dst Register, src Register) {
	// 66 REX.W 0F 6E /r (MOVQ xmm, r64/m64) - GPR64 to XMM
	a.emitByte(0x66)
	a.encodeRex(true, dst, NoReg, src)
	a.emitByte(0x0F)
	a.emitByte(0x6E)
	a.encodeModRM(dst, RegOp(src))
}

func (a *Assembler) Cvtss2sd(dst Register, src Register) {
	// F3 0F 5A /r (CVTSS2SD xmm1, xmm2)
	a.emitByte(0xF3)
	a.encodeRex(false, dst, NoReg, src)
	a.emitByte(0x0F)
	a.emitByte(0x5A)
	a.encodeModRM(dst, RegOp(src))
}