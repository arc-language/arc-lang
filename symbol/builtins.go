package symbol

import "github.com/arc-language/arc-lang/builder/types"

// InitGlobalScope populates a scope with the language's built-in primitives.
// This ensures Pass 1 and Pass 2 agree on what "int" or "void" means.
func InitGlobalScope(s *Scope) {
	// Integer Types
	s.Define("int", SymType, types.I64)     // Default int
	s.Define("int8", SymType, types.I8)
	s.Define("int16", SymType, types.I16)
	s.Define("int32", SymType, types.I32)
	s.Define("int64", SymType, types.I64)
	
	s.Define("uint8", SymType, types.U8)
	s.Define("uint16", SymType, types.U16)
	s.Define("uint32", SymType, types.U32)
	s.Define("uint64", SymType, types.U64)
	s.Define("byte", SymType, types.U8)
	
	// Float Types
	s.Define("float", SymType, types.F64)   // Default float
	s.Define("float32", SymType, types.F32)
	s.Define("float64", SymType, types.F64)
	
	// Misc
	s.Define("bool", SymType, types.I1)
	s.Define("void", SymType, types.Void)
	s.Define("char", SymType, types.I32) // Rune (UTF-32)
	
	// String (Pointer to i8 for now)
	s.Define("string", SymType, types.NewPointer(types.I8))
}