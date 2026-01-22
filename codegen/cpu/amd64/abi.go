package amd64

import "github.com/arc-language/arc-lang/builder/types"

// SizeOf returns the size in bytes of a type for amd64
func SizeOf(t types.Type) int {
	switch t.Kind() {
	case types.VoidKind:
		return 0
	case types.IntegerKind:
		bits := t.(*types.IntType).BitWidth
		if bits <= 8 { return 1 }
		if bits <= 16 { return 2 }
		if bits <= 32 { return 4 }
		return 8
	case types.FloatKind:
		bits := t.(*types.FloatType).BitWidth
		if bits <= 32 { return 4 }
		return 8
	case types.PointerKind, types.FunctionKind:
		return 8
	case types.ArrayKind:
		at := t.(*types.ArrayType)
		return int(at.Length) * SizeOf(at.ElementType)
	case types.StructKind:
		st := t.(*types.StructType)
		size := 0
		
		// --- FIX: Handle Class Header ---
		if st.IsClass {
			size = 8 // Implicit i64 RefCount header
		}

		if st.Packed {
			for _, f := range st.Fields { size += SizeOf(f) }
			if size == 0 { return 1 }
			return size
		}
		
		// Aligned size
		for _, f := range st.Fields {
			align := AlignOf(f)
			if size%align != 0 {
				size += align - (size % align)
			}
			size += SizeOf(f)
		}
		
		if size == 0 { size = 1 }

		// Pad end
		sa := AlignOf(st)
		if size%sa != 0 {
			size += sa - (size % sa)
		}
		return size
	default:
		return 8
	}
}

func AlignOf(t types.Type) int {
	switch t.Kind() {
	case types.IntegerKind:
		bits := t.(*types.IntType).BitWidth
		if bits <= 8 { return 1 }
		if bits <= 16 { return 2 }
		if bits <= 32 { return 4 }
		return 8
	case types.StructKind:
		st := t.(*types.StructType)
		max := 1
		
		// Classes are always at least 8-byte aligned due to header
		if st.IsClass {
			max = 8
		}

		for _, f := range st.Fields {
			a := AlignOf(f)
			if a > max { max = a }
		}
		return max
	case types.ArrayKind:
		return AlignOf(t.(*types.ArrayType).ElementType)
	default:
		sz := SizeOf(t)
		if sz > 8 { return 8 }
		return sz
	}
}

// GetStructFieldOffset calculates the byte offset of a field
func GetStructFieldOffset(st *types.StructType, idx int) int {
	off := 0
	
	// --- FIX: Handle Class Header and Index Shifting ---
	targetIdx := idx
	
	if st.IsClass {
		// Index 0 refers to the RefCount Header
		if idx == 0 {
			return 0
		}
		// Index 1 refers to Fields[0], Index 2 to Fields[1], etc.
		off = 8 // Start after header
		targetIdx = idx - 1 // Shift index for the loop below
	}

	for i := 0; i < targetIdx; i++ {
		f := st.Fields[i]
		if !st.Packed {
			a := AlignOf(f)
			if off%a != 0 { off += a - (off % a) }
		}
		off += SizeOf(f)
	}
	
	// Align target
	if !st.Packed && targetIdx < len(st.Fields) {
		a := AlignOf(st.Fields[targetIdx])
		if off%a != 0 { off += a - (off % a) }
	}
	return off
}