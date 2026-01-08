# Compute Templates (Version 2.0)

Compute templates allow you to define custom execution contexts. Mark a struct with `<compute>` to enable the `Type func() {}(args)` syntax pattern.

The core idea is simple: the compiler lifts inline functions to named entry points, and the same binary runs everywhere with different flags.

---

## How It Works

When you write:
```arc
let job = process func(x: int32, y: int32) int32 {
    return x + y
}(ProcessConfig{}, 10, 20)
```

The compiler:

1. Lifts the inline function to a named function
2. Adds it to a dispatch table in main
3. At the callsite, spawns the same binary with `--compute=<id>`
```arc
// Compiler generates this entry point
func __process_0(x: int32, y: int32) int32 {
    return x + y
}

// Compiler generates dispatch in main
func main() {
    if os.args.has("--compute") {
        let id = os.args.get("--compute")
        match id {
            "__process_0" => {
                let args = deserialize(stdin)
                let result = __process_0(args.0, args.1)
                serialize(stdout, result)
            }
        }
        return
    }
    
    // Normal main continues...
    actual_main()
}
```

No magic on the child side. Just: run binary, see flag, call function, exit.

---

## Compiler Intrinsics

Two intrinsics are available inside `<compute>` init:

| Intrinsic | Returns | Description |
|-----------|---------|-------------|
| `@compute_id()` | `string` | The generated entry point ID for this callsite |
| `@self_binary()` | `*byte` | Bytes of the current executable |

---

## Built-in Contexts

### thread - OS Threads

Threads share memory, so no serialization needed.
```arc
struct thread<compute> {
    handle: ThreadHandle
    
    init(self t: *thread<compute>, config: ThreadConfig, args...) {
        t.handle = runtime.spawn_thread(@compute_id(), config, args...)
    }
    
    func join(self t: *thread<compute>) auto {
        return runtime.thread_join(t.handle)
    }
    
    func detach(self t: *thread<compute>) {
        runtime.thread_detach(t.handle)
    }
}

struct ThreadConfig {
    stack_size: usize
    priority: int32
    name: string
}
```

**Usage:**
```arc
let t = thread func(data: *int32, len: usize) int64 {
    let sum: int64 = 0
    for let i: usize = 0; i < len; i++ {
        sum += cast<int64>(data[i])
    }
    return sum
}(ThreadConfig{stack_size: 8_MB}, array_ptr, array_len)

let result = t.join()
```

---

### process - OS Processes

Separate memory space. Same binary, different entry point.
```arc
struct process<compute> {
    pid: int32
    stdin: Pipe
    stdout: Pipe
    
    init(self p: *process<compute>, config: ProcessConfig, args...) {
        let id = @compute_id()
        
        let pipes = pipe_create()
        p.stdin = pipes.write
        p.stdout = pipes.read
        
        p.pid = os.spawn(
            os.executable_path(),
            ["--compute=${id}"],
            stdin: pipes.child_read,
            stdout: pipes.child_write,
            env: config.env,
            cwd: config.working_dir
        )
        
        // Send args to child
        let args_bytes = serialize(args...)
        p.stdin.write(args_bytes)
        p.stdin.close()
    }
    
    func wait(self p: *process<compute>) auto {
        let output = p.stdout.read_all()
        os.waitpid(p.pid)
        return deserialize(output)
    }
    
    func kill(self p: *process<compute>) {
        os.kill(p.pid, SIGKILL)
    }
}

struct ProcessConfig {
    env: map<string, string>
    working_dir: string
}
```

**Usage:**
```arc
let p = process func(code: string) int32 {
    return eval(code)
}(ProcessConfig{}, user_code)

let result = p.wait()
```

---

### container - Sandboxed Processes

Same as process, plus namespaces and cgroups for isolation.
```arc
struct container<compute> {
    pid: int32
    stdout: Pipe
    cgroup_path: string
    
    init(self c: *container<compute>, config: ContainerConfig, args...) {
        let id = @compute_id()
        
        // Setup resource limits
        c.cgroup_path = cgroup.create(config.cpu_limit, config.memory_limit)
        
        let pipes = pipe_create()
        c.stdout = pipes.read
        
        c.pid = os.spawn(
            os.executable_path(),
            ["--compute=${id}"],
            stdin: pipes.child_read,
            stdout: pipes.child_write,
            namespaces: CLONE_NEWNS | CLONE_NEWPID | CLONE_NEWNET,
            cgroup: c.cgroup_path
        )
        
        let args_bytes = serialize(args...)
        pipes.write.write(args_bytes)
        pipes.write.close()
    }
    
    func wait(self c: *container<compute>) auto {
        let output = c.stdout.read_all()
        os.waitpid(c.pid)
        cgroup.cleanup(c.cgroup_path)
        return deserialize(output)
    }
    
    func kill(self c: *container<compute>) {
        os.kill(c.pid, SIGKILL)
        cgroup.cleanup(c.cgroup_path)
    }
}

struct ContainerConfig {
    cpu_limit: int32      // Percentage 0-100
    memory_limit: usize   // Bytes
    network: bool
    readonly_fs: bool
}
```

**Usage:**
```arc
let c = container func(input: string) string {
    return dangerous_parse(input)
}(ContainerConfig{cpu_limit: 50, memory_limit: 128_MB, network: false}, data)

let result = c.wait()
c.kill()
```

---

### cloud - Remote Execution

Upload binary to remote machine, execute there.
```arc
struct cloud<compute> {
    conn: SSHConnection
    remote_pid: int32
    
    init(self c: *cloud<compute>, config: CloudConfig, args...) {
        let id = @compute_id()
        let binary = @self_binary()
        
        c.conn = ssh.connect(config.host, config.auth)
        c.conn.upload("./worker", binary)
        c.conn.chmod("./worker", 0o755)
        
        let args_bytes = serialize(args...)
        c.remote_pid = c.conn.exec("./worker --compute=${id}", stdin: args_bytes)
    }
    
    func wait(self c: *cloud<compute>) auto {
        let output = c.conn.wait(c.remote_pid)
        return deserialize(output)
    }
    
    func destroy(self c: *cloud<compute>) {
        c.conn.exec("kill ${c.remote_pid}")
        c.conn.exec("rm ./worker")
        c.conn.close()
    }
    
    func logs(self c: *cloud<compute>) string {
        return c.conn.get_stderr(c.remote_pid)
    }
}

struct CloudConfig {
    host: string
    auth: AuthConfig
}

struct AuthConfig {
    type: string   // "ssh_key", "password"
    key: string
}
```

**Usage:**
```arc
let job = cloud func(data: *byte, len: usize) float64 {
    return heavy_compute(data, len)
}(CloudConfig{host: "gpu-server.internal", auth: ssh_key}, dataset, dataset_len)

let result = job.wait()
job.destroy()
```

---

## Custom Compute Templates

You can define your own execution contexts.
```arc
struct worker_pool<compute> {
    pool: *Pool
    task_id: int32
    
    init(self w: *worker_pool<compute>, pool: *Pool, args...) {
        let id = @compute_id()
        let args_bytes = serialize(args...)
        
        w.pool = pool
        w.task_id = pool.submit(id, args_bytes)
    }
    
    func wait(self w: *worker_pool<compute>) auto {
        let output = w.pool.get_result(w.task_id)
        return deserialize(output)
    }
}

// Usage
let pool = Pool{workers: 8}

let task1 = worker_pool func(x: int32) int32 {
    return expensive(x)
}(&pool, 100)

let task2 = worker_pool func(x: int32) int32 {
    return expensive(x)
}(&pool, 200)

let r1 = task1.wait()
let r2 = task2.wait()
```

---

## Rules

**Required:**
- Struct must have `<compute>` marker
- Must have `init(self: *StructName<compute>, ...)`

**Syntax patterns:**
```arc
let handle = YourStruct func(params) ReturnType {
    // body
}(init_args)
```

```arc
// or from module with namespace 
let handle  = docker.Container func(params) docker.ReturnType {
    // body
}(init_args)
```

**Compiler behavior:**
1. Lifts `func(params) { body }` to a named function
2. Generates dispatch table entry in main
3. Calls `YourStruct.init()` with `@compute_id()` available
4. Handle returned for later method calls

---

## Comparison

| Context | Isolation | Upload | Use Case |
|---------|-----------|--------|----------|
| `thread` | Shared memory | No | CPU parallelism |
| `process` | Separate memory | No | Fault isolation |
| `container` | Namespaced | No | Security, resource limits |
| `cloud` | Remote machine | Yes | Scale, GPUs, deployment |

All use the same pattern: same binary, `--compute=id` flag, serialize args over stdin, deserialize result from stdout. Only `thread` skips serialization since it shares memory.