# Arc Language Execution Model

Arc separates **logic definition** from **execution context**. Write your code as inline functions, then choose where to run them: async event loop, OS threads, isolated processes, or sandboxed containers.

This allows Arc to scale from lightweight concurrency to security isolation—all with the same language and syntax.

---

## The 4 Execution Models

### 1. `async` - Async Functions (Cooperative Multitasking)

**What:** Asynchronous functions that run on the event loop. Execution is cooperative and non-blocking.

**When to use:**
- High-concurrency I/O (web servers, 10k+ connections)
- Network requests, file I/O
- Tasks that yield frequently

**Cost:** ~200ns context switch, minimal memory overhead

**Syntax:**
```arc
let result = await async func(url: string) {
    let data = await http.get(url)
    return process(data)
}("https://api.example.com")
```

**Restrictions:**
- Must use non-blocking I/O (async APIs)
- Cannot call blocking C functions (use `thread` instead)

---

### 2. `thread` - OS Threads (Preemptive Multitasking)

**What:** Real OS threads managed by the kernel. Each thread has its own stack.

**When to use:**
- Blocking C library calls (libc, database drivers)
- CPU-intensive work
- True parallel computation on multiple cores

**Cost:** ~5µs switching, ~1MB stack

**Syntax:**
```arc
let handle = thread func(path: string) {
    let file = libc.fopen(path, "r")
    libc.sleep(1000)  // Blocks this thread only
    libc.fclose(file)
}("/tmp/data.txt")

handle.join()  // Wait for thread to finish
```

---

### 3. `process` - OS Processes (Memory Isolation)

**What:** Separate OS process with isolated memory space. Uses fork/clone syscall.

**When to use:**
- Fault tolerance (crashes don't affect parent)
- Plugins or untrusted code (limited isolation)
- Tasks that need complete memory separation

**Cost:** Milliseconds setup, copy-on-write memory

**Syntax:**
```arc
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

### 4. `container` - Sandboxed Processes (Security Isolation)

**What:** Process with Linux namespaces and cgroups. Isolated network, filesystem, and PID view.

**When to use:**
- Security-critical tasks (user-submitted code)
- Multi-tenant execution
- "Serverless" functions inside your binary
- Limiting resource usage (CPU, memory, network)

**Cost:** Milliseconds setup

**Syntax:**
```arc
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
| `async` | Shared memory | Cooperative | I/O-bound, high concurrency | ~200ns |
| `thread` | Shared memory | Preemptive | CPU-bound, blocking calls | ~5µs |
| `process` | Separate memory | Full | Fault tolerance | ~1-10ms |
| `container` | Sandboxed + limits | Full | Security, multi-tenant | ~10-50ms |

---

## Examples

### Web Server (async)
```arc
func main() {
    let server = http.listen(":8080")
    for req in server.accept() {
        let response = await async func(r: Request) Response {
            let data = await db.query("SELECT * FROM users")
            return Response{body: data}
        }(req)
        
        await req.send(response)
    }
}
```

### Parallel Async Operations (async)
```arc
let user_id = 42

let results = await all([
    async func() { return await fetch_user(user_id) }(),
    async func() { return await fetch_posts(user_id) }(),
    async func() { return await fetch_comments(user_id) }()
])
```

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
- **async**: Lightweight concurrency, I/O-bound
- **thread**: Real parallelism, CPU-bound
- **process**: Isolation and fault tolerance
- **container**: Security and resource limits

All with the same language, same syntax, same binary. No external dependencies, no frameworks, no runtime bloat.

**All execution contexts use inline functions only** - making execution boundaries explicit, data flow clear, and preventing hidden complexity.