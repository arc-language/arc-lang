package symbol

import "github.com/arc-language/arc-lang/builder/types"

func InitGlobalScope(s *Scope) {
	// Integer Types
	s.Define("int", SymType, types.I64)
	s.Define("int8", SymType, types.I8)
	s.Define("int16", SymType, types.I16)
	s.Define("int32", SymType, types.I32)
	s.Define("int64", SymType, types.I64)
	
	s.Define("uint8", SymType, types.U8)
	s.Define("uint16", SymType, types.U16)
	s.Define("uint32", SymType, types.U32)
	s.Define("uint64", SymType, types.U64)
	
	// Aliases
	s.Define("byte", SymType, types.U8)
	s.Define("usize", SymType, types.U64) // Arch dependent usually, fixed to u64 for now
	s.Define("isize", SymType, types.I64)
	
	// Float Types
	s.Define("float", SymType, types.F64)
	s.Define("float32", SymType, types.F32)
	s.Define("float64", SymType, types.F64)
	
	// Misc
	s.Define("bool", SymType, types.I1)
	s.Define("void", SymType, types.Void)
	s.Define("char", SymType, types.I32) // Rune
	s.Define("string", SymType, types.NewPointer(types.I8))
}