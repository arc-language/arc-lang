package irgen

import (

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/context"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
)

// SourceUnit represents a single parsed source file.
// Defined here to be shared between the Driver and IRGen.
type SourceUnit struct {
	Path string
	Tree parser.ICompilationUnitContext
}

// Generator orchestrates the entire compilation process from AST to IR.
type Generator struct {
	*parser.BaseArcParserVisitor
	ctx              *context.Context
	analysis         *semantics.AnalysisResult
	currentScope     *symbol.Scope
	deferStack       *DeferStack
	loopStack        []loopInfo
	currentNamespace string
	Phase            int // 1=Prototypes, 2=Bodies
}

type loopInfo struct {
	breakBlock    *ir.BasicBlock
	continueBlock *ir.BasicBlock
}

// GenerateProject compiles parsed trees into an IR Module.
func GenerateProject(units []*SourceUnit, moduleName string, analysis *semantics.AnalysisResult) *ir.Module {
	ctx := context.NewContext(moduleName)
	gen := &Generator{
		BaseArcParserVisitor: &parser.BaseArcParserVisitor{},
		ctx:                  ctx,
		analysis:             analysis,
		currentScope:         analysis.GlobalScope,
		deferStack:           NewDeferStack(),
	}

	// Pass 1: Generate Prototypes (Function definitions, Globals, Structs)
	gen.Phase = 1
	for _, unit := range units {
		gen.Visit(unit.Tree)
	}

	// Pass 2: Generate Bodies (Instructions)
	gen.Phase = 2
	for _, unit := range units {
		gen.Visit(unit.Tree)
	}

	return ctx.Module
}

func (g *Generator) enterScope(ctx antlr.ParserRuleContext) {
	if s, ok := g.analysis.Scopes[ctx]; ok {
		g.currentScope = s
	}
}

func (g *Generator) exitScope() {
	if g.currentScope.Parent != nil {
		g.currentScope = g.currentScope.Parent
	}
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	// 1. Hardware Markers (Generics)
	isGPU := false
	isROCm := false
	isCUDA := false
	isTPU := false

	if gp := ctx.GenericParams(); gp != nil {
		if gpl := gp.GenericParamList(); gpl != nil {
			for _, param := range gpl.AllGenericParam() {
				for _, id := range param.AllIDENTIFIER() {
					tag := id.GetText()
					if tag == "gpu" { 
						isGPU = true 
					} else if tag == "rocm" { 
						isROCm = true 
					} else if tag == "cuda" { 
						isCUDA = true 
					} else if tag == "tpu" { 
						isTPU = true 
					}
				}
			}
		}
	}

	// 2. Resolve Name
	// Get all identifiers. The function name is the LAST one.
	ids := ctx.AllIDENTIFIER()
	nameToken := ids[len(ids)-1]
	name := nameToken.GetText()
	
	irName := name

	var parentName string
	isMethod := false
	if parent := ctx.GetParent(); parent != nil {
		if _, ok := parent.(*parser.ClassMemberContext); ok {
			if classDecl, ok := parent.GetParent().(*parser.ClassDeclContext); ok {
				parentName = classDecl.IDENTIFIER().GetText()
				isMethod = true
			}
		} else if _, ok := parent.(*parser.StructMemberContext); ok {
			if structDecl, ok := parent.GetParent().(*parser.StructDeclContext); ok {
				parentName = structDecl.IDENTIFIER().GetText()
				isMethod = true
			}
		}
	}

	// Handle Flat Methods (func foo(self x: T))
	if !isMethod && ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil {
				t := g.resolveType(param.Type_())
				if ptr, ok := t.(*types.PointerType); ok {
					t = ptr.ElementType
				}
				if st, ok := t.(*types.StructType); ok {
					parentName = st.Name
					isMethod = true
				}
				break
			}
		}
	}

	if isMethod {
		irName = parentName + "_" + name
	} else if g.currentNamespace != "" && name != "main" {
		irName = g.currentNamespace + "_" + name
	}

	lookupName := name
	if isMethod {
		lookupName = parentName + "_" + name
	} else if g.currentNamespace != "" && name != "main" {
		lookupName = g.currentNamespace + "." + name
	}

	sym, _ := g.currentScope.Resolve(lookupName)

	// --- Phase 1: Prototype ---
	if g.Phase == 1 {
		var retType types.Type = types.Void
		if ctx.ReturnType() != nil {
			if ctx.ReturnType().Type_() != nil {
				retType = g.resolveType(ctx.ReturnType().Type_())
			} else if ctx.ReturnType().TypeList() != nil {
				var tupleTypes []types.Type
				for _, t := range ctx.ReturnType().TypeList().AllType_() {
					tupleTypes = append(tupleTypes, g.resolveType(t))
				}
				retType = types.NewStruct("", tupleTypes, false)
			}
		}
		var paramTypes []types.Type
		if ctx.ParameterList() != nil {
			for _, param := range ctx.ParameterList().AllParameter() {
				paramTypes = append(paramTypes, g.resolveType(param.Type_()))
			}
		}

		fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)

		// Set Hardware CallConvs
		if isTPU {
			fn.CallConv = ir.CC_TPU
		} else if isROCm {
			fn.CallConv = ir.CC_ROCM
		} else if isCUDA {
			fn.CallConv = ir.CC_PTX
		} else if isGPU {
			fn.CallConv = ir.CC_PTX
		}

		// Set Concurrency Flags (String Detection)
		if ctx.ASYNC() != nil {
			fn.FuncType.IsAsync = true
		} else if len(ids) > 1 && ids[0].GetText() == "process" {
			fn.FuncType.IsProcess = true
		}

		if sym != nil {
			sym.IRValue = fn
		}
		return nil
	}

	// --- Phase 2: Body ---
	if g.Phase == 2 {
		if sym == nil || sym.IRValue == nil {
			return nil
		}

		fn := sym.IRValue.(*ir.Function)
		g.ctx.EnterFunction(fn)
		g.enterScope(ctx)
		defer g.exitScope()

		var paramNames []string
		if ctx.ParameterList() != nil {
			for _, param := range ctx.ParameterList().AllParameter() {
				paramNames = append(paramNames, param.IDENTIFIER().GetText())
			}
		}

		for i, arg := range fn.Arguments {
			if i < len(paramNames) {
				arg.SetName(paramNames[i])
				alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
				g.ctx.Builder.CreateStore(arg, alloca)
				if s, ok := g.currentScope.Resolve(paramNames[i]); ok {
					s.IRValue = alloca
				}
			}
		}

		if ctx.Block() != nil {
			g.deferStack = NewDeferStack()
			g.Visit(ctx.Block())
		}

		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.deferStack.Emit(g)
			if fn.FuncType.ReturnType == types.Void {
				g.ctx.Builder.CreateRetVoid()
			} else {
				g.ctx.Builder.CreateRet(g.getZeroValue(fn.FuncType.ReturnType))
			}
		}
		
		g.ctx.ExitFunction()
	}
	return nil
}

// Visit acts as the main dispatcher for the AST nodes.
func (g *Generator) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}

	switch ctx := tree.(type) {
	case *parser.CompilationUnitContext:
		return g.VisitCompilationUnit(ctx)
	case *parser.TopLevelDeclContext:
		return g.VisitTopLevelDecl(ctx)
	case *parser.FunctionDeclContext:
		return g.VisitFunctionDecl(ctx)
	case *parser.MutatingDeclContext:
		return g.VisitMutatingDecl(ctx)
	case *parser.VariableDeclContext:
		return g.VisitVariableDecl(ctx)
	case *parser.ExternCDeclContext:
		return g.VisitExternCDecl(ctx)
	case *parser.ExternCppDeclContext:
		return g.VisitExternCppDecl(ctx)
	case *parser.StructDeclContext:
		return g.VisitStructDecl(ctx)
	case *parser.ClassDeclContext:
		return g.VisitClassDecl(ctx)
	case *parser.EnumDeclContext:
		return g.VisitEnumDecl(ctx)
	case *parser.ConstDeclContext:
		return g.VisitConstDecl(ctx)
	case *parser.BlockContext:
		return g.VisitBlock(ctx)
	case *parser.NamespaceDeclContext:
		return g.VisitNamespaceDecl(ctx)

	case *parser.StatementContext:
		if ctx.VariableDecl() != nil {
			return g.VisitVariableDecl(ctx.VariableDecl().(*parser.VariableDeclContext))
		}
		if ctx.ConstDecl() != nil {
			return g.VisitConstDecl(ctx.ConstDecl().(*parser.ConstDeclContext))
		}
		if ctx.ReturnStmt() != nil {
			return g.VisitReturnStmt(ctx.ReturnStmt().(*parser.ReturnStmtContext))
		}
		if ctx.IfStmt() != nil {
			return g.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
		}
		if ctx.ForStmt() != nil {
			return g.VisitForStmt(ctx.ForStmt().(*parser.ForStmtContext))
		}
		if ctx.SwitchStmt() != nil {
			return g.VisitSwitchStmt(ctx.SwitchStmt().(*parser.SwitchStmtContext))
		}
		if ctx.BreakStmt() != nil {
			return g.VisitBreakStmt(ctx.BreakStmt().(*parser.BreakStmtContext))
		}
		if ctx.ContinueStmt() != nil {
			return g.VisitContinueStmt(ctx.ContinueStmt().(*parser.ContinueStmtContext))
		}
		if ctx.DeferStmt() != nil {
			return g.VisitDeferStmt(ctx.DeferStmt().(*parser.DeferStmtContext))
		}
		if ctx.AssignmentStmt() != nil {
			return g.VisitAssignmentStmt(ctx.AssignmentStmt().(*parser.AssignmentStmtContext))
		}
		if ctx.ExpressionStmt() != nil {
			return g.Visit(ctx.ExpressionStmt().Expression())
		}
		if ctx.Block() != nil {
			return g.VisitBlock(ctx.Block().(*parser.BlockContext))
		}
		if ctx.TryStmt() != nil {
			return g.VisitTryStmt(ctx.TryStmt().(*parser.TryStmtContext))
		}
		if ctx.ThrowStmt() != nil {
			return g.VisitThrowStmt(ctx.ThrowStmt().(*parser.ThrowStmtContext))
		}
		return nil

	case *parser.ReturnStmtContext:
		return g.VisitReturnStmt(ctx)
	case *parser.IfStmtContext:
		return g.VisitIfStmt(ctx)
	case *parser.ForStmtContext:
		return g.VisitForStmt(ctx)
	case *parser.SwitchStmtContext:
		return g.VisitSwitchStmt(ctx)
	case *parser.BreakStmtContext:
		return g.VisitBreakStmt(ctx)
	case *parser.ContinueStmtContext:
		return g.VisitContinueStmt(ctx)
	case *parser.DeferStmtContext:
		return g.VisitDeferStmt(ctx)
	case *parser.AssignmentStmtContext:
		return g.VisitAssignmentStmt(ctx)
	case *parser.TryStmtContext:
		return g.VisitTryStmt(ctx)
	case *parser.ThrowStmtContext:
		return g.VisitThrowStmt(ctx)

	case *parser.ExpressionContext:
		return g.VisitExpression(ctx)
	case *parser.LogicalOrExpressionContext:
		return g.VisitLogicalOrExpression(ctx)
	case *parser.LogicalAndExpressionContext:
		return g.VisitLogicalAndExpression(ctx)
	case *parser.BitOrExpressionContext:
		return g.VisitBitOrExpression(ctx)
	case *parser.BitXorExpressionContext:
		return g.VisitBitXorExpression(ctx)
	case *parser.BitAndExpressionContext:
		return g.VisitBitAndExpression(ctx)
	case *parser.EqualityExpressionContext:
		return g.VisitEqualityExpression(ctx)
	case *parser.RelationalExpressionContext:
		return g.VisitRelationalExpression(ctx)
	case *parser.ShiftExpressionContext:
		return g.VisitShiftExpression(ctx)
	case *parser.RangeExpressionContext:
		return g.VisitRangeExpression(ctx)
	case *parser.AdditiveExpressionContext:
		return g.VisitAdditiveExpression(ctx)
	case *parser.MultiplicativeExpressionContext:
		return g.VisitMultiplicativeExpression(ctx)
	case *parser.UnaryExpressionContext:
		return g.VisitUnaryExpression(ctx)
	case *parser.PostfixExpressionContext:
		return g.VisitPostfixExpression(ctx)
	case *parser.PrimaryExpressionContext:
		return g.VisitPrimaryExpression(ctx)

	case *parser.LiteralContext:
		return g.VisitLiteral(ctx)
	case *parser.StructLiteralContext:
		return g.VisitStructLiteral(ctx)

	default:
		return tree.Accept(g)
	}
}