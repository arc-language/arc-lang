// Package mangler provides utilities for C++ ABI name mangling.
package mangler

import (
	"strconv"
	"strings"

	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/symbol"
)

// MangleItanium generates a mangled symbol name based on the Itanium C++ ABI.
func MangleItanium(sym *symbol.Symbol, isMethod bool) string {
	fnType, ok := sym.Type.(*types.FunctionType)
	if !ok {
		return sym.Name
	}

	components := parseComponents(sym.Name)

	var sb strings.Builder
	sb.WriteString("_Z")

	// Nested names need 'N' prefix
	if len(components) > 1 {
		sb.WriteString("N")
	}

	for _, part := range components {
		sb.WriteString(strconv.Itoa(len(part)))
		sb.WriteString(part)
	}

	if len(components) > 1 {
		sb.WriteString("E")
	}

	params := fnType.ParamTypes
	// For methods, the first parameter is the implicit 'this' pointer (self).
	// Itanium ABI does not encode 'this' in the parameter list for the mangled name.
	if isMethod && len(params) > 0 {
		params = params[1:]
	}

	if len(params) == 0 {
		sb.WriteString("v")
	} else {
		for _, p := range params {
			encodeType(p, &sb)
		}
	}
	
	// Note: Itanium ABI typically doesn't encode return type for standard functions,
	// only for templates or special cases.

	if fnType.Variadic {
		sb.WriteString("z") // Ellipsis
	}

	return sb.String()
}

// parseComponents splits the Arc symbol name into C++ namespace/class/method components.
// Arc represents C++ methods as "Namespace.Class_Method".
// We convert the last underscore to a dot, then split by dot.
func parseComponents(name string) []string {
	// Heuristic: If there is an underscore, treat the last segment after it as the method name,
	// provided it comes after the last dot (nesting).
	if lastUnderscore := strings.LastIndex(name, "_"); lastUnderscore != -1 {
		if strings.LastIndex(name, ".") < lastUnderscore {
			name = name[:lastUnderscore] + "." + name[lastUnderscore+1:]
		}
	}
	
	return strings.Split(name, ".")
}

func encodeType(t types.Type, sb *strings.Builder) {
	switch T := t.(type) {
	case *types.VoidType:
		sb.WriteString("v")
	case *types.IntType:
		switch T.BitWidth {
		case 1:  sb.WriteString("b") // bool
		case 8:  if T.Signed { sb.WriteString("c") } else { sb.WriteString("h") } // char / uchar
		case 16: if T.Signed { sb.WriteString("s") } else { sb.WriteString("t") } // short / ushort
		case 32: if T.Signed { sb.WriteString("i") } else { sb.WriteString("j") } // int / uint
		case 64: if T.Signed { sb.WriteString("l") } else { sb.WriteString("m") } // long / ulong (Linux 64)
		// Note: long long (64) is 'x'/'y'. Arc maps int64 to 'l' by default here.
		default: sb.WriteString("l")
		}
	case *types.FloatType:
		switch T.BitWidth {
		case 32: sb.WriteString("f")
		case 64: sb.WriteString("d")
		default: sb.WriteString("d")
		}
	case *types.PointerType:
		sb.WriteString("P")
		encodeType(T.ElementType, sb)
	case *types.StructType:
		// Handle nested struct names (e.g., Math.Vector3)
		parts := strings.Split(T.Name, ".")
		if len(parts) > 1 {
			sb.WriteString("N")
			for _, part := range parts {
				sb.WriteString(strconv.Itoa(len(part)))
				sb.WriteString(part)
			}
			sb.WriteString("E")
		} else {
			sb.WriteString(strconv.Itoa(len(T.Name)))
			sb.WriteString(T.Name)
		}
	default:
		sb.WriteString("v") // Fallback
	}
}