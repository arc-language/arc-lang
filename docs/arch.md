

## **Recommendation: Add an IR Transformation Pass**

Your compiler already has a clean pipeline:
```
Parser → IRVisitor (IR Builder) → Codegen → Object File
```

For async/await, you should add a **transformation pass** between IR building and codegen:

```
Parser → IRVisitor → Async Transform → Codegen → Object File
```




This is an **excellent** idea.

In Go, packaging is not just about organization; it’s about **enforcing dependencies**. If you keep everything in `compiler`, Pass 2 can accidentally read private state from Pass 1, or the Driver can accidentally touch low-level IR builder methods it shouldn't.

However, splitting packages in Go introduces one major risk: **Circular Dependencies**. (e.g., Package A needs Package B, but Package B needs Package A).

Here is the **Domain-Driven Design** approach I recommend for a compiler of this size.

### The Recommended Layout

Instead of naming them `pass1` or `pass2`, name them by **what they do**.

```text
/
├── cmd/
│   └── arc/              # Main entry point (calls compiler)
├── pkg/
│   ├── compiler/         # THE DRIVER (Orchestrator)
│   │   └── driver.go     # Calls Parser -> Semantics -> IRGen
│   │
│   ├── semantics/        # PASS 1 (The Validator)
│   │   ├── analyzer.go   # The Visitor
│   │   └── errors.go     # Semantic specific errors
│   │
│   ├── irgen/            # PASS 2 (The Translator)
│   │   ├── generator.go  # The Visitor
│   │   └── defer.go      # Defer stack logic
│   │
│   ├── symbol/           # SHARED DATA (The Glue)
│   │   ├── scope.go      # Scope implementation
│   │   ├── symbol.go     # Symbol struct definition
│   │   └── table.go      # Global table logic
│   │
│   ├── builder/          # (Existing) LLVM IR Builder
│   │   ├── ir/
│   │   └── types/
│   │
│   └── codegen/          # (Existing) ASM/Object generation
```

---

### Why this structure?

#### 1. `pkg/symbol` (The Foundation)
This is the most critical extraction. Both `semantics` and `irgen` need to know what a "Symbol" and a "Scope" are.
*   **Why split?** If `semantics` defines `Symbol`, and `irgen` imports `semantics`, `irgen` is now dependent on the entire analysis logic just to read a variable name.
*   **Content:** Only data structures. No heavy logic.

#### 2. `pkg/semantics` (The Brain)
*   **Imports:** `parser`, `symbol`, `builder/types`.
*   **Does not import:** `irgen`, `builder/ir`.
*   **Responsibility:** It takes the AST, validates it, and produces a `symbol.Table` (or `AnalysisResult`). It knows **nothing** about LLVM IR.

#### 3. `pkg/irgen` (The Builder)
*   **Imports:** `parser`, `symbol`, `builder/ir`, `builder/types`.
*   **Responsibility:** It trusts that `semantics` did its job. It takes the `AnalysisResult` and blindly builds IR.

#### 4. `pkg/compiler` (The Boss)
*   **Imports:** `semantics`, `irgen`, `codegen`.
*   **Responsibility:**
    1.  Read file.
    2.  `tree := parser.Parse(...)`
    3.  `analysis := semantics.Analyze(tree)`
    4.  `module := irgen.Generate(tree, analysis)`
    5.  `codegen.Write(module)`

---

### How the Code Changes

Here is how the interaction looks with this split.

#### 1. The Shared Symbol Package
`pkg/symbol/structs.go`
```go
package symbol

import "github.com/arc-language/arc-lang/pkg/builder/types"

// Shared between Semantic Analysis and IR Gen
type Symbol struct {
    Name string
    Type types.Type
    // We store IRValue as interface{} to avoid importing 'ir' package here
    // checking circular dependencies. Or keep it generic.
    Data interface{} 
}
```

#### 2. The Semantic Package
`pkg/semantics/analyzer.go`
```go
package semantics

import (
    "github.com/arc-language/arc-lang/pkg/symbol"
    "github.com/arc-language/arc-lang/pkg/parser"
)

type Result struct {
    GlobalScope *symbol.Scope
    NodeTypes   map[interface{}]types.Type
}

func Analyze(tree parser.ICompilationUnitContext) (*Result, error) {
    // ... logic ...
    return &Result{...}, nil
}
```

#### 3. The IRGen Package
`pkg/irgen/generator.go`
```go
package irgen

import (
    "github.com/arc-language/arc-lang/pkg/semantics" // To read results
    "github.com/arc-language/arc-lang/pkg/builder/ir"
)

func Generate(tree parser.ICompilationUnitContext, res *semantics.Result) *ir.Module {
    // ... logic ...
    // Use res.GlobalScope to look up types
}
```

#### 4. The Driver
`pkg/compiler/driver.go`
```go
package compiler

import (
    "github.com/arc-language/arc-lang/pkg/semantics"
    "github.com/arc-language/arc-lang/pkg/irgen"
)

func Compile(path string) {
    // 1. Parse
    tree := parse(path)

    // 2. Semantics
    analysis, err := semantics.Analyze(tree)
    if err != nil {
        panic(err)
    }

    // 3. IR Gen
    mod := irgen.Generate(tree, analysis)
    
    // 4. Codegen
    // ...
}
```

### Advantages of this approach:

1.  **Parallel Development:** One person can work on `semantics` (type checking rules) while another works on `irgen` (LLVM instruction selection) with minimal merge conflicts.
2.  **Testability:** You can write unit tests for `semantics` that just check if "string + int" returns an error, without needing to set up the entire LLVM backend.
3.  **Strict Boundaries:** `semantics` physically *cannot* generate IR because it doesn't import the `ir` package. This prevents "hacking in a quick fix" where you generate code during analysis.

### Conclusion

**Yes, do it.**

Move the logic into:
1.  `pkg/symbol` (The data structures)
2.  `pkg/semantics` (Pass 1)
3.  `pkg/irgen` (Pass 2)
4.  `pkg/compiler` (The glue)





This is the "Pipeline of Errors". In a production-grade compiler, errors are distinct based on **when** they are caught.

Here is exactly where each error lives in your new architecture:

1.  **Syntax Errors (Grammar)** → Caught in **`pkg/compiler`** (The Parsing Phase).
    *   *Examples:* Missing `;`, unmatched `}`, invalid keyword `func` instead of `fn`.
2.  **Semantic Errors (Logic/Types)** → Caught in **`pkg/semantics`** (Pass 1).
    *   *Examples:* Undefined variable `x`, type mismatch `int + string`, assigning to a `const`.
3.  **Generation Errors (Backend)** → Caught in **`pkg/irgen`** (Pass 2).
    *   *Examples:* Feature not implemented yet, backend limitation. (Ideally, this should be empty if Pass 1 is perfect).

---

### The New Package Structure

First, we need a shared way to report errors so `semantics` doesn't depend on `compiler`. We'll add a tiny `diagnostic` package.

```text
/pkg
├── diagnostic/       # [NEW] Shared error types (Color printing, file positions)
├── symbol/           # Shared data (Scope, Symbol)
├── semantics/        # [PASS 1] Type Checking & Logic Validation
├── irgen/            # [PASS 2] Code Generation
└── compiler/         # [DRIVER] Parsing & Orchestration
```

---

### 1. The Shared Error System (`pkg/diagnostic`)

This prevents circular dependencies. Both `compiler` and `semantics` import this to report errors.

```go
// pkg/diagnostic/errors.go
package diagnostic

import "fmt"

type Error struct {
	File    string
	Line    int
	Column  int
	Message string
}

func (e *Error) String() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.File, e.Line, e.Column, e.Message)
}

// Bag collects errors from different passes
type Bag struct {
	Errors []*Error
}

func NewBag() *Bag {
	return &Bag{Errors: make([]*Error, 0)}
}

func (b *Bag) Report(file string, line, column int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	b.Errors = append(b.Errors, &Error{
		File:    file,
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

func (b *Bag) HasErrors() bool {
	return len(b.Errors) > 0
}
```

---

### 2. Syntax Errors (`pkg/compiler`)

Syntax errors happen **before** any of your code runs, inside the ANTLR parser. You catch them in the Driver.

```go
// pkg/compiler/parser.go
package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/diagnostic"
)

// syntaxListener pipes ANTLR errors into our Diagnostic Bag
type syntaxListener struct {
	*antlr.DefaultErrorListener
	filename string
	bag      *diagnostic.Bag
}

func (l *syntaxListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// EXPLICIT SYNTAX ERROR HANDLING
	l.bag.Report(l.filename, line, column, "Syntax Error: %s", msg)
}

func Parse(filename string) (parser.ICompilationUnitContext, *diagnostic.Bag) {
	bag := diagnostic.NewBag()
	
	// ... setup ANTLR input ...
	
	p := parser.NewArcParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&syntaxListener{
		filename: filename,
		bag:      bag,
	})
	
	tree := p.CompilationUnit()
	return tree, bag
}
```

---

### 3. Type Checking Errors (`pkg/semantics`)

This is where the "real" compiler work happens. This is Pass 1.

```go
// pkg/semantics/analyzer.go
package semantics

import (
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/diagnostic"
	"github.com/arc-language/arc-lang/pkg/symbol"
)

type Analyzer struct {
	bag          *diagnostic.Bag
	currentScope *symbol.Scope
	// ...
}

func Analyze(tree parser.ICompilationUnitContext, bag *diagnostic.Bag) (*Result, error) {
	a := &Analyzer{
		bag: bag,
		// ...
	}
	a.Visit(tree)
	
	if bag.HasErrors() {
		return nil, fmt.Errorf("semantic analysis failed")
	}
	return &Result{...}, nil
}

// EXAMPLE: Type Checking Logic
func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// 1. CHECK DECLARATION
	if _, exists := a.currentScope.ResolveLocal(name); exists {
		// LOGIC ERROR: Variable already exists
		a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"Variable '%s' is already defined in this scope", name)
		return nil
	}

	// 2. CHECK TYPES
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		declType := a.resolveType(ctx.Type_())

		if !exprType.Equal(declType) {
			// TYPE ERROR: Mismatch
			a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				"Type Mismatch: Cannot assign type '%s' to variable of type '%s'", 
				exprType.String(), declType.String())
		}
	}
	
	return nil
}

func (a *Analyzer) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	// 3. CHECK UNDEFINED
	name := ctx.IDENTIFIER().GetText()
	if _, ok := a.currentScope.Resolve(name); !ok {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
			"Undefined variable '%s'", name)
		return types.Void // Return poison type to suppress cascade errors
	}
	// ...
}
```

---

### 4. The Orchestrator (`pkg/compiler`)

The driver runs the pipeline and stops if any stage reports errors.

```go
// pkg/compiler/compiler.go
package compiler

import (
	"fmt"
	"github.com/arc-language/arc-lang/pkg/diagnostic"
	"github.com/arc-language/arc-lang/pkg/irgen"
	"github.com/arc-language/arc-lang/pkg/semantics"
)

func Compile(path string) {
	// STAGE 1: PARSING
	// Captures Syntax Errors (missing semicolons, bad keywords)
	tree, syntaxErrors := Parse(path)
	if syntaxErrors.HasErrors() {
		printErrors(syntaxErrors)
		return // Stop here
	}

	// STAGE 2: SEMANTICS
	// Captures Type Errors (int = string, undefined vars)
	// We pass the SAME bag or a new one to collect these errors
	semanticErrors := diagnostic.NewBag()
	analysis, err := semantics.Analyze(tree, semanticErrors)
	
	if semanticErrors.HasErrors() {
		printErrors(semanticErrors)
		return // Stop here, do not attempt to generate IR
	}

	// STAGE 3: IR GENERATION
	// No user errors should happen here if Stage 2 worked.
	// We trust the 'analysis' result completely.
	mod := irgen.Generate(tree, analysis)
	
	// STAGE 4: CODEGEN
	// ...
}

func printErrors(bag *diagnostic.Bag) {
	for _, err := range bag.Errors {
		// You can add color support here (Red for Error, Yellow for Warning)
		fmt.Printf("%s\n", err.String())
	}
}
```

### Summary of Benefits

1.  **Clean Separation**: Pass 2 (`irgen`) **never** has to check `if variable_exists`. It just calls `scope.Resolve()` and knows it will succeed because Pass 1 (`semantics`) guaranteed it.
2.  **Fast Feedback**: If there is a syntax error, the compiler stops instantly. It doesn't try to type-check broken syntax.
3.  **Testability**: You can write tests specifically for `pkg/semantics` that expect specific type errors without needing to parse a file or generate IR.