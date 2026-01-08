package symbol

import (
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

// SymbolKind categorizes what a name represents in the language.
type SymbolKind int

const (
	SymVar      SymbolKind = iota // Variables, Parameters
	SymConst                      // Compile-time constants
	SymFunc                       // Functions
	SymType                       // Structs, Classes, Type Aliases
	SymModule                     // Imported packages/modules
)

func (k SymbolKind) String() string {
	switch k {
	case SymVar:
		return "Variable"
	case SymConst:
		return "Constant"
	case SymFunc:
		return "Function"
	case SymType:
		return "Type"
	case SymModule:
		return "Module"
	default:
		return "Unknown"
	}
}

// Symbol represents a named entity in the source code.
// It acts as the bridge between Semantic Analysis and Code Generation.
type Symbol struct {
	// --- Fixed Properties (Set during Pass 1) ---
	
	// Name is the identifier string (e.g., "x", "calculate")
	Name string
	
	// Kind determines how the compiler treats this symbol
	Kind SymbolKind
	
	// Type is the high-level type info (e.g., i64, *i8, func(i32)void)
	Type types.Type
	
	// IsDefined tracks if the symbol is fully defined or just forward-declared.
	// Useful for recursive structs or function prototypes.
	IsDefined bool

	// --- C++ Interop / Virtual Method Support ---
	
	// IsVirtual indicates if this is a C++ virtual method that requires vtable lookup.
	IsVirtual bool
	
	// VTableIndex is the index in the vtable where this method's pointer resides.
	VTableIndex int

	// --- Mutable Properties (Set during Pass 2) ---

	// IRValue holds the LLVM IR representation.
	// - For SymVar: It's the *ir.AllocaInst (pointer to memory)
	// - For SymFunc: It's the *ir.Function
	// - For SymConst: It's the *ir.Constant
	// This remains nil until Pass 2 visits the declaration.
	IRValue ir.Value
}

// String returns a debug representation of the symbol
func (s *Symbol) String() string {
	valStr := "nil"
	if s.IRValue != nil {
		valStr = s.IRValue.String()
	}
	virt := ""
	if s.IsVirtual {
		virt = fmt.Sprintf(" [Virtual #%d]", s.VTableIndex)
	}
	return fmt.Sprintf("%s (%s)%s : %s -> %s", s.Name, s.Kind, virt, s.Type, valStr)
}