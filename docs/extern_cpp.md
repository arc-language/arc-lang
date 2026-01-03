# C++ Interoperability (`extern cpp`)

Arc structs are **already binary-compatible** with C++ POD (Plain Old Data) types. Since Arc compiles to `.o` files with proper ABI-compliant struct layouts, you can pass Arc structs directly to C++ functions without conversion.

## What Works Now ✅

### Struct Layout Compatibility

Arc's struct layout matches C++ ABI (both Itanium/GCC and MSVC):
- Fields aligned to natural boundaries
- Struct alignment = max field alignment
- Proper tail padding
- Field order preserved
- `@packed` attribute supported

**Example - DirectX Buffer Descriptor:**

```arc
// Arc struct - binary identical to C++
struct D3D11_BUFFER_DESC {
    byte_width: uint32           // Offset 0
    usage: uint32                // Offset 4
    bind_flags: uint32           // Offset 8
    cpu_access_flags: uint32     // Offset 12
    misc_flags: uint32           // Offset 16
    structure_byte_stride: uint32 // Offset 20
}
// Size: 24 bytes, aligned to 4

// C++ equivalent (DirectX SDK)
// struct D3D11_BUFFER_DESC {
//     UINT ByteWidth;
//     UINT Usage;
//     UINT BindFlags;
//     UINT CPUAccessFlags;
//     UINT MiscFlags;
//     UINT StructureByteStride;
// };
```

### Complete DirectX Example

```arc
import "std/io"

// Opaque COM interface pointer
struct ID3D11Device {}
struct ID3D11DeviceContext {}
struct ID3D11Buffer {}
struct IDXGIAdapter {}

// Result type
type HRESULT = int32

// Descriptor structs (POD - binary compatible!)
struct D3D11_BUFFER_DESC {
    byte_width: uint32
    usage: uint32
    bind_flags: uint32
    cpu_access_flags: uint32
    misc_flags: uint32
    structure_byte_stride: uint32
}

struct D3D11_SUBRESOURCE_DATA {
    sys_mem: *void
    sys_mem_pitch: uint32
    sys_mem_slice_pitch: uint32
}

// Constants
const D3D11_USAGE_DEFAULT: uint32 = 0
const D3D11_BIND_VERTEX_BUFFER: uint32 = 1

extern cpp "d3d11.lib" {
    // Factory function - exported by DLL
    func D3D11CreateDevice(
        adapter: *IDXGIAdapter,
        driver_type: uint32,
        software: *void,
        flags: uint32,
        feature_levels: *uint32,
        num_levels: uint32,
        sdk_version: uint32,
        device: **ID3D11Device,
        feature_level: *uint32,
        context: **ID3D11DeviceContext
    ) HRESULT
    
    // COM Interface - methods are vtable calls
    class ID3D11Device {
        // IUnknown methods (vtable offsets 0-2)
        func QueryInterface(self this: *ID3D11Device, riid: *void, out: **void) HRESULT
        func AddRef(self this: *ID3D11Device) uint32
        func Release(self this: *ID3D11Device) uint32
        
        // ID3D11Device methods (vtable offset 3+)
        func CreateBuffer(
            self this: *ID3D11Device,
            desc: *D3D11_BUFFER_DESC,      // ← Arc struct passed directly!
            data: *D3D11_SUBRESOURCE_DATA,
            buffer: **ID3D11Buffer
        ) HRESULT
        
        func CreateTexture2D(
            self this: *ID3D11Device,
            desc: *void,
            data: *void,
            texture: **void
        ) HRESULT
    }
    
    class ID3D11Buffer {
        func QueryInterface(self this: *ID3D11Buffer, riid: *void, out: **void) HRESULT
        func AddRef(self this: *ID3D11Buffer) uint32
        func Release(self this: *ID3D11Buffer) uint32
    }
}

func main() {
    let device: *ID3D11Device = null
    let context: *ID3D11DeviceContext = null
    
    // Create D3D11 Device
    let hr = d3d11.D3D11CreateDevice(
        null,  // Default adapter
        1,     // D3D_DRIVER_TYPE_HARDWARE
        null,
        0,     // Flags
        null,  // Feature levels
        0,
        7,     // D3D11_SDK_VERSION
        &device,
        null,
        &context
    )
    
    if hr != 0 {
        io.printf("Failed to create D3D11 device: %d\n", hr)
        return
    }
    
    // Create Vertex Buffer using Arc struct
    let vertices: array<float32, 9> = {
        0.0, 0.5, 0.0,
        0.5, -0.5, 0.0,
        -0.5, -0.5, 0.0
    }
    
    let desc = D3D11_BUFFER_DESC{
        byte_width: 36,  // 9 floats * 4 bytes
        usage: D3D11_USAGE_DEFAULT,
        bind_flags: D3D11_BIND_VERTEX_BUFFER,
        cpu_access_flags: 0,
        misc_flags: 0,
        structure_byte_stride: 0
    }
    
    let init_data = D3D11_SUBRESOURCE_DATA{
        sys_mem: cast<*void>(&vertices),
        sys_mem_pitch: 0,
        sys_mem_slice_pitch: 0
    }
    
    let buffer: *ID3D11Buffer = null
    
    // Pass Arc structs directly to C++ method! ✅
    hr = device.CreateBuffer(&desc, &init_data, &buffer)
    
    if hr != 0 {
        io.printf("Failed to create buffer: %d\n", hr)
    } else {
        io.printf("Buffer created successfully!\n")
    }
    
    // Cleanup
    if buffer != null {
        buffer.Release()
    }
    device.Release()
    context.Release()
}
```

## Implementation TODO 📋

To make `extern cpp` fully functional for interfacing with C++ compiled libraries, the following features need implementation:

### High Priority (Essential)

#### 1. **Name Mangling** 🔴
- [ ] **MSVC Name Mangling** (Windows)
  - Implement mangler for Windows C++ ABI
  - Handle namespaces: `namespace::Class::method` → `?method@Class@namespace@@QAEXH@Z`
  - Support function overloading signatures
  - Handle const/volatile qualifiers
  
- [ ] **Itanium Name Mangling** (Linux/Unix/macOS)
  - Implement Itanium C++ ABI mangler
  - Handle namespaces: `namespace::Class::method` → `_ZN9namespace5Class6methodEv`
  - Support complex type encoding
  - Handle template instantiations (basic)

**Example:**
```arc
extern cpp "math.lib" {
    namespace DirectX {
        func XMVectorAdd(v1: XMVECTOR, v2: XMVECTOR) XMVECTOR
    }
}
// Compiler mangles to: ?XMVectorAdd@DirectX@@YA?ATVVECTOR@@V1@0@Z (MSVC)
// Or: _ZN8DirectX11XMVectorAddENS_7XMVECTORES0_ (Itanium)
```

#### 2. **Calling Conventions** 🔴
- [ ] **`__thiscall`** (MSVC member functions)
  - First parameter (`this` pointer) passed in RCX (x64) or ECX (x86)
  - Remaining params follow normal conventions
  
- [ ] **`__stdcall`** (Win32 API)
  - Callee cleans stack (x86 only, x64 ignores)
  
- [ ] **`__vectorcall`** (SIMD optimization)
  - XMM registers for float vector types

**Already implemented:** System V AMD64 ABI (Linux) ✅

**Example:**
```arc
extern cpp "user32.lib" {
    // __stdcall on x86, default on x64
    func MessageBoxA(hwnd: *void, text: *byte, caption: *byte, type: uint32) int32
}
```

#### 3. **Vtable Call Generation** 🔴
- [ ] Generate vtable offset calculations for virtual methods
- [ ] Handle vtable pointer at offset 0 of class instances
- [ ] Support virtual inheritance (complex, lower priority)

**Current workaround:** Manual vtable offset tracking
```arc
class ID3D11Device {
    // Compiler knows these are vtable[0], vtable[1], vtable[2]...
    func QueryInterface(self this: *ID3D11Device, ...) HRESULT  // vtable[0]
    func AddRef(self this: *ID3D11Device) uint32                 // vtable[1]
    func Release(self this: *ID3D11Device) uint32                // vtable[2]
    func CreateBuffer(self this: *ID3D11Device, ...) HRESULT     // vtable[3]
}
```

**Generated code should be:**
```asm
; this->CreateBuffer(args...)
mov rax, [rcx]        ; Load vtable pointer
call [rax + 24]       ; Call vtable[3] (offset 3 * 8 bytes)
```

#### 4. **Constructor/Destructor Support** 🟡
- [ ] Map C++ constructors to Arc factory functions
- [ ] Map C++ destructors to Arc cleanup
- [ ] Handle RAII patterns

**Example:**
```arc
extern cpp "physics.lib" {
    class PhysicsWorld {
        // Constructor - special mangled name
        func new() *PhysicsWorld
        
        // Destructor - special mangled name
        func delete(self this: *PhysicsWorld) void
        
        func step(self this: *PhysicsWorld, dt: float32) void
    }
}

// Usage with Arc's defer
func simulate() {
    let world = PhysicsWorld.new()
    defer world.delete()
    
    world.step(0.016)
}
```

### Medium Priority (Common Use Cases)

#### 5. **Function Overloading** 🟡
- [ ] Allow multiple functions with same name, different signatures
- [ ] Generate unique mangled names based on parameter types

**Example:**
```arc
extern cpp "graphics.lib" {
    namespace Renderer {
        func draw(x: int32, y: int32) void              // draw_2i
        func draw(x: float32, y: float32, z: float32) void  // draw_3f
        func draw(pos: Vector3) void                    // draw_Vector3
    }
}
```

#### 6. **Operator Overloading Mapping** 🟡
- [ ] Map C++ operators to Arc method names
- [ ] Support arithmetic operators (+, -, *, /)
- [ ] Support comparison operators (==, !=, <, >)
- [ ] Support array subscript operator[]

**Example:**
```arc
extern cpp "math.lib" {
    struct Vector3 {
        x: float32
        y: float32
        z: float32
        
        // C++ operator+ maps to Arc method
        func operator_add "operator+" (a: Vector3, b: Vector3) Vector3
        func operator_mul "operator*" (v: Vector3, scalar: float32) Vector3
    }
}

// Usage
let v1 = Vector3{x: 1.0, y: 2.0, z: 3.0}
let v2 = Vector3{x: 4.0, y: 5.0, z: 6.0}
let v3 = v1.operator_add(v2)  // Calls C++ operator+
```

#### 7. **Const Method Support** 🟡
- [ ] Distinguish `const` vs non-`const` methods
- [ ] Generate different mangled names
- [ ] Enforce immutability at compile time

**Example:**
```arc
extern cpp "buffer.lib" {
    class Buffer {
        // const method - read-only
        func size(self this: *const Buffer) usize
        func data(self this: *const Buffer) *const byte
        
        // non-const method - can modify
        mutating resize(self this: *Buffer, n: usize) void
        mutating clear(self this: *Buffer) void
    }
}

const buf = get_buffer()
let sz = buf.size()      // OK - const method
// buf.resize(100)       // ERROR - can't call mutating method on const
```

#### 8. **Reference Parameters** 🟡
- [ ] Map C++ `&` (lvalue reference) to Arc reference syntax
- [ ] Map C++ `const&` to read-only references
- [ ] Handle reference return types

**Example:**
```arc
extern cpp "math.lib" {
    // C++: void modify(Vector3& v)
    func modify(v: &Vector3) void
    
    // C++: void read(const Vector3& v)  
    func read(v: &const Vector3) void
    
    // C++: const Vector3& get_velocity()
    func get_velocity() &const Vector3
}
```

### Low Priority (Advanced Features)

#### 9. **Nested Namespace Support** 🟢
- [ ] Parse `namespace A::B::C` syntax
- [ ] Generate proper mangled names with namespace hierarchy
- [ ] Support using-declarations

**Example:**
```arc
extern cpp "DirectXMath.lib" {
    namespace DirectX.PackedVector {
        struct XMCOLOR {
            r: uint8
            g: uint8
            b: uint8
            a: uint8
        }
        
        func XMColorRGBToSRGB(rgb: XMVECTOR) XMCOLOR
    }
}

// Usage
let color: DirectX.PackedVector.XMCOLOR = ...
```

#### 10. **STL Type Wrappers** 🟢
- [ ] Opaque wrappers for `std::string`
- [ ] Opaque wrappers for `std::vector<T>`
- [ ] Opaque wrappers for `std::map<K,V>`
- [ ] Optional: Auto-conversion between Arc and C++ types

**Example:**
```arc
extern cpp "engine.lib" {
    // Opaque wrapper
    struct CppString {
        func new() *CppString
        func from_cstr(s: *byte) *CppString
        func c_str(self this: *CppString) *byte
        func delete(self this: *CppString) void
    }
    
    func GetName() *CppString
    func SetName(name: *CppString) void
}

// Usage
let name_cpp = GetName()
let name_arc = string.from_cstr(name_cpp.c_str())
defer name_cpp.delete()
```

#### 11. **Template Instantiation Info** 🟢
- [ ] Parse common template instantiations
- [ ] Generate mangled names for `std::vector<int>`, etc.
- [ ] Limit to explicit instantiations (no template metaprogramming)

**Example:**
```arc
extern cpp "containers.lib" {
    class VectorInt "std::vector<int>" {
        func new() *VectorInt
        func push_back(self this: *VectorInt, val: int32) void
        func size(self this: *VectorInt) usize
        func delete(self this: *VectorInt) void
    }
}
```

## What's NOT Supported (By Design)

Arc will **never** be a full C++ compiler. The following are intentionally out of scope:

- ❌ C++ template metaprogramming
- ❌ Parsing C++ headers automatically
- ❌ C++ SFINAE / concepts
- ❌ C++ exception handling (use `try`/`except` for Arc exceptions)
- ❌ C++ move semantics / rvalue references
- ❌ Multiple inheritance (except interface-style pure virtual)

**For complex C++ libraries:** Write a C wrapper and use `extern libc`.

## Compatibility Matrix

| Feature | Arc Struct | C++ POD Struct | C++ Class | Status |
|---------|-----------|----------------|-----------|--------|
| Field layout | ✅ | ✅ | ⚠️ | Compatible |
| Alignment | ✅ | ✅ | ⚠️ | Compatible |
| Size | ✅ | ✅ | ⚠️ | Compatible |
| Pass by value | ✅ | ✅ | ❌ | Compatible (POD only) |
| Pass by pointer | ✅ | ✅ | ✅ | Compatible |
| Methods | ✅ | N/A | ✅ | Needs vtable impl |
| Constructors | ❌ | N/A | ✅ | TODO |
| Destructors | ✅ (`deinit`) | N/A | ✅ | TODO (mapping) |
| Virtual functions | ❌ | N/A | ✅ | TODO |
| Inheritance | ❌ | N/A | ⚠️ | Out of scope |

**Legend:**
- ✅ Fully compatible
- ⚠️ Compatible for simple cases
- ❌ Not compatible
- N/A: Not applicable

## Implementation Phases

### Phase 1: Windows DirectX Support (MVP)
- [x] Struct layout compatibility (DONE ✅)
- [ ] MSVC name mangling
- [ ] `__thiscall` convention
- [ ] Vtable offset calculation
- [ ] COM interface support (QueryInterface, AddRef, Release)

**Deliverable:** Call DirectX APIs from Arc

### Phase 2: Constructor/Destructor
- [ ] Constructor mapping
- [ ] Destructor mapping  
- [ ] RAII patterns with `defer`

**Deliverable:** Use C++ RAII libraries

### Phase 3: Cross-Platform
- [ ] Itanium name mangling (Linux/macOS)
- [ ] System V AMD64 calling convention (done ✅)
- [ ] Platform detection for mangling scheme

**Deliverable:** Linux C++ interop

### Phase 4: Advanced Features
- [ ] Function overloading
- [ ] Operator overloading
- [ ] Const methods
- [ ] Reference parameters
- [ ] STL wrappers

**Deliverable:** Ergonomic C++ API access

## Current Workarounds

Until full `extern cpp` is implemented, use these patterns:

**1. C Wrapper for Complex C++**
```cpp
// wrapper.cpp (compile with gcc/clang)
extern "C" {
    void* physics_create_world() {
        return new PhysicsWorld();
    }
    
    void physics_step(void* world, float dt) {
        static_cast<PhysicsWorld*>(world)->step(dt);
    }
    
    void physics_destroy(void* world) {
        delete static_cast<PhysicsWorld*>(world);
    }
}
```

```arc
// Arc
extern libc "physics_wrapper.lib" {
    func physics_create_world() *void
    func physics_step(world: *void, dt: float32) void
    func physics_destroy(world: *void) void
}
```

**2. Manual Vtable Handling**
```arc
// For COM interfaces, manually track vtable offsets
struct ID3D11DeviceVtbl {
    query_interface: *void  // [0]
    add_ref: *void          // [1]
    release: *void          // [2]
    create_buffer: *void    // [3]
    // ... more methods
}

// Cast and call manually (temporary until compiler handles it)
let device_vtbl = cast<*ID3D11DeviceVtbl>(*device)
// Call via function pointer...
```

## Summary

**Today:** Arc structs are binary-compatible with C++ POD types ✅

**To ship `extern cpp`:**
1. Implement name mangling (MSVC + Itanium)
2. Add vtable call generation
3. Support calling conventions (`__thiscall`)
4. Map constructors/destructors

**Goal:** Interface with 90% of real-world C++ libraries without writing C wrappers.

**Non-goal:** Become a C++ compiler. For complex C++ (templates, multiple inheritance, etc.), use C wrappers.