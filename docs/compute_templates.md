# Compute Templates (Version 2.1)

Compute templates allow you to define custom execution contexts using the `compute` keyword. 

The core idea is simple: the compiler lifts inline functions to named entry points, and the context (defined by the `compute` type) decides **where** and **how** that function runs (in a thread, a subprocess, a container, or on a remote server).

---

## Syntax Definition

Instead of marking a struct with a generic attribute, use the dedicated `compute` keyword. Under the hood, this generates a struct, but enables special compiler lifecycle hooks.

### Definition
```arc
compute Name {
    // 1. State fields (just like a struct)
    field: Type
    
    // 2. Initialization
    // 'args...' captures the arguments passed to the inline function
    init(self n: *Name, config: ConfigType, args...) {
        // Setup logic
    }
    
    // 3. Methods
    func wait(self n: *Name) auto {
        // Retrieval logic
    }
}
```

### Usage
At the callsite, the syntax looks like instantiating a "function in a context".

```arc
let handle = Name func(param: Type) ReturnType {
    // Function body
}(config_arg, param_value)
```

---

## How It Works

When you write:
```arc
let job = process func(x: int32, y: int32) int32 {
    return x + y
}(ProcessConfig{}, 10, 20)
```

The compiler:

1.  **Lifts** the inline function to a named global function (e.g., `__process_entry_0`).
2.  **Injects** it into a dispatch table in `main` so the binary can route to it via flags.
3.  **Rewrites** the callsite to initialize the `process` struct.

**Generated Code Equivalent:**
```arc
// 1. Lifted Function
func __process_entry_0(x: int32, y: int32) int32 {
    return x + y
}

// 2. Main Dispatch (Hidden)
func main() {
    if os.args.has("--compute") {
        let id = os.args.get("--compute")
        match id {
            "__process_entry_0" => {
                let args = deserialize(stdin)
                let result = __process_entry_0(args.0, args.1) // args... unpacked
                serialize(stdout, result)
            }
        }
        return
    }
    actual_main()
}

// 3. Call Rewrite
// The compiler calls init with the ID and the arguments
let job = process.init(alloc(process), ProcessConfig{}, 10, 20)
```

---

## Compiler Intrinsics

These are available inside the `init` method of a `compute` type.

| Intrinsic | Type | Description |
|-----------|------|-------------|
| `@compute_id()` | `string` | The generated ID of the lifted function (e.g., `"__process_entry_0"`). |
| `@self_binary()` | `*byte` | A pointer to the current executable's bytes (useful for cloning/uploading). |

---

## Standard Contexts

### 1. Threads (`thread`)
Threads share memory, so we do not need to serialize arguments. We pass them directly to the runtime spawner.

```arc
compute thread {
    handle: ThreadHandle
    
    // args... captures the arguments intended for the function
    init(self t: *thread, config: ThreadConfig, args...) {
        // runtime.spawn_thread takes (func_id, config, variadic_args)
        t.handle = runtime.spawn_thread(@compute_id(), config, args...)
    }
    
    func join(self t: *thread) auto {
        return runtime.thread_join(t.handle)
    }
}
```

**Usage:**
```arc
let t = thread func(arr: *int) {
    sort(arr)
}(ThreadConfig{}, my_array) // my_array is passed by pointer, safe in threads

t.join()
```

### 2. Processes (`process`)
Processes have separate memory spaces. Arguments must be serialized and piped to the child process.

```arc
compute process {
    pid: int32
    stdin: Pipe
    stdout: Pipe
    
    init(self p: *process, config: ProcessConfig, args...) {
        let id = @compute_id()
        let pipes = pipe_create()
        
        p.stdin = pipes.write
        p.stdout = pipes.read
        
        // Spawn self with magic flag
        p.pid = os.spawn(
            os.executable_path(),
            ["--compute=${id}"],
            stdin: pipes.child_read,
            stdout: pipes.child_write,
            env: config.env
        )
        
        // Serialize arguments and send to child
        let payload = serialize(args...)
        p.stdin.write(payload)
        p.stdin.close()
    }
    
    func wait(self p: *process) auto {
        let output = p.stdout.read_all()
        os.waitpid(p.pid)
        return deserialize(output)
    }
}
```

**Usage:**
```arc
let p = process func(a: int, b: int) int {
    return a + b
}(ProcessConfig{}, 10, 20)

let res = p.wait() // 30
```

### 3. Cloud (`cloud`)
Uploads the binary to a remote server via SSH and executes it there.

```arc
compute cloud {
    conn: SSHConnection
    remote_pid: int32
    
    init(self c: *cloud, config: CloudConfig, args...) {
        let id = @compute_id()
        let binary = @self_binary()
        
        c.conn = ssh.connect(config.host, config.auth)
        
        // Upload self to remote
        c.conn.upload("/tmp/worker", binary)
        c.conn.chmod("/tmp/worker", 0o755)
        
        // Execute remote binary with flag
        let payload = serialize(args...)
        c.remote_pid = c.conn.exec("/tmp/worker --compute=${id}", stdin: payload)
    }
    
    func wait(self c: *cloud) auto {
        let output = c.conn.wait(c.remote_pid)
        return deserialize(output)
    }
}
```

**Usage:**
```arc
let job = cloud func(data: []byte) Result {
    return heavy_gpu_task(data)
}(CloudConfig{host: "192.168.1.50"}, large_dataset)

let result = job.wait()
```

---

## Custom Compute Contexts

You can define application-specific contexts, like a Worker Pool.

```arc
compute worker_pool {
    pool: *Pool
    task_id: int32
    
    init(self w: *worker_pool, pool: *Pool, args...) {
        let id = @compute_id()
        let payload = serialize(args...)
        
        w.pool = pool
        w.task_id = pool.submit(id, payload)
    }
    
    func wait(self w: *worker_pool) auto {
        let result_bytes = w.pool.get_result(w.task_id)
        return deserialize(result_bytes)
    }
}
```

**Usage:**
```arc
let my_pool = Pool{workers: 8}

let task = worker_pool func(x: int) int {
    return x * x
}(&my_pool, 5) // Pass pool config as first arg, then func args

let result = task.wait()
```

---

## Summary of Rules

1.  **Keyword:** Use `compute Name { ... }` to define the context.
2.  **Structure:** Must contain an `init` method.
3.  **Arguments:** `init` must accept `(self, [ConfigArgs], args...)`.
    *   `[ConfigArgs]` are passed in the second set of parentheses at the callsite.
    *   `args...` are the arguments passed to the inline function.
4.  **Serialization:** Unless memory is shared (like threads), `args...` must be serialized in `init` and deserialized by the generated entry point logic (handled by compiler infrastructure in `main`).
5.  **Return Type:** The return type of methods returning `auto` is inferred from the inline function's return type.