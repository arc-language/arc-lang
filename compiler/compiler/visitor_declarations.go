// --- START OF FILE compiler/visitor_declarations.go ---
package compiler

import (
	"path/filepath"
	"strings"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// ============================================================================
// IMPORT DECLARATIONS
// ============================================================================

func (v *IRVisitor) VisitImportDecl(ctx *parser.ImportDeclContext) interface{} {
	// 1. Get the import path string (remove quotes)
	rawPath := ctx.STRING_LITERAL().GetText()
	importPath := strings.Trim(rawPath, "\"")

	v.logger.Info("Processing import: %s from %s", importPath, v.currentFile)

	// 2. Resolve absolute directory path
	currentDir := filepath.Dir(v.currentFile)
	absPath, err := v.ctx.Importer.ResolvePath(currentDir, importPath)
	if err != nil {
		v.ctx.Logger.Error("Import resolution failed for '%s': %v", importPath, err)
		return nil
	}

	// 3. Compile that package (recursively)
	pkgInfo, err := v.compiler.CompilePackage(absPath) 
	if err != nil {
		v.ctx.Logger.Error("Failed to compile package '%s': %v", importPath, err)
		return nil
	}

	v.logger.Info("Successfully imported package '%s' (namespace: %s)", importPath, pkgInfo.Name)
	return nil
}

// ============================================================================
// EXTERN DECLARATIONS
// ============================================================================

func (v *IRVisitor) VisitExternDecl(ctx *parser.ExternDeclContext) interface{} {
	var namespaceName string
	
	if ctx.IDENTIFIER() != nil {
		namespaceName = ctx.IDENTIFIER().GetText()
		v.logger.Debug("Processing extern namespace: %s", namespaceName)
		
		// Temporarily switch namespace for these externs
		oldNamespace := v.ctx.currentNamespace
		v.ctx.SetNamespace(namespaceName)
		
		// Process all extern members
		for _, member := range ctx.AllExternMember() {
			v.Visit(member)
		}
		
		// Restore namespace
		v.ctx.currentNamespace = oldNamespace
	} else {
		// No namespace, just global extern declarations
		for _, member := range ctx.AllExternMember() {
			v.Visit(member)
		}
	}
	
	return nil
}

func (v *IRVisitor) VisitExternMember(ctx *parser.ExternMemberContext) interface{} {
	if ctx.ExternFunctionDecl() != nil {
		return v.Visit(ctx.ExternFunctionDecl())
	}
	return nil
}

func (v *IRVisitor) VisitExternFunctionDecl(ctx *parser.ExternFunctionDeclContext) interface{} {
	if ctx.IDENTIFIER() == nil {
		return nil
	}
	name := ctx.IDENTIFIER().GetText()
	
	// Check for explicit external name (e.g. func printf "printf")
	var externalName string
	if ctx.STRING_LITERAL() != nil {
		rawName := ctx.STRING_LITERAL().GetText()
		externalName = strings.Trim(rawName, "\"")
	} else {
		externalName = name
	}
	
	var retType types.Type = types.Void
	if ctx.Type_() != nil {
		retType = v.resolveType(ctx.Type_())
	}
	
	paramTypes := make([]types.Type, 0)
	variadic := false
	
	if ctx.ExternParameterList() != nil {
		paramCtx := ctx.ExternParameterList()
		if paramCtx.ELLIPSIS() != nil {
			variadic = true
		}
		for _, typeCtx := range paramCtx.AllType_() {
			paramTypes = append(paramTypes, v.resolveType(typeCtx))
		}
	}
	
	// Use external name for the actual function declaration
	fn := v.ctx.Builder.DeclareFunction(externalName, retType, paramTypes, variadic)
	
	// Register in current namespace using the internal name
	if v.ctx.currentNamespace != nil {
		v.ctx.currentNamespace.Functions[name] = fn
		v.logger.Debug("Declared extern function '%s' (external: '%s') in namespace '%s'", 
			name, externalName, v.ctx.currentNamespace.Name)
	} else {
		v.ctx.currentScope.Define(name, fn)
		v.logger.Debug("Declared extern function '%s' (external: '%s') in global scope", 
			name, externalName)
	}
	
	return nil
}

// ============================================================================
// FUNCTION DECLARATIONS
// ============================================================================

func (v *IRVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	if ctx.IDENTIFIER() == nil {
		return nil
	}
	name := ctx.IDENTIFIER().GetText()
	
	// Check for Generic Definition
	if ctx.GenericParams() != nil {
		// If we are NOT in an instantiation context (overrideFunctionName is empty),
		// this is the generic definition. Store it and skip code generation.
		if v.overrideFunctionName == "" {
			v.ctx.GenericFunctionDecls[name] = ctx
			v.logger.Info("Registered generic function definition: %s", name)
			return nil
		}
		// If overrideFunctionName is set, we proceed to generate code for this instantiation.
		// The generic types T, U, etc., are already mapped in v.ctx.CurrentTypeParams.
	}

	// Check if this is an async function
	isAsync := ctx.ASYNC() != nil
	
	// Check if this is a method inside a class/struct
	var methodPrefix string
	if parent := ctx.GetParent(); parent != nil {
		if classMember, ok := parent.(*parser.ClassMemberContext); ok {
			if classDecl, ok := classMember.GetParent().(*parser.ClassDeclContext); ok {
				className := classDecl.IDENTIFIER().GetText()
				methodPrefix = className + "_"
				name = methodPrefix + name
			}
		} else if structMember, ok := parent.(*parser.StructMemberContext); ok {
			if structDecl, ok := structMember.GetParent().(*parser.StructDeclContext); ok {
				structName := structDecl.IDENTIFIER().GetText()
				// If we are instantiating a generic struct, use the mangled struct name override
				if v.overrideStructName != "" {
					structName = v.overrideStructName
				}
				methodPrefix = structName + "_"
				name = methodPrefix + name
			}
		}
	}
	
	// Handle Namespacing (if not a method, and not main)
	var irName string
	if v.overrideFunctionName != "" {
		irName = v.overrideFunctionName
	} else {
		irName = name
		// Special Case: main function should not be mangled
		isMain := name == "main" && (v.ctx.currentNamespace == nil || v.ctx.currentNamespace.Name == "main" || v.ctx.currentNamespace.Name == "")
		if !isMain && v.ctx.currentNamespace != nil && v.ctx.currentNamespace.Name != "" {
			irName = v.ctx.currentNamespace.Name + "_" + name
		}
	}

	if isAsync {
		v.logger.Debug("Declaring async function: %s (IR: %s)", name, irName)
	} else {
		v.logger.Debug("Declaring function: %s (IR: %s)", name, irName)
	}
	
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		retTypeCtx := ctx.ReturnType()
		if retTypeCtx.Type_() != nil {
			retType = v.resolveType(retTypeCtx.Type_())
		} else if retTypeCtx.TypeList() != nil {
			// Tuple return type
			typeListCtx := retTypeCtx.TypeList()
			tupleTypes := make([]types.Type, 0)
			for _, t := range typeListCtx.AllType_() {
				tupleTypes = append(tupleTypes, v.resolveType(t))
			}
			retType = types.NewStruct("", tupleTypes, false)
		}
	}
	
	paramTypes := make([]types.Type, 0)
	paramNames := make([]string, 0)
	variadic := false
	
	if ctx.ParameterList() != nil {
		paramCtx := ctx.ParameterList()
		if paramCtx.ELLIPSIS() != nil {
			variadic = true
		}
		for _, param := range paramCtx.AllParameter() {
			if param.SELF() != nil {
				// Self parameter for methods
				paramName := "self"
				if param.IDENTIFIER() != nil {
					paramName = param.IDENTIFIER().GetText()
				}
				paramType := v.resolveType(param.Type_())
				paramNames = append(paramNames, paramName)
				paramTypes = append(paramTypes, paramType)
			} else {
				paramName := param.IDENTIFIER().GetText()
				paramType := v.resolveType(param.Type_())
				paramNames = append(paramNames, paramName)
				paramTypes = append(paramTypes, paramType)
			}
		}
	}

	fn := v.ctx.Builder.CreateFunction(irName, retType, paramTypes, variadic)
	
	// Mark as coroutine if async
	if isAsync {
		fn.Attributes = append(fn.Attributes, ir.AttrCoroutine)
		v.logger.Info("Marked function '%s' as async/coroutine", irName)
	}
	
	// Register function in the current namespace (if not instantiating with override)
	// If instantiating, it's cached in InstantiatedFunctions by the caller
	if v.ctx.currentNamespace != nil && v.overrideFunctionName == "" {
		v.ctx.currentNamespace.Functions[name] = fn
	}

	for i, paramName := range paramNames {
		fn.Arguments[i].SetName(paramName)
	}
	
	v.ctx.EnterFunction(fn)
	
	if ctx.Block() != nil {
		entry := v.ctx.Builder.CreateBlock("entry")
		v.ctx.SetInsertBlock(entry)
		
		// Allocate space for parameters and store them
		for i, arg := range fn.Arguments {
			alloc := v.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
			v.ctx.Builder.CreateStore(arg, alloc)
			v.ctx.currentScope.Define(paramNames[i], alloc)
		}
		
		// If async, emit coroutine setup
		if isAsync {
			v.emitAsyncFunctionPrologue(fn)
		}
		
		v.Visit(ctx.Block())
		
		// Add default return if needed
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			// If async, emit coroutine cleanup
			if isAsync {
				v.emitAsyncFunctionEpilogue(fn)
			}
			
			if retType.Kind() == types.VoidKind {
				v.ctx.Builder.CreateRetVoid()
			} else {
				zero := v.getZeroValue(retType)
				v.ctx.Builder.CreateRet(zero)
			}
		}
	}
	
	v.ctx.ExitFunction()
	return nil
}

// emitAsyncFunctionPrologue emits the coroutine initialization code
func (v *IRVisitor) emitAsyncFunctionPrologue(fn *ir.Function) {
	v.logger.Debug("Emitting async function prologue for %s", fn.Name())
	
	// Create coroutine ID
	coroId := v.ctx.Builder.CreateCoroId("")
	
	// Begin coroutine (allocate frame)
	coroHandle := v.ctx.Builder.CreateCoroBegin(coroId, "")
	
	// Store the handle somewhere accessible (in a stack slot for now)
	handleAlloca := v.ctx.Builder.CreateAlloca(coroHandle.Type(), "coro.handle")
	v.ctx.Builder.CreateStore(coroHandle, handleAlloca)
	
	// Store in scope so we can access it later
	v.ctx.currentScope.Define("__coro_handle__", handleAlloca)
	v.ctx.currentScope.Define("__coro_id__", coroId)
}

// emitAsyncFunctionEpilogue emits the coroutine cleanup code
func (v *IRVisitor) emitAsyncFunctionEpilogue(fn *ir.Function) {
	v.logger.Debug("Emitting async function epilogue for %s", fn.Name())
	
	// Retrieve the coroutine handle
	if handleSym, ok := v.ctx.currentScope.Lookup("__coro_handle__"); ok {
		if alloca, isAlloca := handleSym.Value.(*ir.AllocaInst); isAlloca {
			ptrType := alloca.Type().(*types.PointerType)
			handle := v.ctx.Builder.CreateLoad(ptrType.ElementType, alloca, "")
			
			// Mark end of coroutine
			v.ctx.Builder.CreateCoroEnd(handle)
			
			// Get memory to free
			if idSym, ok := v.ctx.currentScope.Lookup("__coro_id__"); ok {
				coroId := idSym.Value
				mem := v.ctx.Builder.CreateCoroFree(coroId, handle, "")
				_ = mem
			}
		}
	}
}

// ============================================================================
// VARIABLE & CONSTANT DECLARATIONS
// ============================================================================

func (v *IRVisitor) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	// Check for tuple pattern (let (a, b) = ...)
	if ctx.TuplePattern() != nil {
		return v.visitTupleVariableDecl(ctx)
	}
	
	// Safety check: Parser might return nil IDENTIFIER on syntax error
	if ctx.IDENTIFIER() == nil {
		return nil
	}

	name := ctx.IDENTIFIER().GetText()
	
	v.logger.Debug("Declaring variable: %s", name)
	
	var varType types.Type
	if ctx.Type_() != nil {
		varType = v.resolveType(ctx.Type_())
	}
	
	var initValue ir.Value
	if ctx.Expression() != nil {
		initValue = v.Visit(ctx.Expression()).(ir.Value)
		
		if varType == nil {
			varType = initValue.Type()
		}
		
		// Type check and cast if necessary
		if !initValue.Type().Equal(varType) {
			initValue = v.castValue(initValue, varType)
		}
	} else {
		if varType == nil {
			v.ctx.Logger.Error("Variable '%s' needs type annotation or initializer", name)
			return nil
		}
		initValue = v.getZeroValue(varType)
	}

	alloca := v.ctx.Builder.CreateAlloca(varType, name+".addr")

	// Workaround for backend limitation:
	// If initialization value is a ConstantArray, decompose it into individual stores
	if arrType, ok := varType.(*types.ArrayType); ok {
		if constArr, ok := initValue.(*ir.ConstantArray); ok {
			v.logger.Debug("Expanding array initialization for %s into %d stores", name, len(constArr.Elements))
			for i, elem := range constArr.Elements {
				// GEP to index i
				idx := v.ctx.Builder.ConstInt(types.I32, int64(i))
				zero := v.ctx.Builder.ConstInt(types.I32, 0)
				gep := v.ctx.Builder.CreateInBoundsGEP(arrType, alloca, []ir.Value{zero, idx}, "")
				v.ctx.Builder.CreateStore(elem, gep)
			}
			v.ctx.currentScope.Define(name, alloca)
			return nil
		}
	}

	v.ctx.Builder.CreateStore(initValue, alloca)
	v.ctx.currentScope.Define(name, alloca)
	
	return nil
}

func (v *IRVisitor) visitTupleVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	tuplePatternCtx := ctx.TuplePattern()
	names := make([]string, 0)
	for _, id := range tuplePatternCtx.AllIDENTIFIER() {
		names = append(names, id.GetText())
	}
	
	if ctx.Expression() == nil {
		v.ctx.Logger.Error("Tuple destructuring requires an initializer")
		return nil
	}
	
	// Evaluate the expression (should be a tuple/struct)
	tupleVal := v.Visit(ctx.Expression()).(ir.Value)
	tupleType, ok := tupleVal.Type().(*types.StructType)
	if !ok {
		v.ctx.Logger.Error("Tuple destructuring requires a tuple value, got %T", tupleVal.Type())
		return nil
	}
	
	if len(names) != len(tupleType.Fields) {
		v.ctx.Logger.Error("Tuple destructuring: expected %d values, got %d", 
			len(tupleType.Fields), len(names))
		return nil
	}
	
	// Spill the tuple value to a temporary stack slot to allow safe field extraction
	// This avoids potential backend issues with extractvalue on registers (or treating them as pointers)
	tempAlloca := v.ctx.Builder.CreateAlloca(tupleType, "tuple.destruct.temp")
	v.ctx.Builder.CreateStore(tupleVal, tempAlloca)
	
	// Extract each field and create variables
	for i, name := range names {
		// Get the actual field type from the struct
		fieldType := tupleType.Fields[i]
		
		// Create GEP to access the field from memory
		zero := v.ctx.Builder.ConstInt(types.I32, 0)
		idx := v.ctx.Builder.ConstInt(types.I32, int64(i))
		gep := v.ctx.Builder.CreateInBoundsGEP(tupleType, tempAlloca, []ir.Value{zero, idx}, "")
		
		// Load the field value
		fieldVal := v.ctx.Builder.CreateLoad(fieldType, gep, "")
		
		// Create alloca with the FIELD type
		alloca := v.ctx.Builder.CreateAlloca(fieldType, name+".addr")
		
		// Store the extracted field value
		v.ctx.Builder.CreateStore(fieldVal, alloca)
		
		// Define the variable in the current scope
		v.ctx.currentScope.Define(name, alloca)
		
		v.logger.Debug("Tuple destructure: %s = field %d (type: %v)", name, i, fieldType)
	}
	
	return nil
}

func (v *IRVisitor) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	if ctx.IDENTIFIER() == nil {
		return nil
	}
	name := ctx.IDENTIFIER().GetText()
	v.logger.Debug("Declaring constant: %s", name)

	if ctx.Expression() == nil {
		v.ctx.Logger.Error("Constant '%s' must have an initializer", name)
		return nil
	}

	initValue := v.Visit(ctx.Expression()).(ir.Value)
	
	// Check if this is a top-level constant (global)
	isGlobal := v.ctx.currentScope == v.ctx.globalScope
	
	if isGlobal {
		// Create mangled name for global
		globalName := name
		if v.ctx.currentNamespace != nil && v.ctx.currentNamespace.Name != "" {
			globalName = v.ctx.currentNamespace.Name + "_" + name
		}
		
		// Create IR Global for the constant
		if constant, ok := initValue.(ir.Constant); ok {
			g := v.ctx.Builder.CreateGlobalConstant(globalName, constant)
			// Register in scope as the global pointer
			v.ctx.currentScope.DefineConst(name, g)
			v.logger.Info("Declared global constant '%s' (IR: %s)", name, globalName)
		} else {
			v.ctx.Logger.Error("Top-level constant '%s' must have a constant initializer", name)
		}
	} else {
		// Local constant
		v.ctx.currentScope.DefineConst(name, initValue)
	}

	return nil
}