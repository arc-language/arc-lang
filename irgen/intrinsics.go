package irgen

import (

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (g *Generator) GenerateIntrinsicCall(name string, args []ir.Value, typeArgs []types.Type) ir.Value {
	switch name {
	case "alloca":
		// alloca<T>(count)
		var typ types.Type = types.I8
		if len(typeArgs) > 0 {
			typ = typeArgs[0]
		}
		
		if len(args) > 0 {
			return g.ctx.Builder.CreateAllocaWithCount(typ, args[0], "")
		}
		return g.ctx.Builder.CreateAlloca(typ, "")

	case "cast":
		// cast<T>(val) - Value conversion (truncation, extension, float<->int value)
		if len(args) > 0 {
			if len(typeArgs) > 0 {
				return g.emitCast(args[0], typeArgs[0])
			}
			return args[0]
		}

	case "bit_cast":
		// bit_cast<T>(val) - Bitwise reinterpretation (e.g., float bits as uint32)
		if len(args) > 0 {
			if len(typeArgs) > 0 {
				return g.ctx.Builder.CreateBitCast(args[0], typeArgs[0], "")
			}
			return args[0]
		}

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
		
	// --- Variadic Argument Intrinsics ---
	
	case "va_start":
		if len(args) == 1 {
			return g.ctx.Builder.CreateVaStart(args[0])
		}
		
	case "va_end":
		if len(args) == 1 {
			return g.ctx.Builder.CreateVaEnd(args[0])
		}
		
	case "va_arg":
		// va_arg<T>(list)
		if len(args) == 1 && len(typeArgs) > 0 {
			return g.ctx.Builder.CreateVaArg(args[0], typeArgs[0], "")
		}
	}

	return nil
}