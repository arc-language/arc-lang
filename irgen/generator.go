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

// VisitFunctionDecl handles function declarations.
// It includes logic to detect the <gpu> and <tpu> generic tags to set specific calling conventions.
func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	// 1. Detect Hardware Markers in Generics
	// We scan the generic parameter list for specific tags like "gpu" or "tpu".
	isGPU := false
	isTPU := false

	if gp := ctx.GenericParams(); gp != nil {
		if gpl := gp.GenericParamList(); gpl != nil {
			for _, id := range gpl.AllIDENTIFIER() {
				tag := id.GetText()
				if tag == "gpu" {
					isGPU = true
				} else if tag == "tpu" {
					isTPU = true
				}
			}
		}
	}

	name := ctx.IDENTIFIER().GetText()
	irName := name

	// 2. Handle Name Mangling (Methods / Namespaces)
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

	// Construct the actual IR name
	if isMethod {
		irName = parentName + "_" + name
	} else if g.currentNamespace != "" && name != "main" {
		irName = g.currentNamespace + "_" + name
	}

	// Construct the lookup name for the symbol table
	lookupName := name
	if isMethod {
		lookupName = parentName + "_" + name
	} else if g.currentNamespace != "" && name != "main" {
		lookupName = g.currentNamespace + "." + name
	}

	sym, _ := g.currentScope.Resolve(lookupName)

	// --- Phase 1: Create Function Prototype ---
	// In this phase, we define the function signature and add it to the module.
	// We do not generate the body yet.
	if g.Phase == 1 {
		var retType types.Type = types.Void
		if ctx.ReturnType() != nil {
			if ctx.ReturnType().Type_() != nil {
				retType = g.resolveType(ctx.ReturnType().Type_())
			}
		}
		var paramTypes []types.Type
		if ctx.ParameterList() != nil {
			for _, param := range ctx.ParameterList().AllParameter() {
				paramTypes = append(paramTypes, g.resolveType(param.Type_()))
			}
		}

		// Create the function in the IR module
		fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)

		// Apply specific calling conventions based on tags found earlier
		if isGPU {
			fn.CallConv = ir.CC_PTX
		} else if isTPU {
			fn.CallConv = ir.CC_TPU
		}

		// Link the IR function to the semantic symbol
		if sym != nil {
			sym.IRValue = fn
		}
		return nil
	}

	// --- Phase 2: Generate Body ---
	// In this phase, we generate the instructions inside the function.
	if g.Phase == 2 {
		if sym == nil || sym.IRValue == nil {
			return nil
		}

		fn := sym.IRValue.(*ir.Function)
		g.ctx.EnterFunction(fn)
		g.enterScope(ctx)
		defer g.exitScope()

		// 3. Setup Arguments
		// We create stack allocations (alloca) for parameters so they are mutable l-values.
		var paramNames []string
		if ctx.ParameterList() != nil {
			for _, param := range ctx.ParameterList().AllParameter() {
				paramNames = append(paramNames, param.IDENTIFIER().GetText())
			}
		}

		for i, arg := range fn.Arguments {
			if i < len(paramNames) {
				arg.SetName(paramNames[i])
				
				// Create a stack slot for the argument
				alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
				
				// Store the incoming argument value into the stack slot
				g.ctx.Builder.CreateStore(arg, alloca)
				
				// Update the symbol table so variable lookups find the alloca (l-value)
				if s, ok := g.currentScope.Resolve(paramNames[i]); ok {
					s.IRValue = alloca
				}
			}
		}

		// 4. Generate Block Instructions
		if ctx.Block() != nil {
			// Reset defer stack for the new function scope
			g.deferStack = NewDeferStack()
			g.Visit(ctx.Block())
		}

		// 5. Implicit Return
		// If the last block doesn't have a terminator (ret, br), insert a return.
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			// Run any deferred statements before the implicit return
			g.deferStack.Emit(g)
			
			if fn.FuncType.ReturnType == types.Void {
				g.ctx.Builder.CreateRetVoid()
			} else {
				// Safety: return a zero value for non-void functions missing a return
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