package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// IRVisitor implements the ANTLR visitor pattern to generate IR
type IRVisitor struct {
	*parser.BaseArcParserVisitor
	compiler    *Compiler
	ctx         *Context
	currentFile string
	logger      *Logger
	
	// Method call tracking
	pendingMethodSelf ir.Value
}

// NewIRVisitor creates a new IR visitor
func NewIRVisitor(c *Compiler, filename string) *IRVisitor {
	logger := NewLogger("[Visitor]")
	logger.Debug("Created visitor for file: %s", filename)
	
	return &IRVisitor{
		BaseArcParserVisitor: &parser.BaseArcParserVisitor{},
		compiler:             c,
		ctx:                  c.context,
		currentFile:          filename,
		logger:               logger,
	}
}

// Visit overrides the base Visit to add explicit dispatching
func (v *IRVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}

	// Explicitly dispatch to the correct visitor method based on context type
	switch ctx := tree.(type) {
	case *parser.CompilationUnitContext:
		return v.VisitCompilationUnit(ctx)
	case *parser.TopLevelDeclContext:
		return v.VisitTopLevelDecl(ctx)
	case *parser.NamespaceDeclContext:
		return v.VisitNamespaceDecl(ctx)
	case *parser.ImportDeclContext:
		return v.VisitImportDecl(ctx)
	case *parser.ExternDeclContext:
		return v.VisitExternDecl(ctx)
	case *parser.ExternMemberContext:
		return v.VisitExternMember(ctx)
	case *parser.ExternFunctionDeclContext:
		return v.VisitExternFunctionDecl(ctx)
	case *parser.FunctionDeclContext:
		return v.VisitFunctionDecl(ctx)
	case *parser.StructDeclContext:
		return v.VisitStructDecl(ctx)
	case *parser.ClassDeclContext:
		return v.VisitClassDecl(ctx)
	case *parser.ClassMemberContext:
		return v.VisitClassMember(ctx)
	case *parser.ClassFieldContext:
		return v.VisitClassField(ctx)
	case *parser.DeinitDeclContext:
		return v.VisitDeinitDecl(ctx)
	case *parser.EnumDeclContext:
		return v.VisitEnumDecl(ctx)
	case *parser.EnumMemberContext:
		return v.VisitEnumMember(ctx)
	case *parser.BlockContext:
		return v.VisitBlock(ctx)
	case *parser.StatementContext:
		return v.VisitStatement(ctx)
	case *parser.VariableDeclContext:
		return v.VisitVariableDecl(ctx)
	case *parser.ConstDeclContext:
		return v.VisitConstDecl(ctx)
	case *parser.AssignmentStmtContext:
		return v.VisitAssignmentStmt(ctx)
	case *parser.ReturnStmtContext:
		return v.VisitReturnStmt(ctx)
	case *parser.IfStmtContext:
		return v.VisitIfStmt(ctx)
	case *parser.ForStmtContext:
		return v.VisitForStmt(ctx)
	case *parser.SwitchStmtContext:
		return v.VisitSwitchStmt(ctx)
	case *parser.TryStmtContext:
		return v.VisitTryStmt(ctx)
	case *parser.ThrowStmtContext:
		return v.VisitThrowStmt(ctx)
	case *parser.BreakStmtContext:
		return v.VisitBreakStmt(ctx)
	case *parser.ContinueStmtContext:
		return v.VisitContinueStmt(ctx)
	case *parser.DeferStmtContext:
		return v.VisitDeferStmt(ctx)
	case *parser.ExpressionStmtContext:
		return v.VisitExpressionStmt(ctx)
	case *parser.ExpressionContext:
		return v.VisitExpression(ctx)
	case *parser.LogicalOrExpressionContext:
		return v.VisitLogicalOrExpression(ctx)
	case *parser.LogicalAndExpressionContext:
		return v.VisitLogicalAndExpression(ctx)
	case *parser.BitOrExpressionContext:
		return v.VisitBitOrExpression(ctx)
	case *parser.BitXorExpressionContext:
		return v.VisitBitXorExpression(ctx)
	case *parser.BitAndExpressionContext:
		return v.VisitBitAndExpression(ctx)
	case *parser.ShiftExpressionContext:
		return v.VisitShiftExpression(ctx)
	case *parser.EqualityExpressionContext:
		return v.VisitEqualityExpression(ctx)
	case *parser.RelationalExpressionContext:
		return v.VisitRelationalExpression(ctx)
	case *parser.RangeExpressionContext:
		return v.VisitRangeExpression(ctx)
	case *parser.AdditiveExpressionContext:
		return v.VisitAdditiveExpression(ctx)
	case *parser.MultiplicativeExpressionContext:
		return v.VisitMultiplicativeExpression(ctx)
	case *parser.UnaryExpressionContext:
		return v.VisitUnaryExpression(ctx)
	case *parser.PostfixExpressionContext:
		return v.VisitPostfixExpression(ctx)
	case *parser.PrimaryExpressionContext:
		return v.VisitPrimaryExpression(ctx)
	case *parser.LiteralContext:
		return v.VisitLiteral(ctx)
	case *parser.StructLiteralContext:
		return v.VisitStructLiteral(ctx)
	case *parser.CastExpressionContext:
		return v.VisitCastExpression(ctx)
	case *parser.AllocaExpressionContext:
		return v.VisitAllocaExpression(ctx)
	case *parser.SyscallExpressionContext:
		return v.VisitSyscallExpression(ctx)
	case *parser.IntrinsicExpressionContext:
		return v.VisitIntrinsicExpression(ctx)
	case *parser.ArgumentListContext:
		return v.VisitArgumentList(ctx)
	case *parser.LeftHandSideContext:
		return v.VisitLeftHandSide(ctx)
	default:
		return v.BaseArcParserVisitor.Visit(tree)
	}
}

// ============================================================================
// COMPILATION UNIT & TOP LEVEL
// ============================================================================

func (v *IRVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	v.logger.Info("Starting compilation of %s", v.currentFile)
	
	// Pass 0: Imports
	v.logger.Debug("Pass 0 - Processing imports")
	for _, imp := range ctx.AllImportDecl() {
		v.Visit(imp)
	}

	// Process Namespace declaration if present
	for _, ns := range ctx.AllNamespaceDecl() {
		v.Visit(ns)
	}

	// Pass 1: Register all type declarations (structs, classes, and enums)
	v.logger.Debug("Pass 1 - Registering types")
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.StructDecl() != nil {
			v.registerStructType(decl.StructDecl().(*parser.StructDeclContext))
		} else if decl.ClassDecl() != nil {
			v.registerClassType(decl.ClassDecl().(*parser.ClassDeclContext))
		}
		// Note: Enums are processed in pass 2 since they create constants
	}
	
	// Pass 2: Process everything else
	v.logger.Debug("Pass 2 - Processing declarations")
	
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.FunctionDecl() != nil {
			v.Visit(decl.FunctionDecl())
		} else if decl.ExternDecl() != nil {
			v.Visit(decl.ExternDecl())
		} else if decl.ConstDecl() != nil {
			v.Visit(decl.ConstDecl())
		} else if decl.VariableDecl() != nil {
			v.Visit(decl.VariableDecl())
		} else if decl.StructDecl() != nil {
			v.Visit(decl.StructDecl())
		} else if decl.ClassDecl() != nil {
			v.Visit(decl.ClassDecl())
		} else if decl.EnumDecl() != nil {
			v.Visit(decl.EnumDecl())
		}
	}
	
	v.logger.Info("Compilation complete for %s", v.currentFile)
	return nil
}

func (v *IRVisitor) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil {
		return v.Visit(ctx.FunctionDecl())
	}
	if ctx.StructDecl() != nil {
		return v.Visit(ctx.StructDecl())
	}
	if ctx.ClassDecl() != nil {
		return v.Visit(ctx.ClassDecl())
	}
	if ctx.EnumDecl() != nil {
		return v.Visit(ctx.EnumDecl())
	}
	if ctx.ExternDecl() != nil {
		return v.Visit(ctx.ExternDecl())
	}
	if ctx.ConstDecl() != nil {
		return v.Visit(ctx.ConstDecl())
	}
	if ctx.VariableDecl() != nil {
		return v.Visit(ctx.VariableDecl())
	}
	return nil
}

func (v *IRVisitor) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	v.logger.Info("Setting current namespace to '%s'", name)
	v.ctx.SetNamespace(name)
	return nil
}

// ============================================================================
// HELPERS
// ============================================================================

func (v *IRVisitor) resolveType(ctx parser.ITypeContext) types.Type {
	if ctx == nil {
		return types.Void
	}
	
	typeCtx := ctx.(*parser.TypeContext)
	
	if typeCtx.PrimitiveType() != nil {
		name := typeCtx.PrimitiveType().GetText()
		if typ, ok := v.ctx.GetType(name); ok {
			return typ
		}
		v.logger.Warning("Unknown primitive type '%s', defaulting to i64", name)
		return types.I64
	}
	
	if typeCtx.PointerType() != nil {
		elemType := v.resolveType(typeCtx.PointerType().Type_())
		return types.NewPointer(elemType)
	}
	
	if typeCtx.ReferenceType() != nil {
		elemType := v.resolveType(typeCtx.ReferenceType().Type_())
		return types.NewPointer(elemType)
	}
	
	if typeCtx.ArrayType() != nil {
		arrCtx := typeCtx.ArrayType()
		elemType := v.resolveType(arrCtx.Type_())
		
		// Get array size
		var size int64 = 0
		if arrCtx.ArraySize() != nil {
			sizeCtx := arrCtx.ArraySize()
			if sizeCtx.INTEGER_LITERAL() != nil {
				// Parse integer literal
				sizeText := sizeCtx.INTEGER_LITERAL().GetText()
				var err error
				size, err = parseInt(sizeText)
				if err != nil {
					v.ctx.Logger.Error("Invalid array size: %s", sizeText)
					size = 0
				}
			} else if sizeCtx.IDENTIFIER() != nil {
				// Constant identifier
				name := sizeCtx.IDENTIFIER().GetText()
				if sym, ok := v.ctx.currentScope.Lookup(name); ok {
					if constInt, ok := sym.Value.(*ir.ConstantInt); ok {
						size = constInt.Value
					}
				}
			}
		}
		
		return types.NewArray(elemType, size)
	}

	// Handle Qualified Type (namespace.Type)
	if typeCtx.QualifiedType() != nil {
		qCtx := typeCtx.QualifiedType()
		parts := make([]string, len(qCtx.AllIDENTIFIER()))
		for i, node := range qCtx.AllIDENTIFIER() {
			parts[i] = node.GetText()
		}
		
		if len(parts) >= 2 {
			nsName := parts[0]
			typeName := parts[1] // For now support 1 level deep: namespace.Type
			
			if ns, ok := v.ctx.NamespaceRegistry[nsName]; ok {
				if typ, ok := ns.Types[typeName]; ok {
					return typ
				}
				v.ctx.Logger.Error("Type '%s' not found in namespace '%s'", typeName, nsName)
				return types.I64
			}
			v.ctx.Logger.Error("Unknown namespace: %s", nsName)
			return types.I64
		}
	}
	
	if typeCtx.IDENTIFIER() != nil {
		name := typeCtx.IDENTIFIER().GetText()
		if typ, ok := v.ctx.GetType(name); ok {
			return typ
		}
		v.ctx.Logger.Error("Unknown type: %s", name)
		return types.I64
	}
	
	return types.I64
}

func (v *IRVisitor) getZeroValue(typ types.Type) ir.Value {
	switch typ.Kind() {
	case types.IntegerKind:
		return v.ctx.Builder.ConstInt(typ.(*types.IntType), 0)
	case types.FloatKind:
		return v.ctx.Builder.ConstFloat(typ.(*types.FloatType), 0.0)
	case types.PointerKind:
		return v.ctx.Builder.ConstNull(typ.(*types.PointerType))
	case types.ArrayKind:
		return v.ctx.Builder.ConstZero(typ)
	case types.StructKind:
		return v.ctx.Builder.ConstZero(typ)
	default:
		return v.ctx.Builder.ConstZero(typ)
	}
}

func (v *IRVisitor) findFieldIndex(structType *types.StructType, fieldName string) int {
	if fieldIndices, ok := v.ctx.StructFieldIndices[structType.Name]; ok {
		if idx, ok := fieldIndices[fieldName]; ok {
			return idx
		}
	}
	return -1
}

func (v *IRVisitor) castValue(val ir.Value, targetType types.Type) ir.Value {
	srcType := val.Type()
	
	if types.IsInteger(srcType) && types.IsInteger(targetType) {
		srcBits := srcType.(*types.IntType).BitWidth
		destBits := targetType.(*types.IntType).BitWidth
		if srcBits > destBits {
			return v.ctx.Builder.CreateTrunc(val, targetType, "")
		} else if srcBits < destBits {
			return v.ctx.Builder.CreateSExt(val, targetType, "")
		}
	}
	
	if types.IsFloat(srcType) && types.IsFloat(targetType) {
		srcBits := srcType.(*types.FloatType).BitWidth
		destBits := targetType.(*types.FloatType).BitWidth
		if srcBits > destBits {
			return v.ctx.Builder.CreateFPTrunc(val, targetType, "")
		} else if srcBits < destBits {
			return v.ctx.Builder.CreateFPExt(val, targetType, "")
		}
	}
	
	return val
}

func (v *IRVisitor) castConstant(constant ir.Constant, targetType types.Type) ir.Constant {
	srcType := constant.Type()
	
	if srcType.Equal(targetType) {
		return constant
	}
	
	if srcInt, ok := constant.(*ir.ConstantInt); ok {
		if targetInt, ok := targetType.(*types.IntType); ok {
			return v.ctx.Builder.ConstInt(targetInt, srcInt.Value)
		}
	}
	
	v.logger.Warning("Cannot cast constant from %v to %v", srcType, targetType)
	return constant
}

// Helper function to parse integer literals
func parseInt(s string) (int64, error) {
	var base int = 10
	
	// Handle hex (0x), octal (0o), binary (0b) prefixes
	if len(s) > 2 {
		switch {
		case s[0:2] == "0x" || s[0:2] == "0X":
			base = 16
			s = s[2:]
		case s[0:2] == "0o" || s[0:2] == "0O":
			base = 8
			s = s[2:]
		case s[0:2] == "0b" || s[0:2] == "0B":
			base = 2
			s = s[2:]
		}
	}
	
	var result int64 = 0
	for _, ch := range s {
		var digit int64
		switch {
		case ch >= '0' && ch <= '9':
			digit = int64(ch - '0')
		case ch >= 'a' && ch <= 'f':
			digit = int64(ch - 'a' + 10)
		case ch >= 'A' && ch <= 'F':
			digit = int64(ch - 'A' + 10)
		case ch == '_':
			continue // Allow underscores as separators
		default:
			return 0, nil
		}
		result = result*int64(base) + digit
	}
	
	return result, nil
}