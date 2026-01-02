# Arc Language Execution Model

Arc separates **logic definition** from **execution context**. Write your code as inline functions, then choose where to run them: async event loop, OS threads, isolated processes, or sandboxed containers.

---

## Configuration Pattern

Execution contexts support **config-first arguments** (except `async`):

```arc
context func(params) { body }(Config{...}, args...)
```

**The compiler detects config at compile time:**
- If first argument is the config type → use it as config
- If first argument matches first parameter type → use default config
- Zero runtime overhead

```arc
// With custom config
thread func(x: int) { work(x) }(ThreadConfig{stack_size: 8_MB}, 42)

// With default config
thread func(x: int) { work(x) }(42)
```

**Note:** `async` has no config - arguments always map directly to parameters.

---

## The 4 Execution Models

### 1. `async` - Async Functions

**When to use:** High-concurrency I/O, network requests, non-blocking operations

**Cost:** ~200ns context switch

**Syntax:**
```arc
let result = await async func(url: string) {
    let data = await http.get(url)
    return process(data)
}("https://api.example.com")
```

**No config available** - always uses event loop defaults.

---

### 2. `thread` - OS Threads

**When to use:** Blocking C calls, CPU-intensive work, parallel computation

**Cost:** ~5µs switching, ~1MB stack

**Syntax:**
```arc
// Default config
let handle = thread func(path: string) {
    let file = libc.fopen(path, "r")
    libc.fclose(file)
}("/tmp/data.txt")

handle.join()

// Custom config
let handle = thread func(data: array<int>) {
    compute(data)
}(ThreadConfig{stack_size: 8_MB, cpu_affinity: array<int>{0, 1}}, dataset)
```

**Config:**
```arc
struct ThreadConfig {
    stack_size: usize
    cpu_affinity: array<int>
    priority: int
    name: string
}
```

---

### 3. `process` - OS Processes

**When to use:** Fault tolerance, plugins, memory isolation

**Cost:** ~1-10ms setup

**Syntax:**
```arc
// Default config
let handle = process func(data: *byte) int32 {
    return dangerous_computation(data)
}(data_ptr)

let result = handle.wait()

// Custom config
let handle = process func(script: string) {
    run_script(script)
}(ProcessConfig{env: map<string, string>{"PATH": "/bin"}, stdout: log_file}, script_path)
```

**Config:**
```arc
struct ProcessConfig {
    env: map<string, string>
    working_dir: string
    stdin: FileDescriptor
    stdout: FileDescriptor
    stderr: FileDescriptor
    rlimits: RLimits
}
```

---

### 4. `container` - Sandboxed Processes

**When to use:** Security-critical tasks, multi-tenant execution, resource limits

**Cost:** ~10-50ms setup

**Syntax:**
```arc
// Default config
let handle = container func(code: string) {
    eval(code)
}(user_code)

// Custom config
let handle = container func(code: string) {
    eval(code)
}(ContainerConfig{
    cpu_limit: 50,
    memory_limit: 128_MB,
    network: false
}, user_code)
```

**Config:**
```arc
struct ContainerConfig {
    cpu_limit: int
    memory_limit: usize
    network: bool
    readonly_fs: bool
    timeout: int
    allowed_syscalls: array<int>
}
```

**Isolation:** filesystem, network, PIDs, resource limits

---

## Quick Comparison

| Model | Isolation | Use Case | Overhead | Config |
|-------|-----------|----------|----------|--------|
| `async` | Shared memory | I/O-bound | ~200ns | No |
| `thread` | Shared memory | CPU-bound, blocking | ~5µs | Optional |
| `process` | Separate memory | Fault tolerance | ~1-10ms | Optional |
| `container` | Sandboxed | Security, limits | ~10-50ms | Optional |

---

## Examples

**Web Server (async)**
```arc
let server = http.listen(":8080")
for req in server.accept() {
    let response = await async func(r: Request) Response {
        let data = await db.query("SELECT * FROM users")
        return Response{body: data}
    }(req)
}
```

**Parallel Computation (thread)**
```arc
let h1 = thread func(data: array<int>) { return compute(data) }(dataset_a)
let h2 = thread func(data: array<int>) { return compute(data) }(dataset_b)

let result = h1.join() + h2.join()
```

**Plugin System (process)**
```arc
let handle = process func(path: string) {
    let plugin = load_library(path)
    plugin.run()
}("/path/to/plugin.so")

if handle.wait() != 0 {
    io.printf("Plugin crashed\n")
}
```

**Untrusted Code (container)**
```arc
let handle = container func(code: string) {
    eval(code)
}(ContainerConfig{
    cpu_limit: 10,
    memory_limit: 64_MB,
    network: false
}, user_code)

let result = handle.wait()
```

---

## Philosophy

- **async**: Lightweight concurrency
- **thread**: Real parallelism  
- **process**: Isolation
- **container**: Security

All inline functions. Config when needed. Zero overhead. Same language, same syntax.