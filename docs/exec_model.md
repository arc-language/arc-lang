# Arc Language Execution Model

Arc separates **logic definition** from **execution context**. Write your code as inline functions, then choose where to run them: OS threads, isolated processes, or sandboxed containers.

This allows Arc to scale from parallel computation to security isolation—all with the same language and syntax.

---

## The 3 Execution Models

### 1. `thread` - OS Threads (Preemptive Multitasking)

**What:** Real OS threads managed by the kernel. Each thread has its own stack.

**When to use:**
- Blocking C library calls (libc, database drivers)
- CPU-intensive work
- True parallel computation on multiple cores

**Cost:** ~5µs switching, ~1MB stack

**Syntax:**
```arc
// Inline anonymous function only
let handle = thread func(path: string) {
    let file = libc.fopen(path, "r")
    libc.sleep(1000)  // Blocks this thread only
    libc.fclose(file)
}("/tmp/data.txt")

handle.join()  // Wait for thread to finish
```

---

### 2. `process` - OS Processes (Memory Isolation)

**What:** Separate OS process with isolated memory space. Uses fork/clone syscall.

**When to use:**
- Fault tolerance (crashes don't affect parent)
- Plugins or untrusted code (limited isolation)
- Tasks that need complete memory separation

**Cost:** Milliseconds setup, copy-on-write memory

**Syntax:**
```arc
// Inline anonymous function only
let handle = process func(data: *byte) int32 {
    // If this crashes, parent process is safe
    let result = dangerous_computation(data)
    return result
}(data_ptr)

let result = handle.wait()  // Blocks until process exits

// Check exit status
if handle.exit_code() != 0 {
    io.printf("Process crashed\n")
}
```

**Communication:**
- Processes share nothing (no shared heap)
- Use pipes, sockets, or shared memory for IPC

---

### 3. `container` - Sandboxed Processes (Security Isolation)

**What:** Process with Linux namespaces and cgroups. Isolated network, filesystem, and PID view.

**When to use:**
- Security-critical tasks (user-submitted code)
- Multi-tenant execution
- "Serverless" functions inside your binary
- Limiting resource usage (CPU, memory, network)

**Cost:** Milliseconds setup

**Syntax:**
```arc
// Inline anonymous function only
let handle = container func(code: string) {
    // Isolated filesystem (chroot)
    // Isolated network (own network namespace)
    // CPU/memory limits enforced by cgroups
    let result = eval(code)
    return result
}("user_code_here")

// Configure container limits
let config = ContainerConfig{
    cpu_limit: 50,        // 50% of one core
    memory_limit: 128_MB,
    network: false,       // No network access
    readonly_fs: true
}

let handle = container func() {
    // Sandboxed execution
}(config)
```

**Isolation provided:**
- Filesystem (chroot/pivot_root)
- Network (separate network namespace)
- PIDs (appears as PID 1 inside container)
- Resource limits (cgroups)

---

## Quick Comparison

| Model | Isolation | Concurrency | Use Case | Overhead |
|-------|-----------|-------------|----------|----------|
| `thread` | Shared memory | Preemptive | CPU-bound, blocking calls | ~5µs |
| `process` | Separate memory | Full | Fault tolerance | ~1-10ms |
| `container` | Sandboxed + limits | Full | Security, multi-tenant | ~10-50ms |

---

## Examples

### Database Driver (thread)
```arc
let handle = thread func(sql: string) Result {
    let conn = libc.connect_db()  // Blocking C call
    let result = libc.execute(conn, sql)  // Blocks
    return result
}("SELECT * FROM large_table")

let result = handle.join()
```

### Plugin System (process)
```arc
let handle = process func(path: string) {
    let plugin = load_library(path)
    plugin.run()
}("/path/to/plugin.so")

// If plugin crashes, main process continues
if handle.wait() != 0 {
    io.printf("Plugin crashed\n")
}
```

### Multi-Tenant Execution (container)
```arc
let config = ContainerConfig{
    cpu_limit: 10,           // 10% CPU
    memory_limit: 64_MB,
    timeout: 5_000,          // 5 seconds
    network: false
}

let handle = container func(code: string) {
    eval_and_run(code)
}(user_code).with_config(config)

let result = handle.wait_timeout(5_000)
```

---

## GPU Execution

For GPU/CUDA execution, see the separate **GPU Execution Model** documentation which covers:
- `async func<gpu>` syntax for GPU kernels
- `await` and `await(device)` for device selection
- Multi-GPU programming
- PTX compilation and CUDA Driver API integration

---

## Philosophy

Arc's execution model gives you the right tool for every job:
- **thread**: When you need real parallelism
- **process**: When you need isolation
- **container**: When you need security

All with the same language, same syntax, same binary. No external dependencies, no frameworks, no runtime bloat.