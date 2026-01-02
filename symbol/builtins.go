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
	// Note: In a real implementation, these would be backed by actual StructTypes
	// loaded from the stdlib. For now, we define them so resolution succeeds.
	s.Define("vector", SymType, types.Void) // Placeholder for generic logic
	s.Define("map", SymType, types.Void)
	s.Define("array", SymType, types.Void)

	// --- Intrinsics (Compiler built-ins) ---
	// Pointers usually represented as *i8 (void*) for generic intrinsics here

	// Stack Allocation
	// func alloca<T>(count: usize = 1) *T
	s.Define("alloca", SymFunc, types.NewPointer(types.I8)) // Return type dynamic in compiler

	// Memory Ops
	// func memset(dest: *void, val: byte, count: usize)
	s.Define("memset", SymFunc, types.Void)
	// func memcpy(dest: *void, src: *void, count: usize)
	s.Define("memcpy", SymFunc, types.Void)
	// func memmove(dest: *void, src: *void, count: usize)
	s.Define("memmove", SymFunc, types.Void)
	// func memcmp(ptr1: *void, ptr2: *void, count: usize) int32
	s.Define("memcmp", SymFunc, types.I32)

	// String Ops
	// func strlen(str: *byte) usize
	s.Define("strlen", SymFunc, types.U64)

	// Variadic / System
	// func syscall(number: usize, ...) usize
	s.Define("syscall", SymFunc, types.U64)
	// func raise(msg: string)
	s.Define("raise", SymFunc, types.Void)
	
	// Variadic Helper (simplification)
	s.Define("va_start", SymFunc, types.NewPointer(types.I8))
	s.Define("va_end", SymFunc, types.Void)
	s.Define("va_arg", SymFunc, types.Void) // Generic return
}