package symbol

import "github.com/arc-language/arc-lang/builder/types"

func InitGlobalScope(s *Scope) {
	// --- Primitive Types ---
	s.Define("int", SymType, types.I64)
	s.Define("int8", SymType, types.I8)
	s.Define("int16", SymType, types.I16)
	s.Define("int32", SymType, types.I32)
	s.Define("int64", SymType, types.I64)
	s.Define("uint8", SymType, types.U8)
	s.Define("uint16", SymType, types.U16)
	s.Define("uint32", SymType, types.U32)
	s.Define("uint64", SymType, types.U64)
	s.Define("byte", SymType, types.U8)
	s.Define("usize", SymType, types.U64)
	s.Define("isize", SymType, types.I64)
	s.Define("float", SymType, types.F64)
	s.Define("float32", SymType, types.F32)
	s.Define("float64", SymType, types.F64)
	s.Define("bool", SymType, types.I1)
	s.Define("void", SymType, types.Void)
	s.Define("char", SymType, types.I32)
	s.Define("string", SymType, types.NewPointer(types.I8))

	// --- Collections (Generic Placeholders) ---
	s.Define("vector", SymType, types.Void)
	s.Define("map", SymType, types.Void)
	s.Define("array", SymType, types.Void)
	
	// --- GPU Types (Previously Keywords) ---
	s.Define("matrix", SymType, types.Void)
	s.Define("vector2", SymType, types.Void)
	s.Define("vector4", SymType, types.Void)

	// --- Intrinsics (Compiler built-ins) ---
	s.Define("alloca", SymFunc, types.NewPointer(types.I8))
	s.Define("memset", SymFunc, types.Void)
	s.Define("memcpy", SymFunc, types.Void)
	s.Define("memmove", SymFunc, types.Void)
	s.Define("memcmp", SymFunc, types.I32)
	s.Define("strlen", SymFunc, types.U64)
	s.Define("syscall", SymFunc, types.U64)
	s.Define("raise", SymFunc, types.Void)
	s.Define("va_start", SymFunc, types.NewPointer(types.I8))
	s.Define("va_end", SymFunc, types.Void)
	s.Define("va_arg", SymFunc, types.Void)
}