# Arc GPU Architecture & PTX Generation

This document specifies how Arc maps high-level syntax to NVIDIA's PTX (Parallel Thread Execution) assembly. It details the "One Parser, Two Backends" architecture, the extended Type System, and the Intrinsic Interface used to expose hardware-specific features without altering the core grammar.

---

## 1. Compiler Architecture

Arc uses a **Unified Frontend** with a **Split Backend**. We do not use a separate parser for GPU code; the differentiation happens during the Intermediate Representation (IR) and Backend stages.

### The Compilation Pipeline

#### Pass 1: Frontend (Parser & Type Checker)
*   **Input:** Source code (`.arc`).
*   **Logic:**
    *   Parses standard definitions (`struct`, `func`, `if`, `let`).
    *   Validates types (including GPU-specific types like `float16`).
    *   Resolves `gpu.*` calls using the `extern gpu` definitions in the standard library.
    *   **Validation Rule:** Functions tagged with `<gpu>` can *only* call other GPU functions or `extern gpu` intrinsics.

#### Pass 2: IR Generation
*   **Logic:** Generates High-Level IR (SSA) for the entire program.
*   **Tagging:** Functions defined as `async func<gpu>` are encapsulated into specific IR Modules tagged with `Target: GPU`.

#### Pass 3: Backend Dispatch (The Fork)
The compiler iterates through generated IR modules and dispatches them to the correct backend package.

| Target | Package | Output | Role |
| :--- | :--- | :--- | :--- |
| **CPU** | `pkg/codegen` | Machine Code (ELF/PE) | Application logic, memory management, kernel launching. |
| **GPU** | `pkg/ptxgen` | PTX Assembly (Text) | High-performance parallel kernels, lowered via intrinsics. |

---

## 2. Control Flow Mapping

The `pkg/ptxgen` backend is stateless. It maps standard Arc control flow structures directly to PTX logic, automatically choosing between branching and predication based on block size/complexity.

**No new syntax is required.**

| Arc Source | Backend Logic | Resulting PTX Assembly |
| :--- | :--- | :--- |
| `if (cond)` | **Short Block?** → Predication | `setp.lt.u32 %p1...`<br>`@%p1 add.f32...` |
| `if (cond)` | **Long Block?** → Branching | `setp.lt.u32 %p1...`<br>`@!%p1 bra LABEL_ELSE;` |
| `for (init; cond; post)` | Loop Structure | `LABEL_LOOP:`<br>`...`<br>`bra LABEL_LOOP;` |
| `break` | Jump to Exit | `bra LABEL_EXIT;` |
| `continue` | Jump to Header | `bra LABEL_LOOP;` |

---

## 3. Type System Extension

To support modern GPU hardware (Volta, Ampere, Hopper), the Type Checker recognizes specific types that map to hardware registers or memory spaces.

### A. Standard Types (Common)
*These map 1:1 to PTX scalar types.*

| Arc Type | PTX Type |
| :--- | :--- |
| `bool` | `.pred` |
| `int8` / `uint8` | `.s8` / `.u8` |
| `int16` / `uint16` | `.s16` / `.u16` |
| `int32` / `uint32` | `.s32` / `.u32` |
| `int64` / `uint64` | `.s64` / `.u64` |
| `float32` | `.f32` |
| `float64` | `.f64` |

### B. GPU Register Types
*These lower to specific hardware capabilities (Tensor Cores, SIMD).*

| Arc Type | PTX Mapping | Description |
| :--- | :--- | :--- |
| `float16` | `.f16` | IEEE 754 Half Precision |
| `bfloat16` | `.bf16` | Brain Floating Point (AI/LLMs) |
| `float8_e4m3` | `.b8` (storage) | FP8 (Hopper Inference) |
| `float8_e5m2` | `.b8` (storage) | FP8 (Hopper Gradients) |
| `vector2<T>` | `.v2` | 2-element tuple (e.g., `ld.global.v2`) |
| `vector4<T>` | `.v4` | 4-element tuple (e.g., `ld.global.v4`) |

### C. Memory Space & Opaque Handles
*`shared` is a storage keyword. PascalCase types are opaque handles in the `gpu` namespace.*

| Arc Type | PTX Mapping | Usage |
| :--- | :--- | :--- |
| `shared<T, N>` | `.shared` | L1 Shared Memory allocation |
| `gpu.Barrier` | `.b64` | Hardware `mbarrier` object |
| `gpu.TensorMap`| `.u64` | TMA Descriptor Pointer (H100) |
| `gpu.Fragment<T,M,N,K>` | `.reg` array | Tensor Core Matrix Fragment |

---

## 4. The `extern gpu` Standard Library

The following definitions exist in `lib/gpu/intrinsics.arc`.
The `extern gpu` keyword tells the compiler: **"Do not link this. Use the internal `pkg/ptxgen` intrinsic mapper to generate assembly."**

*Note: Argument names are omitted; matching is done by type signature.*

```arc
namespace gpu

extern gpu {
    // ==========================================
    // 1. Indexing & Dimensions
    // ==========================================
    func thread_id() int32      // %tid.x
    func thread_id_y() int32    // %tid.y
    func thread_id_z() int32    // %tid.z
    func block_id() int32       // %ctaid.x
    func block_id_y() int32     // %ctaid.y
    func block_id_z() int32     // %ctaid.z
    func lane_id() int32        // %laneid
    func warp_id() int32        // %warpid
    func grid_dim() int32       // %nctaid.x
    func block_dim() int32      // %ntid.x

    // ==========================================
    // 2. Math Intrinsics
    // Compiler selects opcode based on input type (.f16/.f32/.f64)
    // ==========================================
    func sin(float32) float32
    func cos(float32) float32
    func tan(float32) float32
    func pow(float32, float32) float32
    func log2(float32) float32
    func sqrt(float32) float32
    func rsqrt(float32) float32     // 1.0 / sqrt(x)
    func rcp(float32) float32       // 1.0 / x
    
    func abs(float32) float32
    func abs(int32) int32
    func saturate(float32) float32  // Clamp [0.0, 1.0]
    func fma(float32, float32, float32) float32 // (a*b)+c
    
    // ==========================================
    // 3. Bit Manipulation
    // ==========================================
    func popc(uint32) int32         // Population count
    func clz(uint32) int32          // Count leading zeros
    func ffs(uint32) int32          // Find first set bit
    func brev(uint32) uint32        // Bit reverse

    // ==========================================
    // 4. Warp Control & Voting
    // ==========================================
    func sync_threads()             // bar.sync 0
    func sync_warp()                // bar.warp.sync
    func active_mask() uint32       // activemask
    
    // Voting (Predicate Logic)
    func all(bool) bool             // vote.all.pred
    func any(bool) bool             // vote.any.pred
    func ballot(bool) uint32        // vote.ballot.b32
    
    // Shuffling (Register Exchange)
    // <T> resolves to .b32/.b64 depending on type width
    func shuffle<T>(T, int32) T             // shfl.sync.idx
    func shuffle_up<T>(T, int32) T          // shfl.sync.up
    func shuffle_down<T>(T, int32) T        // shfl.sync.down
    func shuffle_xor<T>(T, int32) T         // shfl.sync.bfly

    // ==========================================
    // 5. Atomic Operations
    // ==========================================
    func atomic_add<T>(*T, T) T     // Returns OLD value
    func atomic_sub<T>(*T, T) T
    func atomic_min<T>(*T, T) T
    func atomic_max<T>(*T, T) T
    func atomic_exchange<T>(*T, T) T
    func atomic_cas<T>(*T, T, T) T  // (ptr, compare, val)

    // ==========================================
    // 6. Memory & Async Pipeline
    // ==========================================
    // Fences
    func thread_fence()             // membar.gl
    func thread_fence_block()       // membar.cta
    func thread_fence_system()      // membar.sys
    
    // Load Caching Hints
    func load_cached<T>(*T) T       // ld.global.ca
    func load_streaming<T>(*T) T    // ld.global.cs
    func load_volatile<T>(*T) T     // ld.volatile
    
    // Async Copy (Global -> Shared)
    func async_copy(*void, *void, usize) // cp.async.ca.shared.global
    func async_commit_group()            // cp.async.commit_group
    func async_wait_group(int32)         // cp.async.wait_group

    // ==========================================
    // 7. Tensor Memory Accelerator (H100+)
    // ==========================================
    func tma_copy(gpu.TensorMap, *void, gpu.Barrier) // cp.async.bulk.tensor
    func barrier_wait(gpu.Barrier)                   // mbarrier.test_wait
}
```

---

## 5. Implementation Strategy

### Backend Logic (`pkg/ptxgen`)

When the backend visits a `CallExpr` node:

1.  **Check Flag:** Is `node.Func.IsExternGPU` true?
2.  **Lookup:** Find the handler in `IntrinsicMap[node.Func.Name]`.
3.  **Generate:** Execute the handler to emit the string.

```go
// Conceptual Go Implementation
var IntrinsicHandlers = map[string]IntrinsicGenerator{
    "sin": func(gen *Generator, args []Arg, res Reg) {
        // Automatically handle precision based on arg type
        suffix := ".f32"
        if args[0].Type == Float64 { suffix = ".f64" }
        gen.Emit("sin.approx%s %s, %s;", suffix, res, args[0])
    },
    
    "shuffle": func(gen *Generator, args []Arg, res Reg) {
        // Handle the 5-operand PTX shuffle instruction
        // shfl.sync.idx.b32 d, a, b, c, mask;
        mask := "0x1f"
        gen.Emit("shfl.sync.idx.b32 %s, %s, %s, %s, 0xffffffff;", 
            res, args[0], args[1], mask)
    },
}
```