# Arc AMD GPU Execution Model

**Key Concept:** Functions marked `async func<gpu.rocm>` compile **entirely** to HIP/GCN assembly. Everything - control flow, loops, operations - becomes AMD GPU instructions.

---

## How `<gpu.rocm>` Functions Work

### The Entire Function Becomes HIP/GCN

```arc
async func double_array<gpu.rocm>(arr: *float32, n: usize) *float32 {
    // EVERYTHING in here → HIP/GCN instructions
    let idx = gpu.thread_id()
    
    if idx < n {  // ← GCN s_cmp + s_cbranch
        return arr[idx] * 2.0  // ← GCN flat_load + v_mul_f32 + flat_store
    }
}
```

### Control Flow Mapping

| Your Code | Compiles To |
|-----------|-------------|
| `if ... else` | GCN `s_cmp` + `s_cbranch` or wavefront exec masks |
| `for` loop | GCN loop with `s_cmp` + `s_cbranch` |
| `while` loop | GCN loop with conditional branches |
| Arithmetic | GCN `v_add_f32`, `v_mul_f32`, etc. |
| Array access | GCN `flat_load_dword`, `flat_store_dword` |

**Not Supported:**
- System calls, recursive functions, dynamic memory allocation within kernels

---

## Memory Model

### Unified Memory (Recommended)

```arc
async func process<gpu.rocm>(arr: *float32, n: usize) *float32 {
    let result = gpu.unified_malloc<float32>(n)  // CPU+GPU accessible
    let idx = gpu.thread_id()
    if idx < n {
        result[idx] = arr[idx] * 2.0
    }
    return result
}
```

- Calls `hipMallocManaged()` - accessible from both CPU and GPU
- ROCm runtime handles transfers automatically via HSA
- Simpler but may have performance overhead

### Explicit Memory (Performance Critical)

```arc
async func process_explicit<gpu.rocm>(gpu_arr: *float32, gpu_result: *float32, n: usize) {
    let idx = gpu.thread_id()
    if idx < n {
        gpu_result[idx] = gpu_arr[idx] * 2.0
    }
}

func use_explicit() {
    let cpu_data = alloca<float32>(1024)
    let gpu_data: *float32 = null
    let gpu_result: *float32 = null
    
    rocm.hipMalloc(&gpu_data, 1024 * sizeof<float32>)
    rocm.hipMalloc(&gpu_result, 1024 * sizeof<float32>)
    
    rocm.hipMemcpyHtoD(gpu_data, cpu_data, 1024 * sizeof<float32>)
    await process_explicit(gpu_data, gpu_result, 1024)
    rocm.hipMemcpyDtoH(cpu_data, gpu_result, 1024 * sizeof<float32>)
    
    rocm.hipFree(gpu_data)
    rocm.hipFree(gpu_result)
}
```

---

## Arc's Native Linking

Arc links directly to `libamdhip64.so` - no hipcc required:

```
Arc Source (async func<gpu.rocm>)
    ↓
Arc Parser → AST
    ↓
HIP Backend → GCN Assembly (embedded in executable)
    ↓
Arc Linker → Links to libamdhip64.so
    ↓
Runtime → ROCm Driver → AMD GPU Execution
```

---

## ROCm Software Stack

```
Arc generates HIP/GCN → HIP Runtime → HSA Runtime → AMD GPU Driver → GPU
                       ↓
                       libamdhip64.so
                                      ↓
                                      libhsa-runtime64.so
```

**HIP Runtime API** - Device management, kernel execution, memory transfers  
**HSA Runtime** - Low-level hardware communication  
**AMD GPU Driver** - Direct GPU communication

### Arc's Approach

```arc
extern hip {
    func hipInit(flags: uint32) int32
    func hipGetDeviceCount(count: *int32) int32
    func hipMallocManaged(ptr: **void, size: usize, flags: uint32) int32
    func hipModuleLoad(module: **void, fname: *byte) int32
    func hipModuleLaunchKernel(
        func: *void,
        gridDimX: uint32, gridDimY: uint32, gridDimZ: uint32,
        blockDimX: uint32, blockDimY: uint32, blockDimZ: uint32,
        sharedMemBytes: uint32, stream: *void,
        kernelParams: **void, extra: **void
    ) int32
    func hipDeviceSynchronize() int32
    func hipFree(ptr: *void) int32
}
```

---

## Two-Phase Compilation

### Build Time (AOT)

```bash
arc build --target=gpu.rocm my_program.arc
```

1. Compile `async func<gpu.rocm>` → GCN ISA
2. Embed compiled GPU code in executable
3. Link to `libamdhip64.so`

### Runtime Execution

```arc
let result = await double_array(data, 1024)  // First call
```

1. Load embedded GPU code from executable
2. `hipModuleLoad()` → module handle
3. `hipModuleLaunchKernel()` executes on GPU
4. `await` synchronizes with `hipDeviceSynchronize()`
5. Return result pointer

**Subsequent calls reuse loaded module**

---

## Complete Example

```arc
async func matmul<gpu.rocm>(A: *float32, B: *float32, C: *float32, N: usize) {
    let total = N * N
    let idx = gpu.thread_id()
    
    if idx < total {
        let row = idx / N
        let col = idx % N
        let mut sum: float32 = 0.0
        
        for k in 0..N {
            sum += A[row * N + k] * B[k * N + col]
        }
        
        C[row * N + col] = sum
    }
}

func main() {
    const N: usize = 1024
    
    let A = gpu.unified_malloc<float32>(N * N)
    let B = gpu.unified_malloc<float32>(N * N)
    let C = gpu.unified_malloc<float32>(N * N)
    
    // Initialize on CPU
    for i in 0..(N * N) {
        A[i] = 1.0
        B[i] = 2.0
    }
    
    // Execute on AMD GPU
    await matmul(A, B, C, N)
    
    io.printf("C[0] = %f\n", C[0])
    
    gpu.unified_free(A)
    gpu.unified_free(B)
    gpu.unified_free(C)
}
```

---

## AMD GPU Architectures

**GCN (Graphics Core Next)**
- Wavefront size: **64 threads**
- Used in: Radeon RX series, older Instinct

**RDNA (Gaming/Consumer)**
- Wavefront size: **32 threads**
- Used in: Radeon RX 5000/6000/7000, PlayStation 5

**CDNA (Compute/Datacenter)**
- Wavefront size: **64 threads**
- Matrix cores (MFMA)
- Used in: Instinct MI100, MI200, MI300

Arc runtime auto-detects architecture and adjusts wavefront sizes.

---

## SPMD Parallelism

```arc
async func parallel_work<gpu.rocm>(data: *float32, n: usize) *float32 {
    let tid = gpu.thread_id()  // Which thread am I? (0 to n-1)
}
```

**AMD Wavefront Model:**
- **Wavefront** = Group of threads executing together
- 32 threads (RDNA) or 64 threads (GCN/CDNA)
- All threads in wavefront execute in lockstep
- Divergent branches handled via execution masks

---

## AMD-Specific Features

### LDS (Local Data Share)

Fast on-chip shared memory:

```arc
async func blocked_matmul<gpu.rocm>(A: *float32, B: *float32, C: *float32, N: usize) {
    let shared_A: [float32; 256] @lds
    let shared_B: [float32; 256] @lds
    
    let local_id = gpu.local_thread_id()
    
    // Load into LDS
    shared_A[local_id] = A[/* ... */]
    shared_B[local_id] = B[/* ... */]
    
    gpu.barrier()  // Synchronize wavefront
    
    // Compute using LDS data (much faster)
    let mut sum: float32 = 0.0
    for i in 0..256 {
        sum += shared_A[i] * shared_B[i]
    }
}
```

### Wavefront Intrinsics

```arc
async func wavefront_reduce<gpu.rocm>(data: *float32, n: usize) float32 {
    let tid = gpu.thread_id()
    let value = data[tid]
    
    // AMD wavefront reduction (across 32 or 64 threads)
    let sum = gpu.wavefront_reduce_add(value)
    
    if gpu.is_first_lane() {
        return sum
    }
}
```

---

## Multi-Device Execution

```arc
func multi_amd_gpu() {
    let num_gpus = gpu.device_count()
    
    let r0 = await(0) process_chunk(data0, size)
    let r1 = await(1) process_chunk(data1, size)
    let r2 = await(2) process_chunk(data2, size)
    
    let combined = average(r0, r1, r2)
}
```

---

## Runtime Implementation

```arc
namespace gpu

let initialized: bool = false
let device_count: int32 = 0

func init_rocm() {
    if initialized { return }
    
    hip.hipInit(0)
    hip.hipGetDeviceCount(&device_count)
    
    initialized = true
}

func unified_malloc<T>(count: usize) *T {
    init_rocm()
    let ptr: *void = null
    hip.hipMallocManaged(&ptr, count * sizeof<T>, 1)
    return cast<*T>(ptr)
}

func thread_id() usize {
    // Compiler intrinsic - replaced with HIP thread indexing
}
```

---

## Build and Run

```bash
# Build for AMD GPUs
arc build --target=gpu.rocm my_program.arc

# Run
./my_program
```

**Requirements:**
- AMD GPU (GCN 3.0+, RDNA, CDNA)
- ROCm driver and runtime libraries

**NOT Required:**
- HIP compiler toolchain (`hipcc`)
- ROCm development headers
- Any C/C++ dependencies

---

## Summary

1. `async func<gpu.rocm>` compiles entirely to GCN assembly
2. Unified memory is simplest - CPU and GPU share pointers
3. AOT compilation - Arc→GCN embedded in binary
4. `await` triggers execution and synchronization
5. HIP Runtime manages devices and memory
6. `gpu.thread_id()` enables SIMT parallelism
7. AMD features: MFMA matrix cores, LDS memory, wavefront intrinsics
8. Open-source ROCm stack