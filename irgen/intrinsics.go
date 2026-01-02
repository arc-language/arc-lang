package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

// GenerateIntrinsicCall handles function calls that are actually compiler intrinsics
func (g *Generator) GenerateIntrinsicCall(name string, args []ir.Value) ir.Value {
	switch name {
	case "alloca":
		// alloca(type) or alloca(type, count)
		// Since we passed type via generics or arguments in a real parser, 
		// here we simplify: argument 0 is likely the count if it exists, type is implicit or cast
		// For the 'alloca' function signature in builtins, we treated it as returning *i8.
		// In a real compiler, we'd need to inspect the GenericArgs from the call node to get the Type.
		// Fallback: Default to I8 or I64 for raw buffers.
		typ := types.I8 
		
		if len(args) > 0 {
			// alloca(count) - though signature in builtins.go was func<T>(count)
			return g.ctx.Builder.CreateAllocaWithCount(typ, args[0], "")
		}
		return g.ctx.Builder.CreateAlloca(typ, "")

	case "memset":
		if len(args) == 3 {
			return g.ctx.Builder.CreateMemSet(args[0], args[1], args[2])
		}

	case "memcpy":
		if len(args) == 3 {
			return g.ctx.Builder.CreateMemCpy(args[0], args[1], args[2])
		}

	case "memmove":
		if len(args) == 3 {
			return g.ctx.Builder.CreateMemMove(args[0], args[1], args[2])
		}
		
	case "strlen":
		if len(args) == 1 {
			return g.ctx.Builder.CreateStrLen(args[0], "")
		}

	case "syscall":
		// syscall(num, args...)
		return g.ctx.Builder.CreateSyscall(args)

	case "raise":
		if len(args) == 1 {
			g.ctx.Builder.CreateRaise(args[0])
		}
		return g.getZeroValue(types.Void)
		
	case "bit_cast":
		// Logic usually requires knowing target type from generics
		return args[0] // No-op placeholder without generic info
	}

	return nil
}