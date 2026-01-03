# Arc Language Execution Model

Arc separates **logic definition** from **execution context**. Write your code as inline functions, then choose where to run them: async event loop, OS threads, isolated processes, sandboxed containers, or distributed cloud instances.

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

## The 5 Execution Models

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
// Custom config with resource limits
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
    allowed_syscalls: array<int>
}
```

**Isolation:** filesystem, network, PIDs, resource limits (cgroups/namespaces).

---

### 5. `cloud` - Distributed & GPU Computing

**When to use:** Massive parallelism, GPU training, accessing private VPCs, self-deploying infrastructure.

**Cost:** ~1-3s setup (cold), <100ms (warm/p2p)

**Mechanism:**
- **Ad-hoc:** Compiles closure + captured vars + Arc runtime to a micro-payload.
- **WebRTC:** Uses Data Channels for bidirectional streaming (logs/results) without opening firewalls (NAT traversal).
- **Detached:** Allows "deploy and die" patterns.

**Syntax:**
```arc
// 1. GPU Compute (RPC style)
let model = cloud func(data: bytes) Model {
    return ai.train(data)
}(CloudConfig{ 
    provider: "gcp", 
    hardware: "nvidia-a100",
    transport: "webrtc" // No SSH keys needed
}, dataset)

// 2. Self-Deploying Server (Detached)
cloud func() {
    let s = http.listen(":80")
    s.start() // Runs forever in the cloud
}(CloudConfig{ 
    provider: "aws", 
    region: "us-east-1", 
    detached: true // Program exits, server stays up
})
```

**Config:**
```arc
struct CloudConfig {
    provider: string     // "aws", "gcp", "digitalocean", "ssh"
    region: string
    hardware: string     // "cpu-basic", "gpu-t4", "ram-128gb"
    detached: bool       // If true, local process disconnects after launch
    transport: string    // "webrtc", "http", "ssh"
    auth: AuthConfig     // SigV4, OIDC, or Keys
    vpc: string          // Target private network
}
```

---

## Quick Comparison

| Model | Isolation | Use Case | Overhead | Config |
|-------|-----------|----------|----------|--------|
| `async` | Shared memory | I/O-bound | ~200ns | No |
| `thread` | Shared memory | CPU-bound, blocking | ~5µs | Optional |
| `process` | Separate memory | Fault tolerance | ~5ms | Optional |
| `container` | Sandboxed | Security, limits | ~50ms | Optional |
| `cloud` | **Physical/VM** | **GPU, Deploy, Scale** | **~2s** | **Required** |

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

**Distributed "Botnet" Loop (cloud)**
```arc
// Implicit Massive Parallelism
// Spawns 1000 serverless instances instantly
for chunk in dataset.split(1000) {
    cloud func(data) { 
        // Built-in modules (math/media) available instantly
        return heavy_math.process(data) 
    }(CloudConfig{ provider: "aws_lambda" }, chunk)
}
```

**Self-Deploying Infrastructure (cloud detached)**
```arc
func deploy() {
    print("Deploying to production...")
    
    cloud func() {
        // This runs on a remote VM forever
        app.start_server()
    }(CloudConfig{ 
        detached: true, 
        name: "prod-api-v1" 
    })
    
    print("Deployed. Exiting.")
}
```

---

## Philosophy

- **async**: Lightweight concurrency
- **thread**: Real parallelism  
- **process**: Isolation
- **container**: Security
- **cloud**: Infinite Scale & Deployment

All inline functions. Zero boilerplate. From a single core to a global cluster with one syntax.