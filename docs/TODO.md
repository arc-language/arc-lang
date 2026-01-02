Here is a clean, prioritized `TODO.md` to guide the completion of the Arc compiler.

--- START OF FILE TODO.md ---

# Arc Compiler Roadmap

This roadmap orders tasks by dependency. We must build the foundation (types, memory) before building the house (features, async).

## Phase 1: The Foundation (IR & Memory)
*Focus: Getting basic logic to run correctly.*

- [ ] **Implement `alloca` & Variable Storage**
    *   **Why:** In IR, local variables must be allocated on the stack (`alloca`) so they have memory addresses. We cannot do assignments without this.
- [ ] **Implement Pointer Indirection (`&` and `*`)**
    *   **Why:** Essential for passing data by reference and handling struct fields.
- [ ] **Implement Control Flow (`if`, `while`, `for`, `break`)**
    *   **Why:** Requires managing Basic Blocks and Branch instructions. We need this logic stable before tackling complex iterator loops.
- [ ] **Implement Struct Field Access (`.`)**
    *   **Why:** Requires generating `GetElementPtr` (GEP) instructions to calculate memory offsets. This is the hardest part of basic IR generation.

## Phase 2: Functions & Interop
*Focus: Modularity and talking to the outside world.*

- [ ] **Implement Function Calls & Returns**
    *   **Why:** Need to handle argument passing and return types strictly.
- [ ] **Implement `extern` C Linking**
    *   **Why:** We need `libc` (malloc/free/printf) to implement Arrays/Vectors and print output.
- [ ] **Implement Intrinsics Mapping (`memcpy`, `sizeof`)**
    *   **Why:** Our `builtins.go` defined these, now `irgen` must actually emit the specific LLVM instructions for them.

## Phase 3: Data Structures (The "Library" Phase)
*Focus: Implementing the features removed from the parser.*

- [ ] **Implement `string` Literals & Operations**
    *   **Why:** Strings are complex (pointer + length). Need to generate global constants for literals.
- [ ] **Implement `array<T, N>` (Fixed Size)**
    *   **Why:** The simplest collection. Just a contiguous block of stack memory.
- [ ] **Implement `vector<T>` Layout**
    *   **Why:** Requires heap allocation (`malloc` via extern). Needs a struct definition (`ptr`, `len`, `cap`) generated manually in the compiler.

## Phase 4: Generics (Polymorphism)
*Focus: Making code reusable.*

- [ ] **Implement Monomorphization (Pass 2.5)**
    *   **Why:** We cannot generate IR for `vector<T>` directly. We must detect usage like `vector<int>` and generate a specific struct `vector_int` and specific functions `push_int`.

## Phase 5: Advanced Execution (The "Runtime" Phase)
*Focus: Threads, Async, and GPU.*

- [ ] **Implement Lambda Lifting (Closures)**
    *   **Why:** Anonymous functions capture variables. The compiler must wrap captured vars in a hidden struct and pass it to the function.
- [ ] **Implement `thread` / `process`**
    *   **Why:** These are just wrappers around system calls (pthreads/fork), but they depend on Closures working perfectly first.
- [ ] **Implement `async` / `await` State Machines**
    *   **Why:** The hardest feature. The compiler must chop a function into pieces at every `await` point. Impossible to debug if control flow (Phase 1) isn't rock solid.
- [ ] **Implement GPU PTX Backend**
    *   **Why:** Separate backend target. Requires its own specific intrinsics map.

--- END OF FILE TODO.md ---