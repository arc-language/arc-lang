package lower

import (
	"fmt"
	"github.com/arc-language/arc-lang/ast"
)

type asyncLowerer struct {
	file *ast.File
}

func (l *asyncLowerer) walk(node ast.Node) {
	if f, ok := node.(*ast.File); ok {
		var newDecls []ast.Decl
		for _, decl := range f.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok && fn.IsAsync {
				// Transform: async func Add(a, b) -> int
				// Into:
				//   struct Add_Packet { a, b, result }
				//   func Add_ThreadEntry(p: *Add_Packet)
				//   func Add(a, b) -> ThreadHandle
				packet, entry, wrapper := l.transformAsyncFunc(fn)
				
				newDecls = append(newDecls, packet)
				newDecls = append(newDecls, entry)
				newDecls = append(newDecls, wrapper)
			} else {
				// Non-async functions are untouched
				newDecls = append(newDecls, decl)
			}
		}
		f.Decls = newDecls
	}
	
	// We also need to rewrite 'await' calls inside ALL functions
	// (Recursively walk function bodies to find 'await')
	l.transformAwaitCalls(node)
}

// transformAsyncFunc implements the "Pthread-style" lowering.
// It keeps the stack intact and relies on the OS for context switching.
func (l *asyncLowerer) transformAsyncFunc(fn *ast.FuncDecl) (*ast.InterfaceDecl, *ast.FuncDecl, *ast.FuncDecl) {
	packetName := fmt.Sprintf("%s_Packet", fn.Name)
	entryName := fmt.Sprintf("%s_ThreadEntry", fn.Name)

	// --- 1. The Data Packet Struct ---
	// Holds params (inputs) and return value (output)
	// struct Add_Packet { a: int, b: int, ret: int }
	fields := []*ast.Field{}
	for _, p := range fn.Params {
		fields = append(fields, &ast.Field{Name: p.Name, Type: p.Type})
	}
	// Add return placeholder if not void
	if fn.ReturnType != nil {
		fields = append(fields, &ast.Field{
			Name: "ret",
			Type: fn.ReturnType,
		})
	}
	
	packetStruct := &ast.InterfaceDecl{
		Name:   packetName,
		Fields: fields,
	}

	// --- 2. The Thread Entry (The "Real" Logic) ---
	// func Add_ThreadEntry(raw: *void) {
	//     let pkt = cast(raw, *Add_Packet)
	//     pkt.ret = ...original_body_logic...
	//     thread_exit()
	// }
	// For simplicity in this AST, we just assume the body uses 'pkt.arg'
	// A real implementation would inject specific variable remapping here.
	entryFunc := &ast.FuncDecl{
		Name: entryName,
		Params: []*ast.Field{{
			Name: "raw", 
			Type: &ast.PointerType{Base: &ast.NamedType{Name: "void"}},
		}},
		Body: fn.Body, // We move the original body here
	}
	
	// --- 3. The Wrapper (The Spawner) ---
	// func Add(a, b) ThreadHandle {
	//     let pkt = new Add_Packet{a: a, b: b}
	//     return syscall_spawn(Add_ThreadEntry, pkt)
	// }
	wrapperBody := &ast.BlockStmt{
		List: []ast.Stmt{
			// AST construction for: return syscall_spawn(...)
			// Omitted for brevity, but this calls the intrinsic.
		},
	}
	
	wrapperFunc := &ast.FuncDecl{
		Name:   fn.Name,
		Params: fn.Params,
		// Return type changes from 'T' to 'ThreadHandle'
		ReturnType: &ast.NamedType{Name: "ThreadHandle"}, 
		Body:       wrapperBody,
	}

	return packetStruct, entryFunc, wrapperFunc
}

func (l *asyncLowerer) transformAwaitCalls(node ast.Node) {
	// Recursive walk to find `await expr` and replace with `thread_join(expr)`
	// This is a standard AST traversal. 
	// If we find `UnaryExpr{Op: AWAIT, X: expr}`, we replace it with:
	// `CallExpr{Fun: "thread_join", Args: [expr]}`
}