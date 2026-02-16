package codegen

import (
	"fmt"
	"github.com/arc-language/arc-lang/ast"
	
	"github.com/arc-language/arc-lang/builder"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

type Codegen struct {
	Builder *builder.Builder
	Module  *ir.Module
	TypeGen *TypeGenerator

	// Stack of scopes: Name -> Alloc Instruction (Stack Slot)
	scopes []map[string]ir.Value
}

func New(moduleName string) *Codegen {
	b := builder.New()
	mod := b.CreateModule(moduleName)
	
	return &Codegen{
		Builder: b,
		Module:  mod,
		TypeGen: NewTypeGenerator(),
		scopes:  make([]map[string]ir.Value, 0),
	}
}

func (cg *Codegen) Generate(file *ast.File) (*ir.Module, error) {
	// 1. Declare Functions (Forward Declaration)
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			cg.declareFunction(fn)
		}
	}

	// 2. Generate Function Bodies
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if err := cg.genFuncBody(fn); err != nil {
				return nil, err
			}
		}
	}

	return cg.Module, nil
}

// --- Scope Management ---

func (cg *Codegen) pushScope() {
	cg.scopes = append(cg.scopes, make(map[string]ir.Value))
}

func (cg *Codegen) popScope() {
	cg.scopes = cg.scopes[:len(cg.scopes)-1]
}

func (cg *Codegen) defineVar(name string, val ir.Value) {
	if len(cg.scopes) == 0 {
		panic("no scope to define variable")
	}
	cg.scopes[len(cg.scopes)-1][name] = val
}

func (cg *Codegen) lookupVar(name string) ir.Value {
	for i := len(cg.scopes) - 1; i >= 0; i-- {
		if v, ok := cg.scopes[i][name]; ok {
			return v
		}
	}
	return nil
}

// --- Global Constants ---

// createGlobalString creates a constant [N x i8] array global and returns a *i8 pointer
func (cg *Codegen) createGlobalString(str string) ir.Value {
	// 1. Convert string to i8 constants
	var chars []ir.Constant
	for i := 0; i < len(str); i++ {
		chars = append(chars, cg.Builder.ConstInt(types.I8, int64(str[i])))
	}
	chars = append(chars, cg.Builder.ConstInt(types.I8, 0)) // Null terminator

	// 2. Create Array Constant
	arrType := types.NewArray(types.I8, int64(len(chars)))
	arrConst := &ir.ConstantArray{Elements: chars}
	arrConst.SetType(arrType)

	// 3. Create Global Variable
	name := fmt.Sprintf(".str.%d", cg.Builder.Module().GlobalsCount())
	glob := cg.Builder.CreateGlobalConstant(name, arrConst)

	// 4. Bitcast/GEP to *i8
	zero := cg.Builder.ConstInt(types.I32, 0)
	ptr := cg.Builder.CreateInBoundsGEP(
		arrType,
		glob,
		[]ir.Value{zero, zero},
		name+".ptr",
	)
	return ptr
}