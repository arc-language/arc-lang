package compiler

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// Helper to get namespaced name
func (v *IRVisitor) getNamespacedName(name string) string {
	if v.ctx.currentNamespace != nil && v.ctx.currentNamespace.Name != "" {
		return v.ctx.currentNamespace.Name + "_" + name
	}
	return name
}

// registerStructType registers a struct type in pass 1
func (v *IRVisitor) registerStructType(ctx *parser.StructDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	
	// Register simple name in scope first to check duplicates
	if _, ok := v.ctx.GetType(name); ok {
		v.logger.Debug("Struct type '%s' already registered", name)
		return
	}
	
	v.logger.Debug("Registering struct type: %s", name)
	
	irName := v.getNamespacedName(name)
	
	// Create field map
	fieldMap := make(map[string]int)
	fieldTypes := make([]types.Type, 0)
	
	fieldIndex := 0
	for _, member := range ctx.AllStructMember() {
		if member.StructField() != nil {
			field := member.StructField()
			fieldName := field.IDENTIFIER().GetText()
			fieldType := v.resolveType(field.Type_())
			
			fieldTypes = append(fieldTypes, fieldType)
			fieldMap[fieldName] = fieldIndex
			fieldIndex++
		}
	}
	
	// Register mapping in context using IR Name
	v.ctx.StructFieldIndices[irName] = fieldMap

	// Create struct with IR Name
	structType := types.NewStruct(irName, fieldTypes, false)
	
	// Register simple name in symbol table pointing to this type
	v.ctx.RegisterType(name, structType)
	
	// Also ensure module knows about this type by IR name
	v.ctx.Module.Types[irName] = structType
}

// registerClassType registers a class type in pass 1
func (v *IRVisitor) registerClassType(ctx *parser.ClassDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	
	v.logger.Info("Registering class type: %s", name)
	
	if _, ok := v.ctx.GetType(name); ok {
		v.logger.Debug("Class '%s' already registered", name)
		return
	}
	
	irName := v.getNamespacedName(name)
	
	// Create field map
	fieldMap := make(map[string]int)
	fieldTypes := make([]types.Type, 0)
	
	fieldIndex := 0
	for _, member := range ctx.AllClassMember() {
		if member.ClassField() != nil {
			field := member.ClassField()
			fieldName := field.IDENTIFIER().GetText()
			fieldType := v.resolveType(field.Type_())
			
			fieldTypes = append(fieldTypes, fieldType)
			fieldMap[fieldName] = fieldIndex
			fieldIndex++
		}
	}
	
	// Register mapping using IR name
	v.ctx.ClassFieldIndices[irName] = fieldMap

	// Create struct type with IR name
	structType := types.NewStruct(irName, fieldTypes, false)
	
	// Register simple name in symbol table
	v.ctx.RegisterType(name, structType)
	
	// Mark IR name as class
	v.ctx.classTypes[irName] = true
	
	// Ensure module knows about it
	v.ctx.Module.Types[irName] = structType
	
	v.logger.Debug("Registered class '%s' (IR: %s) with %d fields", name, irName, len(fieldTypes))
}

func (v *IRVisitor) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	v.logger.Debug("Processing struct declaration: %s", name)
	
	for _, member := range ctx.AllStructMember() {
		if member.FunctionDecl() != nil {
			v.Visit(member.FunctionDecl())
		}
	}
	
	return nil
}

func (v *IRVisitor) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	v.logger.Info("Processing class declaration: %s", name)
	
	for i, member := range ctx.AllClassMember() {
		v.logger.Debug("Processing class member %d/%d", i+1, len(ctx.AllClassMember()))
		if member.FunctionDecl() != nil {
			v.Visit(member.FunctionDecl())
		} else if member.DeinitDecl() != nil {
			v.Visit(member.DeinitDecl())
		}
	}
	
	v.logger.Info("Completed class declaration: %s", name)
	return nil
}

func (v *IRVisitor) VisitClassField(ctx *parser.ClassFieldContext) interface{} { 
	return nil 
}

func (v *IRVisitor) VisitClassMember(ctx *parser.ClassMemberContext) interface{} {
	if ctx.ClassField() != nil { 
		return v.Visit(ctx.ClassField()) 
	}
	if ctx.FunctionDecl() != nil { 
		return v.Visit(ctx.FunctionDecl()) 
	}
	if ctx.DeinitDecl() != nil { 
		return v.Visit(ctx.DeinitDecl()) 
	}
	return nil
}

func (v *IRVisitor) VisitDeinitDecl(ctx *parser.DeinitDeclContext) interface{} {
	v.ctx.Logger.Warning("deinit is not yet implemented")
	return nil
}

func (v *IRVisitor) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	v.logger.Info("Processing enum declaration: %s", name)
	
	// Determine underlying type (default int32)
	var underlyingType types.Type = types.I32
	if ctx.PrimitiveType() != nil {
		typeName := ctx.PrimitiveType().GetText()
		if typ, ok := v.ctx.GetType(typeName); ok {
			underlyingType = typ
		}
	}
	
	// Ensure underlying type is an integer type
	intType, ok := underlyingType.(*types.IntType)
	if !ok {
		v.ctx.Logger.Error("Enum underlying type must be an integer type")
		intType = types.I32.(*types.IntType)
	}
	
	// Register enum as an alias to underlying type
	v.ctx.RegisterType(name, intType)
	
	// Process enum members as constants
	value := int64(0)
	for _, member := range ctx.AllEnumMember() {
		memberName := member.IDENTIFIER().GetText()
		
		if member.Expression() != nil {
			// Explicit value
			val := v.Visit(member.Expression())
			if irVal, ok := val.(ir.Value); ok {
				if constInt, ok := irVal.(*ir.ConstantInt); ok {
					value = constInt.Value
				}
			}
		}
		
		// Create constant for each enum member
		constVal := v.ctx.Builder.ConstInt(intType, value)
		
		// Register as namespaced constant: EnumName_MemberName
		fullName := v.getNamespacedName(name + "_" + memberName)
		global := v.ctx.Builder.CreateGlobalConstant(fullName, constVal)
		
		// Also register in current scope for direct access
		v.ctx.currentScope.DefineConst(memberName, global)
		
		v.logger.Debug("Registered enum member: %s.%s = %d", name, memberName, value)
		
		value++
	}
	
	v.logger.Info("Completed enum declaration: %s with underlying type %v", name, intType)
	return nil
}

func (v *IRVisitor) VisitEnumMember(ctx *parser.EnumMemberContext) interface{} {
	// Individual enum members are processed in VisitEnumDecl
	return nil
}