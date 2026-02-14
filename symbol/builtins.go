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

	// --- Collections ---
	s.Define("vector", SymType, types.Void)
	s.Define("map", SymType, types.Void)
	s.Define("array", SymType, types.Void)
	s.Define("matrix", SymType, types.Void)
	s.Define("vector2", SymType, types.Void)
	s.Define("vector4", SymType, types.Void)

	// --- Intrinsics ---
    
    s.Define("alloca", SymFunc, types.NewFunction(types.NewPointer(types.I8), []types.Type{types.I64}, false))

	s.Define("memset", SymFunc, types.NewFunction(types.Void, []types.Type{types.NewPointer(types.Void), types.I32, types.I64}, false))
	s.Define("memcpy", SymFunc, types.NewFunction(types.Void, []types.Type{types.NewPointer(types.Void), types.NewPointer(types.Void), types.I64}, false))
	s.Define("memmove", SymFunc, types.NewFunction(types.Void, []types.Type{types.NewPointer(types.Void), types.NewPointer(types.Void), types.I64}, false))
	s.Define("memcmp", SymFunc, types.NewFunction(types.I32, []types.Type{types.NewPointer(types.Void), types.NewPointer(types.Void), types.I64}, false))
	
    // strlen(*i8) -> u64
    s.Define("strlen", SymFunc, types.NewFunction(types.U64, []types.Type{types.NewPointer(types.I8)}, false))
    
    // syscall(n, ...) -> u64
	s.Define("syscall", SymFunc, types.NewFunction(types.U64, []types.Type{types.I64}, true))
	
    s.Define("raise", SymFunc, types.NewFunction(types.Void, []types.Type{types.NewPointer(types.I8)}, false))
    
    // va_start(v) -> *i8
	s.Define("va_start", SymFunc, types.NewFunction(types.NewPointer(types.I8), []types.Type{types.NewPointer(types.Void)}, false))
    
    // va_end(list) -> void
	s.Define("va_end", SymFunc, types.NewFunction(types.Void, []types.Type{types.NewPointer(types.I8)}, false))
    
    // va_arg is handled specially (generic), but needs a symbol
	s.Define("va_arg", SymFunc, types.NewFunction(types.Void, []types.Type{types.NewPointer(types.I8)}, false))
}
