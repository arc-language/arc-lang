# Compute Templates (Version 1.0)

Compute templates allow you to define custom execution contexts. Mark a struct with `<compute>` to make it usable with the `YourContext func() {}(args)` syntax pattern.

This feature is module-aware, allowing the community to build and share execution models (Docker, Kubernetes, AWS Lambda, ML Tasks) just like standard libraries.

---


## Usage 

```arc
import "github.com/someuser/mydocker2"

func main() {
    print("Orchestrating job...")

    // Usage: Namespace.Type func()
    let job = mydocker2.Container func(x: int) int {
        
        // This code runs INSIDE the container
        // Dependencies linked in the binary are available automatically
        return heavy_math(x)
        
    }(mydocker2.Config{image: "alpine:latest"}, 100)

    // Wait for result (IPC handled by the library struct)
    let result = job.wait()
    
    print("Result from Docker: ${result}")
    job.stop()
}
```



## Compute Templates, basic definition

```arc
// Define a custom compute context
struct custom_thread<compute> {
    handle: ThreadHandle
    stack_ptr: *byte
    
    // Special init - called when: custom_thread func() {}(args)
    init(self ctx: *custom_thread<compute>, closure: func(int) int, arg: int) {
        ctx.stack_ptr = alloc_stack(2_MB)
        ctx.handle = syscall_clone(closure, ctx.stack_ptr, arg)
    }
    
    func join(self ctx: *custom_thread<compute>) int {
        let result = syscall_join(ctx.handle)
        free(ctx.stack_ptr)
        return result
    }
}

// Usage - looks like built-in contexts
let t = custom_thread func(x: int) int { 
    return heavy_work(x) 
}(42)

let result = t.join()
```

---

## Compute Templates, flat method style

```arc
// Struct definition
struct custom_thread<compute> {
    handle: ThreadHandle
    config: ThreadConfig
}

// Special init for compute contexts (must use <compute> in self type)
init(self ctx: *custom_thread<compute>, closure: func(int) int, cfg: ThreadConfig, arg: int) {
    ctx.config = cfg
    ctx.handle = spawn_with_config(closure, cfg, arg)
}

// Regular methods
func join(self ctx: *custom_thread<compute>) int {
    return wait_thread(ctx.handle)
}

func detach(self ctx: *custom_thread<compute>) {
    detach_thread(ctx.handle)
}

// Usage
let t = custom_thread func(x: int) { 
    work(x) 
}(ThreadConfig{stack: 8_MB, priority: 10}, 42)

t.join()
```

---

## Compute Templates, built-in contexts

### thread - OS Threads

```arc
struct thread<compute> {
    handle: ThreadHandle
    stack_ptr: *byte
    
    init(self t: *thread<compute>, closure: func(T) R, config: ThreadConfig, args: T) {
        t.stack_ptr = alloc_stack(config.stack_size)
        t.handle = syscall_clone(closure, t.stack_ptr, args)
    }
    
    func join(self t: *thread<compute>) R {
        let result = syscall_join(t.handle)
        free(t.stack_ptr)
        return result
    }
    
    func detach(self t: *thread<compute>) {
        syscall_detach(t.handle)
    }
}

// Usage
let t = thread func(data: array<int>) int {
    return compute(data)
}(ThreadConfig{stack_size: 8_MB}, dataset)

let result = t.join()
```

**Config:**
```arc
struct ThreadConfig {
    stack_size: usize      // Default: 1MB
    cpu_affinity: array<int>
    priority: int
    name: string
}
```

---

### process - OS Processes

```arc
struct process<compute> {
    pid: int
    ipc_channel: FileDescriptor
    
    init(self p: *process<compute>, closure: func(T) R, config: ProcessConfig, args: T) {
        // Serialize closure and captured variables
        let payload = serialize_closure(closure, args)
        
        // Fork process
        p.pid = syscall_fork()
        
        if p.pid == 0 {
            // Child process
            apply_config(config)
            let result = deserialize_and_execute(payload)
            ipc_send(result)
            exit(0)
        } else {
            // Parent process
            p.ipc_channel = setup_ipc(p.pid)
        }
    }
    
    func wait(self p: *process<compute>) R {
        let result = ipc_recv(p.ipc_channel)
        syscall_waitpid(p.pid)
        return result
    }
    
    func kill(self p: *process<compute>) {
        syscall_kill(p.pid, SIGKILL)
    }
}

// Usage
let p = process func(data: *byte) int {
    return dangerous_computation(data)
}(ProcessConfig{env: {"PATH": "/bin"}}, data_ptr)

let result = p.wait()
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

### container - Sandboxed Processes

```arc
struct container<compute> {
    pid: int
    cgroup_path: string
    namespace_fds: array<int>
    
    init(self c: *container<compute>, closure: func(T) R, config: ContainerConfig, args: T) {
        // Serialize closure
        let payload = serialize_closure(closure, args)
        
        // Create namespaces
        c.namespace_fds = create_namespaces()
        
        // Setup cgroups
        c.cgroup_path = setup_cgroups(config.cpu_limit, config.memory_limit)
        
        // Fork with namespaces
        c.pid = syscall_clone(CLONE_NEWNS | CLONE_NEWPID | CLONE_NEWNET)
        
        if c.pid == 0 {
            // Child process in container
            apply_seccomp_filter(config.allowed_syscalls)
            mount_readonly_fs(config.readonly_fs)
            
            if !config.network {
                disable_network()
            }
            
            let result = deserialize_and_execute(payload)
            ipc_send(result)
            exit(0)
        }
    }
    
    func wait(self c: *container<compute>) R {
        let result = ipc_recv(c.pid)
        syscall_waitpid(c.pid)
        cleanup_cgroups(c.cgroup_path)
        return result
    }
    
    func kill(self c: *container<compute>) {
        syscall_kill(c.pid, SIGKILL)
        cleanup_cgroups(c.cgroup_path)
    }
}

// Usage
let c = container func(code: string) int {
    return eval(code)
}(ContainerConfig{
    cpu_limit: 50,
    memory_limit: 128_MB,
    network: false,
    readonly_fs: true
}, user_code)

let result = c.wait()
```

**Config:**
```arc
struct ContainerConfig {
    cpu_limit: int           // Percentage (0-100)
    memory_limit: usize      // Bytes
    network: bool            // Allow network access
    readonly_fs: bool        // Mount filesystem as read-only
    allowed_syscalls: array<int>  // Seccomp whitelist
}
```

---

### cloud - Remote Execution & Deployment

```arc
struct cloud<compute> {
    provider: string
    instance_id: string
    region: string
    detached: bool
    
    init(self c: *cloud<compute>, closure: func(T) R, config: CloudConfig, args: T) {
        c.provider = config.provider
        c.region = config.region
        c.detached = config.detached
        
        // Compile secondary binary with closure + captured vars + runtime
        let payload = compile_cloud_binary(closure, args)
        
        // Authenticate with cloud provider
        let auth = authenticate_provider(config.provider, config.auth)
        
        // Provision instance
        c.instance_id = cloud_api.create_instance(
            provider: config.provider,
            region: config.region,
            hardware: config.hardware,
            vpc: config.vpc,
            auth: auth
        )
        
        // Upload and execute binary
        cloud_api.upload_binary(c.instance_id, payload)
        cloud_api.start_execution(c.instance_id)
        
        if config.detached {
            // Detach - local process can exit, remote continues
            cloud_api.detach(c.instance_id)
        }
    }
    
    func wait(self c: *cloud<compute>) R {
        if c.detached {
            raise("Cannot wait on detached cloud execution")
        }
        
        // Poll for result
        let result = cloud_api.get_result(c.instance_id)
        return result
    }
    
    func destroy(self c: *cloud<compute>) {
        cloud_api.terminate_instance(c.instance_id)
    }
    
    func logs(self c: *cloud<compute>) string {
        return cloud_api.get_logs(c.instance_id)
    }
}

// Usage - RPC style execution
let job = cloud func(data: bytes) Model {
    return ai.train(data)
}(CloudConfig{
    provider: "gcp",
    hardware: "nvidia-a100",
    region: "us-central1"
}, dataset)

let model = job.wait()
job.destroy()

// Usage - Self-deploying server (detached)
cloud func() {
    let server = http.listen(":80")
    server.start()  // Runs forever in the cloud
}(CloudConfig{
    provider: "aws",
    region: "us-east-1",
    hardware: "cpu-2core-4gb",
    detached: true  // Local program exits, server stays up
})

print("Server deployed, exiting local program...")
```

**Config:**
```arc
struct CloudConfig {
    provider: string     // "aws", "gcp", "digitalocean", "ssh"
    region: string       // "us-east-1", "us-central1", etc.
    hardware: string     // "cpu-basic", "gpu-t4", "nvidia-a100", "ram-128gb"
    detached: bool       // If true, local process disconnects after launch
    auth: AuthConfig     // Authentication credentials
    vpc: string          // Target private network (optional)
}

struct AuthConfig {
    type: string         // "api_key", "oauth", "ssh_key"
    credentials: map<string, string>
}
```

**Compile-time behavior:**

When compiler detects `cloud func()`:

1. **Generates secondary binary:**
   ```
   __cloud_payload_abc123.bin
     ├─ Closure bytecode
     ├─ Captured variables (serialized)
     ├─ Arc minimal runtime
     └─ Required stdlib modules
   ```

2. **Main program replacement:**
   ```arc
   // Original code
   cloud func(data) { train(data) }(config, dataset)
   
   // Becomes
   let payload = embed("__cloud_payload_abc123.bin")
   let instance = cloud_api.deploy(config, payload)
   return cloud_api.get_result(instance)
   ```

---

## Compute Templates, compiler behavior

**The `<compute>` marker tells the compiler:**

1. This struct can be used with `MyStruct func() {}(args)` syntax
2. Call `init(self: *MyStruct<compute>, closure, ...)` when instantiated
3. The closure is the `func() {}` block
4. Args after `}(...)` are passed to init after the closure

**Desugaring:**
```arc
// You write:
let t = custom_thread func(x: int) { work(x) }(42)

// Compiler transforms to:
let t: custom_thread<compute>
init(&t, func(x: int) { work(x) }, 42)
```

---

## Compute Templates, rules and requirements

**Required:**
- Struct must be marked with `<compute>`
- Must have an `init(self: *StructName<compute>, closure, ...)` method
- Self parameter must use `*StructName<compute>` type signature

**Optional:**
- Can have additional methods (join, wait, cancel, destroy, etc.)
- Can store state between init and method calls
- Can be generic over other type parameters

**Syntax pattern:**
```arc
let handle = YourStruct func(params) ReturnType {
    // closure body
}(init_args)
```

**Execution flow:**
1. Compiler captures closure: `func(params) ReturnType { body }`
2. Calls `YourStruct.init(&handle, closure, init_args...)`
3. Returns handle for later method calls

---

## Quick Comparison

| Context | Isolation | Use Case | Overhead | Detach |
|---------|-----------|----------|----------|--------|
| `thread` | Shared memory | CPU-bound, blocking | ~5µs | Yes |
| `process` | Separate memory | Fault tolerance | ~5ms | No |
| `container` | Sandboxed | Security, limits | ~50ms | No |
| `cloud` | Remote VM/Container | GPU, Deploy, Scale | ~2s | Yes |

---

## Benefits

**Extensible execution models:**
- Not limited to built-in contexts
- Build domain-specific compute contexts
- Thread pools, rate limiters, GPU schedulers, distributed queues

**Consistent syntax:**
```arc
let t1 = thread func() { work() }(config, args)
let t2 = custom_thread func() { work() }(args)
let t3 = process func() { work() }(config, args)
```

**Zero magic:**
- All execution contexts are user-definable structs
- No compiler magic beyond the `<compute>` marker
- Full control over initialization and lifecycle