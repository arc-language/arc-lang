# Arc Language Execution Model

Arc separates **logic definition** from **execution context**. Write your code once as a function, then choose where to run it: cooperative threads, OS threads, isolated processes, or sandboxed containers.

This allows Arc to scale from lightweight concurrency to security isolation—all with the same language and syntax.

---

## The 4 Execution Models

### 1. `spawn` - Green Threads (Cooperative Multitasking)

**What:** Lightweight coroutines managed by Arc's runtime scheduler. Runs on the event loop.

**When to use:**
- High-concurrency I/O (web servers, 10k+ connections)
- Async operations (network, file I/O)
- Tasks that yield frequently

**Cost:** ~200ns switching, ~200 bytes memory

**Syntax:**
```arc
// Inline anonymous function
spawn func(url: string) {
    let data = await http.get(url)
    process(data)
}("https://api.example.com")

// Named function
func fetch_data(url: string) {
    let data = await http.get(url)
    process(data)
}

let handle = spawn fetch_data("https://api.example.com")
handle.await  // Wait for completion
```

**Restrictions:**
- Must use non-blocking I/O (async APIs)
- Cannot call blocking C functions (use `thread` instead)

---

### 2. `thread` - OS Threads (Preemptive Multitasking)

**What:** Real OS threads managed by the kernel. Each thread has its own stack.

**When to use:**
- Blocking C library calls (libc, database drivers)
- CPU-intensive work that shouldn't block the event loop
- True parallel computation on multiple cores

**Cost:** ~5µs switching, ~1MB stack

**Syntax:**
```arc
// Call blocking C function safely
func blocking_work(path: string) {
    let file = libc.fopen(path, "r")
    libc.sleep(1000)  // Blocks this thread only
    libc.fclose(file)
}

let handle = thread blocking_work("/tmp/data.txt")
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
// Run risky operation in isolated process
func risky_task(data: *byte) int32 {
    // If this crashes, parent process is safe
    let result = dangerous_computation(data)
    return result
}

let handle = process risky_task(data_ptr)
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
// Run untrusted code in sandbox
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
| `spawn` | Shared memory | Cooperative | I/O-bound | ~200ns |
| `thread` | Shared memory | Preemptive | CPU-bound, blocking calls | ~5µs |
| `process` | Separate memory | Full | Fault tolerance | ~1-10ms |
| `container` | Sandboxed + limits | Full | Security, multi-tenant | ~10-50ms |

---

## Examples

### Web Server (spawn)
```arc
async func handle_request(req: Request) Response {
    let data = await db.query("SELECT * FROM users")
    return Response{body: data}
}

func main() {
    let server = http.listen(":8080")
    for req in server.accept() {
        spawn handle_request(req)  // 10k+ concurrent requests
    }
}
```

### Database Driver (thread)
```arc
func query_blocking_db(sql: string) Result {
    let conn = libc.connect_db()  // Blocking C call
    let result = libc.execute(conn, sql)  // Blocks
    return result
}

let handle = thread query_blocking_db("SELECT * FROM large_table")
let result = handle.join()
```

### Plugin System (process)
```arc
func load_plugin(path: string) {
    let handle = process func(p: string) {
        let plugin = load_library(p)
        plugin.run()
    }(path)
    
    // If plugin crashes, main process continues
    if handle.wait() != 0 {
        io.printf("Plugin crashed\n")
    }
}
```

### Multi-Tenant Execution (container)
```arc
func run_user_code(code: string, user_id: int32) {
    let config = ContainerConfig{
        cpu_limit: 10,           // 10% CPU
        memory_limit: 64_MB,
        timeout: 5_000,          // 5 seconds
        network: false
    }
    
    let handle = container func(c: string) {
        eval_and_run(c)
    }(code).with_config(config)
    
    let result = handle.wait_timeout(5_000)
}
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
- **spawn**: Lightweight, everywhere
- **thread**: When you need real parallelism
- **process**: When you need isolation
- **container**: When you need security

All with the same language, same syntax, same binary. No external dependencies, no frameworks, no runtime bloat.