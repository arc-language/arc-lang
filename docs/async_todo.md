
# ASYNC & RUNTIME REFACTOR TODO

For the Standard Library: If you need specific CPU instructions (like CPUID or RDTSC), add them as Compiler Intrinsics.
User writes: std.cpu.rdtsc()
Compiler sees: OpCall("std.cpu.rdtsc") -> Maps directly to 0F 31 bytes.
Benefit: The compiler knows exactly which registers are used (RAX/RDX), so register allocation still works perfectly.



## 1. Core Philosophy Change
**Del** using `__cpu_xxx` string matching and inline assembly in source files for core runtime features.
**START** using Go's strong struct packing capabilities to generate the runtime environment natively.

The compiler (written in Go) should act as the ultimate factory. It has access to `encoding/binary` and strong typing. We should leverage that to build the "Foundation" of the language (Threading, Signals, Memory) directly in machine code, rather than trying to bootstrap it via "Unmanaged Blocks" in self-hosting source code.

**The Split:**
1.  **Compiler (Go):** Generates Threading, Context Switching, Signals, Syscall ABIs.
2.  **Standard Library (Arc Source):** Handles Sockets, HTTP, File I/O (using the primitives provided by step 1).

---

## 2. Architecture: The 3-Package Solution

We are moving away from the monolithic `compiler.go` loop. We need to split the backend into three distinct phases to handle complex layouts and runtime logic.

### A. Package `layout` (or `abi`)
*Goal: Solve "Building Structs on the Fly" without manual offset math.*

Instead of calculating `offset += 8` manually in the codegen loop, define the OS ABI in Go structs.

*   **Action:** Create `backend/amd64/linux/abi.go`
*   **Action:** Define Linux Kernel structures as Go structs with `WriteTo(buf)` methods.

**Example (SigAction):**
```go
// In package linux_abi
type SigAction struct {
    Handler  uint64
    Flags    uint64
    Restorer uint64
    Mask     uint64
}

func (s *SigAction) WriteBytes(buf *bytes.Buffer) {
    // Go handles the endianness and type widths
    binary.Write(buf, binary.LittleEndian, s)
    // Go handles the padding logic here once, correct forever
    buf.Write(make([]byte, 152 - currentLen)) 
}
```

### B. Package `runtime` (or `gen`)
*Goal: Encapsulate complex assembly sequences.*

The main compiler loop should not know how `CLONE_VM` works or how to clear registers for a child thread.

*   **Action:** Create `backend/amd64/gen/threading.go`
*   **Action:** Move the logic for `__cpu_clone_spawn` out of `compileInst`.
*   **Action:** Create methods like `EmitAsyncTrampoline`, `EmitSegfaultHandler`.

### C. Package `ir` (Refactor)
*Goal: Native support for Concurrency.*

*   **Action:** Remove `CallInst` hacks for `__cpu_`.
*   **Action:** Add real OpCodes:
    *   `OpAsyncTaskCreate` (replaces clone intrinsics)
    *   `OpAsyncTaskAwait`
    *   `OpYield`

---

## 3. Implementation Plan

### Step 1: The Linux Syscall ABI
Stop guessing offsets. Implement the `layout` package to strictly define:
*   `struct sigaction` (152 bytes)
*   `ucontext_t` (for signal stack walking)
*   `stack_t` (for sigaltstack)

**Why:** When the compiler initializes, it will use `layout` to write these bytes into the `.data` section immediately. We pass pointers to these data blocks to the syscalls.

### Step 2: The "Generator" Pattern
Create a `ThreadingBackend` struct that holds a reference to the `Assembler`.

```go
// codegen/gen/threading.go
func (t *ThreadingBackend) SpawnThread(fnPtr Reg, stackPtr Reg) {
    // 1. Emit Syscall (Clone)
    // 2. Emit Parent/Child Branching
    // 3. Emit Child Stack Setup (Zeroing RBP)
    // 4. Emit Trampoline to User Function
}
```

### Step 3: Clean up `compileInst`
The main loop becomes a simple dispatcher.

**Old (Bad):**
```go
case ir.OpCall:
   if name == "__cpu_clone_spawn" { ... 50 lines of complex assembly ... }
```

**New (Good):**
```go
case ir.OpAsyncTaskCreate:
   // The compiler knows exactly what this is. 
   // It delegates to the generator to output the correct machine code sequence.
   c.threading.SpawnThread(inst.Callee, inst.Stack)
```

---

## 4. Summary of Responsibilities

| Feature | Where it lives now | Where it belongs |
| :--- | :--- | :--- |
| **TCP/HTTP/JSON** | Source Code | **Source Code (Stdlib)** |
| **Logic (loops, if)** | Compiler Loop | **Compiler Loop** |
| **Syscall Structs** | Manual byte writing | **`layout` Package** |
| **Thread Creation** | `__cpu` Intrinsic | **`runtime` Generator + Native IR** |
| **Signal Handling** | Inline Assembly | **`runtime` Generator** |

## 5. Notes on "Unmanaged Blocks"

While `unmanaged` functions in source code are clever, they introduce circular dependencies and fragility when defining the core runtime.

By moving the "Foundation" into the Go-based compiler using the `layout` + `gen` approach:
1.  We get **compile-time safety** on struct sizes (via Go).
2.  We generate **optimal machine code** without function call overhead.
3.  We avoid the "chicken and egg" problem of self-hosting the runtime.

