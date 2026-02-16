package codegen

import (
	"fmt"
	"github.com/arc-language/arc-lang/ast"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (cg *Codegen) genExpr(expr ast.Expr) ir.Value {
	switch e := expr.(type) {
	
	case *ast.BasicLit:
		if e.Kind == "INT" {
			// In a real compiler, we'd parse e.Value to int64
			// Placeholder: assuming value is 0 for example purposes
			return cg.Builder.ConstInt(types.I32, 0)
		}
		if e.Kind == "STRING" {
			// Returns *i8
			return cg.createGlobalString(e.Value)
		}

	case *ast.Ident:
		// 1. Lookup the stack slot (Alloca)
		alloca := cg.lookupVar(e.Name)
		if alloca == nil {
			panic(fmt.Sprintf("undefined var %s", e.Name))
		}
		// 2. Load the value
		// alloca type is *T, we want T
		ptrType := alloca.Type().(*types.PointerType)
		return cg.Builder.CreateLoad(ptrType.ElementType, alloca, e.Name)

	case *ast.BinaryExpr:
		lhs := cg.genExpr(e.Left)
		rhs := cg.genExpr(e.Right)
		
		switch e.Op {
		case "+": return cg.Builder.CreateAdd(lhs, rhs, "")
		case "-": return cg.Builder.CreateSub(lhs, rhs, "")
		case "*": return cg.Builder.CreateMul(lhs, rhs, "")
		case "/": return cg.Builder.CreateSDiv(lhs, rhs, "")
		case "==": return cg.Builder.CreateICmpEQ(lhs, rhs, "")
		}

	case *ast.CallExpr:
		// 1. Resolve Function
		var callee *ir.Function
		var fnName string
		
		if id, ok := e.Fun.(*ast.Ident); ok {
			fnName = id.Name
			callee = cg.Module.GetFunction(fnName)
		}

		// 2. Generate Arguments
		var args []ir.Value
		for _, arg := range e.Args {
			args = append(args, cg.genExpr(arg))
		}

		// 3. Handle Intrinsics from Lowering
		// "async_spawn" -> ir.AsyncTaskCreate
		if fnName == "async_spawn" {
			// Args: [FuncPtr, ArgStructPtr]
			return cg.Builder.CreateAsyncTask(callee, args, "handle")
		}

		if callee == nil {
			panic("indirect calls not implemented in this snippet")
		}

		return cg.Builder.CreateCall(callee, args, "")
	}

	return nil
}