<h1 align="center">
  <a href="https://github.com/arc-language/arc-lang"><img src="./.github/arc-logo.jpeg" alt="Arc AI Language" height="150px"></a>
  <br>
  Arc AI Language
  <br>
</h1>
<h4 align="center">A systems programming language for AI inference, training, and direct hardware control with zero-overhead abstractions</h4>
<p align="center">
    <a href="https://github.com/arc-language/arc-lang"><img src="https://img.shields.io/badge/Arc-AI%20Systems%20Language-blue.svg?longCache=true" alt="Arc" /></a>
    <a href="https://github.com/arc-language/arc-lang"><img src="https://img.shields.io/badge/version-1.6-brightgreen" /></a>
    <a href="https://github.com/arc-language/arc-lang"><img src="https://img.shields.io/badge/C%2FC%2B%2B-interop-purple" /></a>
  <br>
    <a href="https://github.com/arc-language/arc-lang"><img src="https://img.shields.io/static/v1?label=Build&message=Documentation&color=brightgreen" /></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-5865F2.svg" alt="License: MIT" /></a>
</p>
<br>

### New Release

Arc v1.6 has been released! See the [release notes](https://github.com/arc-language/arc-lang/releases) to learn about new features including enhanced async/await support, mutating struct methods, and expanded intrinsics for AI workloads and systems programming.

If you aren't ready to upgrade yet, check the [tags](https://github.com/arc-language/arc-lang/tags) for previous stable releases.

We appreciate your feedback! Feel free to open GitHub issues or submit changes to stay updated in development and connect with the maintainers.

-----

### About Arc AI Language

Arc AI Language is a modern systems programming language purpose-built for artificial intelligence and machine learning workloads. Combining Python-like simplicity with C-level performance, Arc provides direct hardware access, GPU acceleration, and zero-overhead abstractions for AI inference, training pipelines, and performance-critical applications.

## Getting Started

Clone and build the Arc compiler:

```bash
git clone https://github.com/arc-language/arc-lang
cd arc-lang/compiler
go build main.go -o arc
```

Compile and run your first AI program:

```bash
./arc build main.arc
```

## Simple AI Inference Example

```arc
namespace main

import "ai"
import "io"

func main() {
    // Load model from disk
    let model = ai.load("models/model.gguf")
    
    // Configure generation parameters
    let config = ai.GenerateConfig{
        temperature: 0.7,
        max_tokens: 512,
        stream: true
    }
    
    // Stream tokens as they generate
    for token in model.generate("Hello AI!", config) {
        io.print(token)
    }
    
    io.println("\nDone!")
}
```

**[Example Applications](examples/README.md)** contain code samples demonstrating AI/ML use cases with Arc.

**[Language Reference](docs/REFERENCE.md)** provides a comprehensive guide to Arc's syntax and features.

**[Standard Library Documentation](docs/STDLIB.md)** covers the built-in packages and AI-specific APIs.

**[AI Library Documentation](docs/AI_LIB.md)** details the AI module for model loading, inference, and training.

Now go build something intelligent! Here are some ideas to spark your creativity:
* Deploy large language models with streaming inference and quantization support
* Build custom training pipelines with GPU acceleration and mixed precision
* Create real-time vision systems leveraging SIMD intrinsics and zero-copy buffers
* Implement reinforcement learning agents with direct hardware control
* Develop edge AI applications for embedded devices with minimal runtime overhead
* Write high-performance vector databases and embedding stores

## Building

See [BUILDING.md](https://github.com/arc-language/arc-lang/blob/main/BUILDING.md) for complete building instructions including compiler setup, dependencies, and platform-specific notes.

### Features

#### AI-First Design
* Native support for tensor operations and model loading (GGUF, ONNX, SafeTensors)
* Streaming inference with token-by-token generation
* GPU acceleration through CUDA, Metal, and Vulkan backends
* Quantization support (FP16, INT8, INT4) for efficient inference
* Zero-copy data loading for large datasets
* Automatic batching and memory pooling for training workloads

#### Modern Systems Language
* Direct hardware access through intrinsics and inline assembly
* Zero-overhead abstractions - pay only for what you use
* No hidden allocations or runtime overhead
* Manual memory management with optional reference counting for classes
* Compile-time evaluation and metaprogramming capabilities

#### Powerful Type System
* Value types (structs) vs reference types (classes) for optimal memory layout
* Fixed-width integers (int8, int16, int32, int64) for precise control
* Architecture-dependent types (usize, isize) for portable pointer arithmetic
* Generic collections with type inference
* Seamless C/C++ interoperability with extern blocks

#### Low-Level Control
* Stack allocation via `alloca` intrinsic for fast temporary buffers
* Direct syscall support without libc dependency
* Pointer arithmetic and raw memory manipulation
* Memory intrinsics: memset, memcpy, memmove, memcmp, memchr
* Bit-level operations and type punning with bit_cast

#### High-Level Ergonomics
* Python-like syntax for rapid AI prototyping
* Async/await for concurrent I/O and distributed training
* Defer statements for guaranteed cleanup (RAII-style)
* Iterator-based for-in loops over datasets and batches
* Methods on both structs and classes (inline or flat declaration)
* Mutating methods for in-place tensor modifications

#### Safe by Default, Unsafe When Needed
* Explicit pointer types distinguish nullable from non-null
* Bounds checking on collections in debug builds
* Reference counting for automatic class lifecycle management
* Manual memory control available when performance demands it

#### Built for Performance
* Compiled to native machine code with aggressive optimizations
* Cache-friendly data structures with predictable memory layout
* SIMD intrinsics for vectorized operations (essential for AI kernels)
* Profile-guided optimization support
* Minimal runtime with optional standard library
* GPU kernel fusion and automatic memory transfers

### Language Philosophy

Arc AI Language follows these core principles:

1. **Explicit is better than implicit** - No hidden allocations, conversions, or side effects
2. **Zero-cost abstractions** - High-level AI features compile to optimal machine code
3. **Control when you need it** - Drop down to CUDA kernels or assembly without friction
4. **Python ergonomics, C performance** - Clean syntax that doesn't sacrifice speed
5. **AI-native primitives** - Tensors, models, and GPU operations as first-class citizens
6. **Interoperable by design** - Play nice with existing ML frameworks and C/C++ codebases

### Contributing

Check out the [contributing guide](https://github.com/arc-language/arc-lang/wiki/Contributing) to join the team of dedicated contributors making this project possible.

### Community

* [GitHub Discussions](https://github.com/arc-language/arc-lang/discussions) - Propose features, ask questions, and share AI projects
* [Issues](https://github.com/arc-language/arc-lang/issues) - Report bugs and request features

### License

MIT License - see [LICENSE](LICENSE) for full text