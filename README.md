<h1 align="center">
  <img src="./.github/arc_logo.jpeg" alt="Arc Language" width="200px">
</h1>

<h4 align="center">AI Language<br>Native GPU/TPU AI development. Zero-overhead systems programming.</h4>

<p align="center">
    <img src="https://img.shields.io/badge/Version-2.0-blue" alt="Version">
    <img src="https://img.shields.io/badge/Targets-CPU%20%7C%20GPU%20%7C%20TPU%20%7C%20QPU-purple" alt="Targets">
    <img src="https://img.shields.io/badge/License-MIT-green" alt="License">
</p>

---

## What is Arc?

Arc lets you write AI models and deploy them to **any hardware—GPUs, TPUs, or quantum processors—from a single codebase.** No context switching between Python and C++.

Write your model once. Deploy everywhere. With full systems control when you need it.

---

## Write Once. Run Anywhere.

```arc
import "ai"

// Train on NVIDIA GPUs
async func train_model<gpu.cuda>(model: NeuralNet, data: Tensor) {
    for epoch in 0..100 {
        let loss = model.forward(data)
        model.backward(loss)
        model.step()
    }
}

// Deploy to Google TPUs
async func inference<tpu>(model: NeuralNet, input: Tensor) Tensor {
    return model.forward(input)
}

// Or run on quantum hardware
async func quantum_circuit<qpu.ibm>() Result {
    let qubits = qpu.alloc(2)
    qpu.h(qubits[0])           // Hadamard gate
    qpu.cx(qubits[0], qubits[1])  // Entanglement
    return qpu.measure(qubits)
}

func main() {
    let model = NeuralNet.create([784, 128, 10])
    
    // Train on GPU
    await train_model(model, training_data)
    
    // Deploy to TPU for inference
    let result = await inference(model, test_input)
}
```

**That's it.** Same language. Same model. Different hardware. Zero friction.

---

## Core Features

### Multi-Target Compilation
Write `async func<target>` and Arc compiles to native code for your hardware:
- **`<gpu.cuda>`** → NVIDIA GPUs
- **`<gpu.rocm>`** → AMD GPUs  
- **`<tpu>`** → Google Cloud TPUs
- **`<qpu.ibm>`** → IBM Quantum Processors

### ML-Native Types & Operations
```arc
import "ai"

let model = Sequential([
    Dense(784, 128, activation: relu),
    Dropout(0.2),
    Dense(128, 10, activation: softmax)
])

// Automatic differentiation
let loss = cross_entropy(model(x), y)
loss.backward()  // Gradients computed automatically
```

### GPU Kernels Made Simple
```arc
async func matrix_multiply<gpu>(A: Tensor, B: Tensor) Tensor {
    let i = gpu.thread_id()
    let result = Tensor.zeros(A.shape[0], B.shape[1])
    
    // Arc handles memory, scheduling, synchronization
    result[i] = dot(A[i], B.transpose())
    return result
}
```

### Systems Programming When Needed
Need to write a custom device driver or optimize memory allocation? Arc is a real systems language:

```arc
import "kernel/driver"

func init() driver.Status {
    let device = driver.create("CustomAccelerator")
    device.on_interrupt(func() {
        // Handle hardware interrupt
    })
    return driver.OK
}
```

### Distributed Training
```arc
// Run training across a cluster
async func distributed_train<cloud>(model: NeuralNet) {
    let rank = cloud.rank()
    let size = cloud.world_size()
    
    // Automatic data parallelism
    let shard = data.partition(rank, size)
    await train(model, shard)
    
    // Gradient synchronization handled automatically
    cloud.all_reduce(model.gradients)
}
```

---

## Supported Hardware

| Platform | Target | Hardware |
|:---|:---|:---|
| **NVIDIA GPUs** | `<gpu.cuda>` | Maxwell → Hopper (RTX, A100, H100) |
| **AMD GPUs** | `<gpu.rocm>` | RDNA / CDNA (RX 7000, MI300) |
| **Apple Silicon** | `<gpu.metal>` | M1 → M4 |
| **Intel GPUs** | `<gpu.oneapi>` | Arc / Data Center Max |
| **Google TPUs** | `<tpu>` | Cloud TPU v2-v5p |
| **IBM Quantum** | `<qpu.ibm>` | Eagle / Heron (Superconducting) |
| **IonQ Quantum** | `<qpu.ionq>` | Harmony / Aria (Trapped Ion) |
| **AWS Trainium** | `<aws.trainium>` | Trn1 / Inf2 |

---

## Getting Started

### Compile and Run (BETA)
It's in beta so you can only run the test_runner and some .arc files.

```bash
git clone https://github.com/arc-language/arc-lang
cd arc-lang/cmd
./build build
./test_runner

./arc build main.arc -o main
./main
```

### Write Your Program

Create `main.arc`:

```arc
extern c {
  func printf(*byte, ...)
}

func main() {
  printf("main\n")
}
```

**No external dependencies.** No CUDA toolkit. No Python environment. Just Arc.

---

## Why Arc?

| Feature | PyTorch/JAX | CUDA/C++ | Arc |
|:---|:---:|:---:|:---:|
| **GPU Training** | ✅ | ✅ | ✅ |
| **TPU Support** | ✅ (JAX only) | ❌ | ✅ |
| **Zero Python Overhead** | ❌ | ✅ | ✅ |
| **Write Custom Kernels** | ⚠️ (Need C++) | ✅ | ✅ |
| **Single Language** | ❌ | ❌ | ✅ |
| **Systems Programming** | ❌ | ✅ | ✅ |
| **Quantum Computing** | ❌ | ❌ | ✅ |

---

## Learn More

- **[Quickstart Guide](docs/quickstart.md)** - Build your first model in 5 minutes
- **[ML Primitives](docs/ml_primitives.md)** - Neural networks, autodiff, optimizers
- **[GPU Programming](docs/gpu_guide.md)** - Write custom CUDA-like kernels
- **[TPU Deployment](docs/tpu_guide.md)** - Deploy to Google Cloud TPUs
- **[Systems Programming](docs/systems.md)** - Kernel drivers, memory management
- **[Language Reference](docs/reference.md)** - Complete syntax and semantics

---

## License

Licensed under either of

*   Apache License, Version 2.0 ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
*   MIT license ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.

---

**Arc: One language. Any chip. Zero compromise.**
