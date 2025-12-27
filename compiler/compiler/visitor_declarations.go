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
	name := ctx.IDENTIFIER().GetText()
	
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
	
	fn := v.ctx.Builder.DeclareFunction(name, retType, paramTypes, variadic)
	
	// Register in current namespace
	if v.ctx.currentNamespace != nil {
		v.ctx.currentNamespace.Functions[name] = fn
		v.logger.Debug("Declared extern function '%s' in namespace '%s'", name, v.ctx.currentNamespace.Name)
	} else {
		v.ctx.currentScope.Define(name, fn)
		v.logger.Debug("Declared extern function '%s' in global scope", name)
	}
	
	return nil
}

// ============================================================================
// FUNCTION DECLARATIONS
// ============================================================================

func (v *IRVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
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
				methodPrefix = structName + "_"
				name = methodPrefix + name
			}
		}
	}
	
	// Handle Namespacing
	var irName string = name
	
	// Special Case: main function should not be mangled
	isMain := name == "main" && (v.ctx.currentNamespace == nil || v.ctx.currentNamespace.Name == "main" || v.ctx.currentNamespace.Name == "")
	
	if !isMain && v.ctx.currentNamespace != nil && v.ctx.currentNamespace.Name != "" {
		irName = v.ctx.currentNamespace.Name + "_" + name
	}

	if isAsync {
		v.logger.Debug("Declaring async function: %s (IR: %s)", name, irName)
	} else {
		v.logger.Debug("Declaring function: %s (IR: %s)", name, irName)
	}
	
	var retType types.Type = types.Void
	if ctx.Type_() != nil {
		retType = v.resolveType(ctx.Type_())
	}
	
	// If async, wrap return type in a Future/Promise
	// For now, we'll keep the return type as-is and mark the function
	// The transformation pass will handle the actual conversion
	
	paramTypes := make([]types.Type, 0)
	paramNames := make([]string, 0)
	variadic := false
	
	if ctx.ParameterList() != nil {
		paramCtx := ctx.ParameterList()
		if paramCtx.ELLIPSIS() != nil {
			variadic = true
		}
		for _, param := range paramCtx.AllParameter() {
			paramName := param.IDENTIFIER().GetText()
			paramType := v.resolveType(param.Type_())
			paramNames = append(paramNames, paramName)
			paramTypes = append(paramTypes, paramType)
		}
	}

	fn := v.ctx.Builder.CreateFunction(irName, retType, paramTypes, variadic)
	
	// Mark as coroutine if async
	if isAsync {
		fn.Attributes = append(fn.Attributes, ir.AttrCoroutine)
		v.logger.Info("Marked function '%s' as async/coroutine", irName)
	}
	
	// Register function in the current namespace
	if v.ctx.currentNamespace != nil {
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
	// In a full implementation, this would be saved in the function's context
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
				
				// Free the memory (would call free or a custom allocator)
				// For now, just create the instruction
				_ = mem
			}
		}
	}
}

// ============================================================================
// VARIABLE & CONSTANT DECLARATIONS
// ============================================================================

func (v *IRVisitor) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
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
        } else {
            // ADD THIS: Cast init value to declared type if needed
            if !initValue.Type().Equal(varType) {
                initValue = v.castValue(initValue, varType)
            }
        }
    } else {
        if varType == nil {
            v.ctx.Logger.Error("Variable '%s' needs type annotation or initializer", name)
            return nil
        }
        initValue = v.getZeroValue(varType)
    }

    alloca := v.ctx.Builder.CreateAlloca(varType, name+".addr")
    v.ctx.Builder.CreateStore(initValue, alloca)
    v.ctx.currentScope.Define(name, alloca)
    
    return nil
}

func (v *IRVisitor) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	v.logger.Debug("Declaring constant: %s", name)
	
	if ctx.Expression() == nil {
		v.ctx.Logger.Error("Constant '%s' must have an initializer", name)
		return nil
	}
	
	initValue := v.Visit(ctx.Expression()).(ir.Value)
	v.ctx.currentScope.DefineConst(name, initValue)
	
	return nil
}