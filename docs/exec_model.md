# Arc Language Execution Model

This document outlines the architecture for Arc's 4-Tier Execution Model. The design philosophy separates **Logic Definition** (`async func`) from **Execution Context** (`spawn`, `thread`, `process`, `container`).

This allows the language to scale from a single concurrent web server to secure, isolated sandboxes without changing the core syntax or relying on heavy external dependencies.

---

## The 4 Tiers of Execution

### 1. `spawn` (Green Threads)
*   **Mechanism:** User-space cooperative multitasking (Stackless Coroutines).
*   **Cost:** ~200ns switching, ~200 bytes memory.
*   **Under the hood:** Runs on a single OS thread (Main Loop). Uses `coro.suspend` and `coro.resume`.
*   **Use Case:** High-concurrency I/O, Web Servers (10k+ connections), Business Logic.
*   **Syntax:** `let h = spawn my_logic()`

### 2. `thread` (OS Threads)
*   **Mechanism:** Kernel-managed preemptive threading (`clone` syscall with shared VM).
*   **Cost:** ~5µs switching, ~1MB stack (configurable).
*   **Under the hood:** Maps 1:1 to a real OS thread.
*   **Use Case:** Blocking C functions (`libc.sleep`, DB drivers), CPU-heavy math that would freeze the Event Loop.
*   **Syntax:** `let h = thread my_logic()` or `thread func() { ... }(stack_size)`

### 3. `process` (OS Processes)
*   **Mechanism:** Kernel-managed memory isolation (`fork` syscall / `clone` without shared VM).
*   **Cost:** Milliseconds setup, Copy-on-Write memory.
*   **Communication:** Pipes / IPC (Cannot share Heap).
*   **Use Case:** Crash resilience, Fault tolerance, Plugins.
*   **Syntax:** `let h = process my_logic()`

### 4. `container` (Sandboxed Processes)
*   **Module Builder:** Works with build.arc files to build modules that need c or c++ deps.
*   **Mechanism:** Linux Namespaces & Cgroups (`clone` with `CLONE_NEWPID`, `CLONE_NEWNET`, etc.).
*   **Cost:** Milliseconds.
*   **Capabilities:** Can isolate Network, Filesystem (`chroot`), and PID view (looks like PID 1).
*   **Use Case:** Security testing, Multi-tenant execution, "Serverless" logic inside the binary.
*   **Syntax:** `let h = container func() { ... }(config)`

---

## Implementation Roadmap

### Phase 1: The Foundation (Async/Await & Coroutines)

To enable Green Threads (`spawn`) and the logic for all other tiers, we must implement **LLVM-style Coroutine Intrinsics** in our custom backend.

**1. Parser & Semantics:**
*   Add support for `intrinsic` keyword or map specific `extern` function names (e.g., `__coro_resume`) to internal IR operations.
*   Update `VisitFunctionDecl` to handle `async`. If `async`, the function must allocate a frame, return a Handle (`ptr<i8>`), and setup the cleanup block.
*   Update `VisitReturnStmt` to store results in the Promise slot and jump to cleanup (Suspend Final) instead of returning normally.

**2. IR Generation:**
*   Implement the custom coroutine instructions in `builder`:
    *   `OpCoroId`, `OpCoroBegin`, `OpCoroSuspend`, `OpCoroResume`, `OpCoroDestroy`, `OpCoroPromise`.

**3. Codegen (AMD64 Backend):**
*   **Manual Assembly Generation:** Since we aren't using LLVM libraries, we implement the "Context Switch" manually in `amd64/coroutines.go`.
    *   **Suspend:** Save callee-saved registers (RBX, RBP, R12-15), RIP, and RSP to the heap frame.
    *   **Resume:** Load registers/RSP from the heap frame and JMP to the saved RIP.
    *   **Begin:** `mmap` a stack chunk/frame.

### Phase 2: The Runtime (Pure Arc)

We build the Scheduler and Thread Pool in `lib/runtime.arc` using the primitives above.

*   **Scheduler:** A simple Queue of Handles. The Main Loop pops a handle and calls `__coro_resume(h)`.
*   **Thread Pool:** Uses `syscall_clone` to create 4-8 dumb OS threads that wait for function pointers (for the `thread` tier).
*   **Container Helper:** Wraps `syscall_clone` with namespace flags to implement the `container` tier.

### Phase 3: Core Modules (Non-Blocking I/O)

Because Green Threads (`spawn`) share the main OS thread, we cannot use blocking syscalls (like `libc.read`) inside standard library modules, or the whole server freezes.

*   **The "Dance":** Core modules (`std.net`, `std.io`) must be rewritten to be async-aware.
    1.  Try Non-Blocking Read.
    2.  If `EAGAIN`: Register File Descriptor with the Runtime Poller (using `poll` syscall).
    3.  `await suspend()` (Yield to scheduler).
    4.  Resume when Poller says data is ready.
*   **Ease of Implementation:** Since we have `async` and `await` keywords working, writing this "dance" in `lib/std/*.arc` is straightforward and readable, unlike C-based callbacks.

---

## Summary

We are avoiding the "One size fits all" trap. By treating `async` as a state-machine transform and separating it from the runner (`spawn` vs `thread`), we get:

1.  **Massive Scalability** (Green Threads/Spawn).
2.  **Compatibility** (OS Threads for blocking C).
3.  **DevOps Powers** (Inline Containers).

This architecture provides a professional-grade runtime environment entirely self-hosted in Arc.