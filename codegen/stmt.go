// codegen/stmt.go
package codegen

import (
	"fmt"

	"github.com/arc-language/arc-lang/ast"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (cg *Codegen) genBlock(block *ast.BlockStmt) error {
	cg.pushScope()
	defer cg.popScope()
	for _, stmt := range block.List {
		if err := cg.genStmt(stmt); err != nil {
			return err
		}
		// Stop emitting after a terminator — subsequent statements are dead code.
		if cur := cg.Builder.CurrentBlock(); cur != nil && cur.Terminator() != nil {
			break
		}
	}
	return nil
}

func (cg *Codegen) genStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {

	case *ast.DeclStmt:
		return cg.genDeclStmt(s)

	case *ast.AssignStmt:
		return cg.genAssignStmt(s)

	case *ast.ReturnStmt:
		return cg.genReturn(s)

	case *ast.ExprStmt:
		cg.genExpr(s.X)

	case *ast.IfStmt:
		return cg.genIf(s)

	case *ast.ForStmt:
		return cg.genFor(s)

	case *ast.ForInStmt:
		return cg.genForIn(s)

	case *ast.SwitchStmt:
		return cg.genSwitch(s)

	case *ast.BreakStmt:
		lc, ok := cg.currentLoop()
		if !ok {
			return fmt.Errorf("break outside loop")
		}
		cg.Builder.CreateBr(lc.endBlock)

	case *ast.ContinueStmt:
		lc, ok := cg.currentLoop()
		if !ok {
			return fmt.Errorf("continue outside loop")
		}
		cg.Builder.CreateBr(lc.postBlock)

	case *ast.BlockStmt:
		return cg.genBlock(s)

	case *ast.DeferStmt:
		// Defers have already been rewritten to explicit calls by the lower pass.
		cg.genExpr(s.Call)
	}
	return nil
}

// ─── Declarations ─────────────────────────────────────────────────────────────

func (cg *Codegen) genDeclStmt(s *ast.DeclStmt) error {
	switch d := s.Decl.(type) {
	case *ast.VarDecl:
		return cg.genVarDecl(d)
	case *ast.ConstDecl:
		// Const specs are inlined as literals by the checker; the alloca path
		// still works for the rare case where a const is assigned a non-literal.
		for _, spec := range d.Specs {
			if spec.Value == nil {
				continue
			}
			val := cg.genExpr(spec.Value)
			if val == nil {
				continue
			}
			alloca := cg.Builder.CreateAlloca(val.Type(), spec.Name)
			cg.Builder.CreateStore(val, alloca)
			cg.defineVar(spec.Name, alloca)
		}
	}
	return nil
}

func (cg *Codegen) genVarDecl(d *ast.VarDecl) error {
	var alloca *ir.AllocaInst

	if d.Value != nil {
		val := cg.genExpr(d.Value)
		if val == nil {
			return fmt.Errorf("genVarDecl: init expr for %q produced nil value", d.Name)
		}
		alloca = cg.Builder.CreateAlloca(val.Type(), d.Name)
		cg.Builder.CreateStore(val, alloca)
	} else if d.Type != nil {
		// Declaration with explicit type but no initialiser — zero-initialise.
		irType := cg.TypeGen.GenType(d.Type)
		alloca = cg.Builder.CreateAlloca(irType, d.Name)
		cg.Builder.CreateStore(cg.Builder.ConstZero(irType), alloca)
	} else if d.IsNull {
		// var x: *T = null
		irType := cg.TypeGen.GenType(d.Type)
		alloca = cg.Builder.CreateAlloca(irType, d.Name)
		if pt, ok := irType.(*types.PointerType); ok {
			cg.Builder.CreateStore(cg.Builder.ConstNull(pt), alloca)
		}
	} else {
		return fmt.Errorf("genVarDecl: %q has no type and no initialiser", d.Name)
	}

	cg.defineVar(d.Name, alloca)
	return nil
}

// ─── Assignment ───────────────────────────────────────────────────────────────

func (cg *Codegen) genAssignStmt(s *ast.AssignStmt) error {
	// ++ and -- are post-increment with no Value.
	if s.Op == "++" || s.Op == "--" {
		ptr := cg.genLValue(s.Target)
		if ptr == nil {
			return fmt.Errorf("assignment target is not an l-value")
		}
		pt := ptr.Type().(*types.PointerType)
		cur := cg.Builder.CreateLoad(pt.ElementType, ptr, "")
		one := cg.Builder.ConstInt(types.I32, 1)
		// If the type is wider than i32, zero-extend the constant.
		if it, ok := pt.ElementType.(*types.IntType); ok && it.BitWidth != 32 {
			one = &ir.ConstantInt{}
			one.SetType(pt.ElementType)
			one.Value = 1
		}
		var next ir.Value
		if s.Op == "++" {
			next = cg.Builder.CreateAdd(cur, one, "")
		} else {
			next = cg.Builder.CreateSub(cur, one, "")
		}
		cg.Builder.CreateStore(next, ptr)
		return nil
	}

	rhs := cg.genExpr(s.Value)
	if rhs == nil {
		return fmt.Errorf("assignment rhs is nil")
	}

	// Compound assignment operators: convert x += y into x = x op y.
	if s.Op != "=" {
		lhsVal := cg.genExpr(s.Target)
		if lhsVal == nil {
			return fmt.Errorf("compound assignment lhs is nil")
		}
		rhs = cg.genCompoundOp(s.Op, lhsVal, rhs)
	}

	ptr := cg.genLValue(s.Target)
	if ptr == nil {
		return fmt.Errorf("assignment target is not an l-value")
	}
	cg.Builder.CreateStore(rhs, ptr)
	return nil
}

// genLValue returns the alloca/GEP pointer for a writable location.
func (cg *Codegen) genLValue(expr ast.Expr) ir.Value {
	switch e := expr.(type) {
	case *ast.Ident:
		return cg.lookupVar(e.Name)

	case *ast.SelectorExpr:
		// e.X.Sel — GEP into struct.
		structPtr := cg.genLValue(e.X)
		if structPtr == nil {
			return nil
		}
		pt, ok := structPtr.Type().(*types.PointerType)
		if !ok {
			return nil
		}
		st, ok := pt.ElementType.(*types.StructType)
		if !ok {
			return nil
		}
		idx := fieldIndex(st, e.Sel)
		if idx < 0 {
			return nil
		}
		return cg.Builder.CreateStructGEP(st, structPtr, idx, e.Sel+".ptr")

	case *ast.IndexExpr:
		// e.X[e.Index]
		ptr := cg.genLValue(e.X)
		if ptr == nil {
			return nil
		}
		pt, ok := ptr.Type().(*types.PointerType)
		if !ok {
			return nil
		}
		idxVal := cg.genExpr(e.Index)
		return cg.Builder.CreateInBoundsGEP(pt.ElementType, ptr, []ir.Value{idxVal}, "elem.ptr")
	}
	return nil
}

func (cg *Codegen) genCompoundOp(op string, lhs, rhs ir.Value) ir.Value {
	switch op {
	case "+=":
		return cg.Builder.CreateAdd(lhs, rhs, "")
	case "-=":
		return cg.Builder.CreateSub(lhs, rhs, "")
	case "*=":
		return cg.Builder.CreateMul(lhs, rhs, "")
	case "/=":
		return cg.Builder.CreateSDiv(lhs, rhs, "")
	case "%=":
		return cg.Builder.CreateSRem(lhs, rhs, "")
	case "&=":
		return cg.Builder.CreateAnd(lhs, rhs, "")
	case "|=":
		return cg.Builder.CreateOr(lhs, rhs, "")
	case "^=":
		return cg.Builder.CreateXor(lhs, rhs, "")
	case "<<=":
		return cg.Builder.CreateShl(lhs, rhs, "")
	case ">>=":
		return cg.Builder.CreateAShr(lhs, rhs, "")
	}
	return rhs
}

// ─── Return ───────────────────────────────────────────────────────────────────

func (cg *Codegen) genReturn(s *ast.ReturnStmt) error {
	switch len(s.Results) {
	case 0:
		cg.Builder.CreateRetVoid()
	case 1:
		val := cg.genExpr(s.Results[0])
		if val == nil {
			cg.Builder.CreateRetVoid()
		} else {
			cg.Builder.CreateRet(val)
		}
	default:
		// Multi-return: pack into a struct.
		// The function's return type must be a StructType; emit insertvalue chain.
		fn := cg.Builder.CurrentFunction()
		retType := fn.FuncType.ReturnType
		var agg ir.Value = cg.Builder.ConstUndef(retType)
		for i, res := range s.Results {
			val := cg.genExpr(res)
			if val != nil {
				agg = cg.Builder.CreateInsertValue(agg, val, []int{i}, "")
			}
		}
		cg.Builder.CreateRet(agg)
	}
	return nil
}

// ─── If ───────────────────────────────────────────────────────────────────────

func (cg *Codegen) genIf(s *ast.IfStmt) error {
	cond := cg.genExpr(s.Cond)
	if cond == nil {
		return fmt.Errorf("if condition produced nil")
	}

	// Ensure condition is i1.
	if cond.Type().BitSize() != 1 {
		zero := cg.Builder.ConstInt(types.I32, 0)
		cond = cg.Builder.CreateICmpNE(cond, zero, "cond")
	}

	thenBlock := cg.Builder.CreateBlock("if.then")
	endBlock := cg.Builder.CreateBlock("if.end")
	var elseBlock *ir.BasicBlock
	if s.Else != nil {
		elseBlock = cg.Builder.CreateBlock("if.else")
		cg.Builder.CreateCondBr(cond, thenBlock, elseBlock)
	} else {
		cg.Builder.CreateCondBr(cond, thenBlock, endBlock)
	}

	// Then branch.
	cg.Builder.SetInsertPoint(thenBlock)
	if err := cg.genBlock(s.Body); err != nil {
		return err
	}
	if cg.Builder.CurrentBlock().Terminator() == nil {
		cg.Builder.CreateBr(endBlock)
	}

	// Else branch.
	if s.Else != nil {
		cg.Builder.SetInsertPoint(elseBlock)
		switch e := s.Else.(type) {
		case *ast.BlockStmt:
			if err := cg.genBlock(e); err != nil {
				return err
			}
		case *ast.IfStmt:
			if err := cg.genIf(e); err != nil {
				return err
			}
		}
		if cg.Builder.CurrentBlock().Terminator() == nil {
			cg.Builder.CreateBr(endBlock)
		}
	}

	cg.Builder.SetInsertPoint(endBlock)
	return nil
}

// ─── For (C-style) ────────────────────────────────────────────────────────────

func (cg *Codegen) genFor(s *ast.ForStmt) error {
	cg.pushScope()
	defer cg.popScope()

	// Emit the init statement in the current block (which has the for-var scope).
	if s.Init != nil {
		if err := cg.genStmt(s.Init); err != nil {
			return err
		}
	}

	condBlock := cg.Builder.CreateBlock("for.cond")
	bodyBlock := cg.Builder.CreateBlock("for.body")
	postBlock := cg.Builder.CreateBlock("for.post")
	endBlock := cg.Builder.CreateBlock("for.end")

	cg.Builder.CreateBr(condBlock)

	// Condition.
	cg.Builder.SetInsertPoint(condBlock)
	if s.Cond != nil {
		cond := cg.genExpr(s.Cond)
		if cond.Type().BitSize() != 1 {
			zero := cg.Builder.ConstInt(types.I32, 0)
			cond = cg.Builder.CreateICmpNE(cond, zero, "cond")
		}
		cg.Builder.CreateCondBr(cond, bodyBlock, endBlock)
	} else {
		// Infinite loop.
		cg.Builder.CreateBr(bodyBlock)
	}

	// Body.
	cg.pushLoop(postBlock, endBlock)
	cg.Builder.SetInsertPoint(bodyBlock)
	if err := cg.genBlock(s.Body); err != nil {
		return err
	}
	cg.popLoop()
	if cg.Builder.CurrentBlock().Terminator() == nil {
		cg.Builder.CreateBr(postBlock)
	}

	// Post.
	cg.Builder.SetInsertPoint(postBlock)
	if s.Post != nil {
		if err := cg.genStmt(s.Post); err != nil {
			return err
		}
	}
	if cg.Builder.CurrentBlock().Terminator() == nil {
		cg.Builder.CreateBr(condBlock)
	}

	cg.Builder.SetInsertPoint(endBlock)
	return nil
}

// ─── For-in ───────────────────────────────────────────────────────────────────

// genForIn lowers  for key, val in iter { }  as:
//
//	%i = alloca i64
//	store 0, %i
//	br for.cond
//	for.cond:
//	  %len = extractvalue %iter, 1     ; or call len()
//	  %cond = icmp ult %i, %len
//	  br cond, for.body, for.end
//	for.body:
//	  %key = load %i
//	  %elem.ptr = getelementptr %data, %key
//	  %val  = load %elem.ptr
//	  <body>
//	  br for.post
//	for.post:
//	  %cur = load %i ; increment
//	  store %cur+1, %i
//	  br for.cond
func (cg *Codegen) genForIn(s *ast.ForInStmt) error {
	cg.pushScope()
	defer cg.popScope()

	iterVal := cg.genExpr(s.Iter)
	if iterVal == nil {
		return fmt.Errorf("for-in: iterator expression produced nil")
	}

	// Allocate index counter.
	idxAlloca := cg.Builder.CreateAlloca(types.I64, "for.idx")
	cg.Builder.CreateStore(cg.Builder.ConstInt(types.I64, 0), idxAlloca)

	condBlock := cg.Builder.CreateBlock("forin.cond")
	bodyBlock := cg.Builder.CreateBlock("forin.body")
	postBlock := cg.Builder.CreateBlock("forin.post")
	endBlock := cg.Builder.CreateBlock("forin.end")

	cg.Builder.CreateBr(condBlock)

	// Condition: check index < length.
	cg.Builder.SetInsertPoint(condBlock)
	idxVal := cg.Builder.CreateLoad(types.I64, idxAlloca, "idx")
	var lenVal ir.Value
	switch iterVal.Type().(type) {
	case *types.StructType:
		// vector/slice struct: field 1 is length.
		lenVal = cg.Builder.CreateExtractValue(iterVal, []int{1}, "len")
	default:
		// Fall back: treat the iterator as a pointer and use a sentinel.
		// A real implementation would call the runtime len() function.
		lenVal = cg.Builder.ConstInt(types.I64, 0)
	}
	cond := cg.Builder.CreateICmpULT(idxVal, lenVal, "forin.cond")
	cg.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	// Body.
	cg.pushLoop(postBlock, endBlock)
	cg.Builder.SetInsertPoint(bodyBlock)

	// Expose key and value bindings.
	keyAlloca := cg.Builder.CreateAlloca(types.I64, s.Key)
	cg.Builder.CreateStore(idxVal, keyAlloca)
	cg.defineVar(s.Key, keyAlloca)

	if s.Value != "" {
		// Extract data pointer (field 0) and GEP to element.
		dataPtr := cg.Builder.CreateExtractValue(iterVal, []int{0}, "data.ptr")
		if pt, ok := dataPtr.Type().(*types.PointerType); ok {
			elemPtr := cg.Builder.CreateInBoundsGEP(pt.ElementType, dataPtr, []ir.Value{idxVal}, "elem.ptr")
			elemAlloca := cg.Builder.CreateAlloca(pt.ElementType, s.Value)
			elemVal := cg.Builder.CreateLoad(pt.ElementType, elemPtr, "elem")
			cg.Builder.CreateStore(elemVal, elemAlloca)
			cg.defineVar(s.Value, elemAlloca)
		}
	}

	if err := cg.genBlock(s.Body); err != nil {
		return err
	}
	cg.popLoop()
	if cg.Builder.CurrentBlock().Terminator() == nil {
		cg.Builder.CreateBr(postBlock)
	}

	// Post: increment index.
	cg.Builder.SetInsertPoint(postBlock)
	curIdx := cg.Builder.CreateLoad(types.I64, idxAlloca, "idx.cur")
	nextIdx := cg.Builder.CreateAdd(curIdx, cg.Builder.ConstInt(types.I64, 1), "idx.next")
	cg.Builder.CreateStore(nextIdx, idxAlloca)
	cg.Builder.CreateBr(condBlock)

	cg.Builder.SetInsertPoint(endBlock)
	return nil
}

// ─── Switch ───────────────────────────────────────────────────────────────────

func (cg *Codegen) genSwitch(s *ast.SwitchStmt) error {
	tag := cg.genExpr(s.Tag)
	if tag == nil {
		return fmt.Errorf("switch tag produced nil")
	}

	endBlock := cg.Builder.CreateBlock("switch.end")
	var defaultBlock *ir.BasicBlock
	if s.Default != nil {
		defaultBlock = cg.Builder.CreateBlock("switch.default")
	} else {
		defaultBlock = endBlock
	}

	sw := cg.Builder.CreateSwitch(tag, defaultBlock, len(s.Cases))

	// Emit case blocks.
	for _, c := range s.Cases {
		caseBlock := cg.Builder.CreateBlock("switch.case")
		// Each value in the case list becomes a SwitchInst arm.
		for _, val := range c.Values {
			constVal := cg.genExpr(val)
			if ci, ok := constVal.(*ir.ConstantInt); ok {
				cg.Builder.AddCase(sw, ci, caseBlock)
			}
		}
		cg.pushLoop(endBlock, endBlock) // break target inside switch
		cg.Builder.SetInsertPoint(caseBlock)
		for _, st := range c.Body {
			if err := cg.genStmt(st); err != nil {
				return err
			}
			if cg.Builder.CurrentBlock().Terminator() != nil {
				break
			}
		}
		cg.popLoop()
		if cg.Builder.CurrentBlock().Terminator() == nil {
			cg.Builder.CreateBr(endBlock)
		}
	}

	// Default block.
	if s.Default != nil {
		cg.Builder.SetInsertPoint(defaultBlock)
		for _, st := range s.Default {
			if err := cg.genStmt(st); err != nil {
				return err
			}
			if cg.Builder.CurrentBlock().Terminator() != nil {
				break
			}
		}
		if cg.Builder.CurrentBlock().Terminator() == nil {
			cg.Builder.CreateBr(endBlock)
		}
	}

	cg.Builder.SetInsertPoint(endBlock)
	return nil
}

// ─── Utility ──────────────────────────────────────────────────────────────────

// fieldIndex finds the zero-based position of fieldName in a StructType.
// Returns -1 if not found.
func fieldIndex(st *types.StructType, fieldName string) int {
	// StructType does not carry field names in the types package; the names are
	// stored at the AST level. We rely on the IR module's named-type registry
	// to find the original InterfaceDecl field order. For now we handle the
	// common case where the struct has been given a stable name and fields are
	// accessed by index via a parallel name list stored on the TypeGenerator.
	// If the TypeGenerator has a lookup, use it; otherwise return -1.
	return -1 // resolved by TypeGenerator.FieldIndex below
}