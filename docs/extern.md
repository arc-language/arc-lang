# Foreign Function Interface (`extern`)

Arc provides seamless interoperability with C and C++ libraries through the `extern` keyword. Arc's ABI is already compatible with C/C++ calling conventions, so `extern` primarily serves as a linker hint and mangling directive.

## Quick Reference

```arc
// C functions (no name mangling)
extern c {
    func printf(*byte, ...) int32
    func malloc(usize) *void
}

// C++ functions (name mangling + vtables)
extern cpp {
    namespace DirectX {
        class ID3D11Device {
            virtual func Release(self *ID3D11Device) uint32
        }
        func CreateDevice(...) HRESULT
    }
}
```

---

## `extern c`

Use `extern c` for C libraries. No name mangling—symbols are used as-is.

### Basic Usage

```arc
namespace mymodule

extern c {
    // Standard C library
    func printf(*byte, ...) int32
    func sprintf(*byte, *byte, ...) int32
    func malloc(usize) *void
    func free(*void) void
    func memcpy(*void, *void, usize) *void
    
    // Math library
    func sin(float64) float64
    func cos(float64) float64
    func sqrt(float64) float64
}

func main() {
    let ptr = malloc(1024)
    defer free(ptr)
    
    printf("Allocated: %p\n", ptr)
    printf("sqrt(2) = %f\n", sqrt(2.0))
}
```

### Symbol Renaming

Use a string literal after the function name to specify the actual symbol:

```arc
extern c {
    // Arc name -> C symbol
    func print "printf" (*byte, ...) int32
    func alloc "malloc" (usize) *void
    func c_free "free" (*void) void
    
    // Platform-specific naming
    func sleep "_sleep" (uint32) void      // Windows
}

// Usage
print("Hello %s\n", "World")
let ptr = alloc(100)
c_free(ptr)
```

### Structs

Arc structs are binary-compatible with C structs. Just define matching layouts:

```arc
// Arc struct - matches C layout exactly
struct FILE {}  // Opaque

struct timeval {
    tv_sec: int64
    tv_usec: int64
}

struct stat {
    st_dev: uint64
    st_ino: uint64
    st_mode: uint32
    st_nlink: uint64
    st_uid: uint32
    st_gid: uint32
    st_size: int64
    // ...
}

extern c {
    func fopen(*byte, *byte) *FILE
    func fclose(*FILE) int32
    func fread(*void, usize, usize, *FILE) usize
    
    func gettimeofday(*timeval, *void) int32
    func stat_file "stat" (*byte, *stat) int32
}
```

### Variadic Functions

```arc
extern c {
    func printf(*byte, ...) int32
    func scanf(*byte, ...) int32
    func ioctl(int32, uint64, ...) int32
}
```

### Constants

```arc
extern c {
    const O_RDONLY: int32 = 0
    const O_WRONLY: int32 = 1
    const O_RDWR: int32 = 2
    const O_CREAT: int32 = 64
    
    func open(*byte, int32, ...) int32
    func close(int32) int32
    func read(int32, *void, usize) isize
    func write(int32, *void, usize) isize
}

func main() {
    let fd = open("test.txt", O_RDONLY)
    if fd >= 0 {
        defer close(fd)
        // ...
    }
}
```

### Callbacks / Function Pointers

Pass Arc functions to C APIs that expect callbacks:

```arc
extern c {
    // Function pointer types
    type Comparator = func(*const void, *const void) int32
    type SignalHandler = func(int32) void
    type ThreadFunc = func(*void) *void
    
    // Functions that take callbacks
    func qsort(*void, usize, usize, Comparator) void
    func signal(int32, SignalHandler) SignalHandler
    func pthread_create(*pthread_t, *void, ThreadFunc, *void) int32
    
    // Inline function pointer syntax also works
    func atexit(func() void) int32
    func bsearch(*void, *void, usize, usize, func(*const void, *const void) int32) *void
}

// Arc functions with matching signatures can be passed directly
func compare_ints(a: *const void, b: *const void) int32 {
    let x = *(a as *const int32)
    let y = *(b as *const int32)
    return x - y
}

func on_signal(sig: int32) void {
    printf("Received signal: %d\n", sig)
}

func main() {
    let arr = [5]int32{5, 2, 8, 1, 9}
    
    // Pass Arc function as callback
    qsort(&arr, 5, sizeof<int32>, compare_ints)
    
    // Register signal handler
    signal(SIGINT, on_signal)
}
```

---

## `extern cpp`

Use `extern cpp` for C++ libraries. Handles name mangling and vtable calls.

### Basic Usage

```arc
namespace graphics

extern cpp {
    func CreateDevice(*void, uint32) *Device
    func DestroyDevice(*Device) void
}

// Access
graphics.CreateDevice(adapter, flags)
```

### Namespaces

C++ namespaces inside `extern cpp` create nested access paths and control name mangling. You can declare namespaces two ways:

**Nested blocks** — best for deep hierarchies:

```arc
namespace graphics

extern cpp {
    namespace DirectX {
        func CreateDevice(...) HRESULT
        func CreateFactory(...) HRESULT
        
        namespace DXGI {
            func GetAdapter(uint32) *IDXGIAdapter
            func CreateSwapChain(...) HRESULT
            
            namespace Debugging {
                func ReportLiveObjects(...) void
            }
        }
    }
}

// Access mirrors the declaration structure:
graphics.DirectX.CreateDevice(...)                    // → DirectX::CreateDevice
graphics.DirectX.DXGI.GetAdapter(0)                   // → DirectX::DXGI::GetAdapter
graphics.DirectX.DXGI.Debugging.ReportLiveObjects()   // → DirectX::DXGI::Debugging::ReportLiveObjects
```

**Dot notation** — best for sparse or one-off declarations:

```arc
namespace mymodule

extern cpp {
    namespace std.chrono {
        func now() TimePoint
    }
    
    namespace std.filesystem {
        func exists(*Path) bool
        func create_directory(*Path) bool
    }
    
    namespace boost.asio.ip {
        func make_address(*byte) Address
    }
}

// main.arc
namespace main

func main () {
    // Access:
    mymodule.std.chrono.now()
    mymodule.std.filesystem.exists(&path)
    mymodule.boost.asio.ip.make_address("127.0.0.1")
}
```

**Mix both styles** as needed:

```arc
extern cpp {
    namespace DirectX {
        func CreateDevice(...) HRESULT
        
        namespace DXGI {
            func GetAdapter(uint32) *IDXGIAdapter
        }
    }
    
    // Equivalent to nesting inside DirectX.D3D12
    namespace DirectX.D3D12 {
        func CreateCommandQueue(...) HRESULT
    }
}
```

### Classes and Virtual Methods

Use `class` for C++ classes with vtables. Mark virtual methods with `virtual`:

```arc
extern cpp {
    namespace DirectX {
        class ID3D11Device {
            // IUnknown (vtable slots 0-2)
            virtual func QueryInterface(self *ID3D11Device, *GUID, **void) HRESULT
            virtual func AddRef(self *ID3D11Device) uint32
            virtual func Release(self *ID3D11Device) uint32
            
            // ID3D11Device methods (vtable slot 3+)
            virtual func CreateBuffer(
                self *ID3D11Device,
                *D3D11_BUFFER_DESC,
                *D3D11_SUBRESOURCE_DATA,
                **ID3D11Buffer
            ) HRESULT
        }
    }
}
```

**Generated code for virtual calls:**
```asm
; device.Release()
mov rax, [rcx]        ; Load vtable pointer
call [rax + 16]       ; Call vtable[2] (offset 2 * 8 bytes)
```

### Constructors and Destructors

```arc
extern cpp {
    class Widget {
        // Constructor - returns pointer to new instance
        new(int32, int32) *Widget
        
        // Multiple constructors (overloaded)
        new() *Widget
        new(*byte) *Widget
        
        // Destructor
        delete(self *Widget) void
        
        virtual func Process(self *Widget) void
    }
}

// Usage
let w = Widget.new(100, 200)
defer w.delete()

w.Process()
```

### Static Methods

Methods without `self` are static:

```arc
extern cpp {
    class Factory {
        // Static methods - no self parameter
        static func Create(*byte) *Factory
        static func GetInstance() *Factory
        static func GetVersion() int32
        
        // Instance method - has self
        virtual func Process(self *Factory) void
        
        delete(self *Factory) void
    }
}

// Usage
let f = Factory.Create("config.json")
defer f.delete()

let version = Factory.GetVersion()
let singleton = Factory.GetInstance()
```

### Opaque Types

Use `opaque` when you don't know (or care about) the internal layout:

```arc
extern cpp {
    namespace Graphics {
        // Opaque - can only hold pointers, no field access
        opaque class RenderContext {}
        opaque class FileHandle {}
        opaque class DatabaseConnection {}
        
        func CreateContext() *RenderContext
        func DestroyContext(*RenderContext) void
    }
}

// Can only use as pointer
let ctx = Graphics.CreateContext()
defer Graphics.DestroyContext(ctx)

// Cannot access internals - it's opaque
// ctx.someField  // ❌ Error
```

### Abstract Classes

Use `abstract` for classes with pure virtual methods:

```arc
namespace main

extern cpp {
    // Abstract - cannot instantiate directly
    abstract class IRenderer {
        virtual func Init(self *IRenderer) bool
        virtual func Draw(self *IRenderer, *Scene) void
        virtual func Shutdown(self *IRenderer) void
    }
    
    // Concrete implementation
    class D3D11Renderer {
        new() *D3D11Renderer
        delete(self *D3D11Renderer) void
        
        virtual func Init(self *D3D11Renderer) bool
        virtual func Draw(self *D3D11Renderer, *Scene) void
        virtual func Shutdown(self *D3D11Renderer) void
    }
}

// Can use abstract as parameter type
func render_frame(renderer: *IRenderer, scene: *Scene) {
    renderer.Draw(scene)
}

func main() {
    let renderer = D3D11Renderer.new()
    defer renderer.delete()
    
    renderer.Init()
    render_frame(renderer, &scene)
    renderer.Shutdown()
}
```

### Const Methods

Use `const` after the parameter list for const methods (affects mangling):

```arc
extern cpp {
    class Buffer {
        new(usize) *Buffer
        delete(self *Buffer) void
        
        // const methods - read only, can call on *const Buffer
        virtual func Size(self *const Buffer) const usize
        virtual func Data(self *const Buffer) const *byte
        virtual func IsEmpty(self *const Buffer) const bool
        
        // non-const methods - can modify
        virtual func Resize(self *Buffer, usize) void
        virtual func Clear(self *Buffer) void
        virtual func Write(self *Buffer, *byte, usize) void
    }
}
```

### Reference Parameters

Map C++ references with `&`:

```arc
extern cpp {
    class Container {
        // const ref - read only, no copy
        virtual func Find(self *Container, &const Key) *Value
        
        // mutable ref - can modify
        virtual func Swap(self *Container, &int32, &int32) void
        
        // Return by reference
        virtual func At(self *Container, usize) &Item
        virtual func Front(self *const Container) const &Item
    }
    
    // Standalone functions with refs
    namespace std {
        func swap(&int32, &int32) void
    }
}
```

### Function Overloading

Same function name with different signatures—mangler generates unique symbols:

```arc
extern cpp {
    namespace Math {
        // Same name, different parameter types
        func Clamp(int32, int32, int32) int32
        func Clamp(float32, float32, float32) float32
        func Clamp(float64, float64, float64) float64
        
        func Abs(int32) int32
        func Abs(float32) float32
        func Abs(float64) float64
        
        func Min(int32, int32) int32
        func Min(float32, float32) float32
    }
}

// Compiler picks correct overload based on argument types
let i = Math.Clamp(x, 0, 100)        // int32 version
let f = Math.Clamp(x, 0.0, 1.0)      // float32 version
```

If auto-mangling fails, use explicit symbol:

```arc
extern cpp {
    namespace Math {
        func ClampInt "?Clamp@Math@@YAHHHH@Z" (int32, int32, int32) int32
        func ClampFloat "?Clamp@Math@@YAMMM@Z" (float32, float32, float32) float32
    }
}
```

### Template Instantiations

For C++ templates, declare the specific instantiation you need:

```arc
extern cpp {
    namespace std {
        // std::vector<int>
        class IntVector "std::vector<int>" {
            new() *IntVector
            delete(self *IntVector) void
            
            virtual func push_back(self *IntVector, int32) void
            virtual func pop_back(self *IntVector) void
            virtual func size(self *const IntVector) const usize
            virtual func capacity(self *const IntVector) const usize
            virtual func at(self *IntVector, usize) &int32
            virtual func data(self *IntVector) *int32
            virtual func clear(self *IntVector) void
            virtual func empty(self *const IntVector) const bool
        }
        
        // std::string
        class String "std::string" {
            new() *String
            new(*byte) *String
            delete(self *String) void
            
            virtual func c_str(self *const String) const *byte
            virtual func size(self *const String) const usize
            virtual func empty(self *const String) const bool
            virtual func clear(self *String) void
        }
        
        // std::unordered_map<std::string, int>
        class StringIntMap "std::unordered_map<std::string, int>" {
            new() *StringIntMap
            delete(self *StringIntMap) void
            
            virtual func size(self *const StringIntMap) const usize
            virtual func clear(self *StringIntMap) void
        }
    }
}

// Usage
let vec = std.IntVector.new()
defer vec.delete()

vec.push_back(10)
vec.push_back(20)
printf("Size: %zu\n", vec.size())
```

### Symbol Override

When auto-mangling fails, specify the exact symbol:

```arc
extern cpp {
    namespace DirectX {
        // Explicit mangled name (MSVC)
        func CreateDevice "?CreateDevice@DirectX@@YAPEAVDevice@@PEAXK@Z" (
            *void, 
            uint32
        ) *Device
        
        class ID3D11Device {
            virtual func Release "?Release@ID3D11Device@@UEAAKXZ" (
                self *ID3D11Device
            ) uint32
        }
    }
}
```

### Callbacks / Function Pointers

```arc
extern cpp {
    // Typedef for callbacks
    type EventCallback = func(*Event, *void) void
    type CompareFunc = func(*const void, *const void) int32
    
    class EventSystem {
        new() *EventSystem
        delete(self *EventSystem) void
        
        virtual func Subscribe(self *EventSystem, *byte, EventCallback, *void) void
        virtual func Unsubscribe(self *EventSystem, *byte, EventCallback) void
    }
    
    namespace Algorithm {
        func Sort(*void, usize, usize, CompareFunc) void
    }
}

// Arc function matching callback signature
func on_event(event: *Event, user_data: *void) void {
    printf("Event received!\n")
}

func main() {
    let events = EventSystem.new()
    defer events.delete()
    
    events.Subscribe("click", on_event, null)
}
```

---

## Struct Attributes

### Alignment

Force specific alignment for SIMD or hardware requirements:

```arc
@align(16)
struct XMVECTOR {
    x: float32
    y: float32
    z: float32
    w: float32
}

@align(32)
struct AVXData {
    values: [8]float32
}

@align(4096)
struct PageAligned {
    data: [4096]byte
}
```

### Packed

Remove padding for binary protocols or file formats:

```arc
@packed
struct FileHeader {
    magic: uint32      // offset 0
    version: uint16    // offset 4 (no padding)
    flags: uint16      // offset 6
    size: uint32       // offset 8
}
// Total size: 12 bytes (not 16)

@packed
struct NetworkPacket {
    type: uint8
    length: uint16
    sequence: uint32
    // ...
}
```

---

## Calling Conventions

Specify calling convention when needed (primarily Windows x86):

```arc
extern c {
    // Default: cdecl on x86, System V / Microsoft x64 ABI on x64
    func printf(*byte, ...) int32
    
    // Windows API (stdcall on x86, ignored on x64)
    stdcall func MessageBoxA(*void, *byte, *byte, uint32) int32
    stdcall func GetLastError() uint32
    stdcall func ExitProcess(uint32) void
}

extern cpp {
    class SomeClass {
        // thiscall: this pointer in ECX (x86) or RCX (x64)
        thiscall virtual func Method(self *SomeClass, int32) void
    }
    
    // vectorcall: SIMD-friendly, floats in XMM registers
    vectorcall func XMVectorAdd(XMVECTOR, XMVECTOR) XMVECTOR
}
```

| Convention | Description | Platform |
|------------|-------------|----------|
| `cdecl` | Caller cleans stack (default) | All |
| `stdcall` | Callee cleans stack | Windows x86 |
| `thiscall` | `this` in ECX/RCX | MSVC C++ |
| `vectorcall` | SIMD registers | Windows |
| `fastcall` | First args in registers | Legacy |

---

## Complete Examples

### DirectX 11

```arc
namespace graphics.d3d11

type HRESULT = int32
type UINT = uint32
type GUID = [16]byte

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

const D3D11_SDK_VERSION: uint32 = 7
const D3D_DRIVER_TYPE_HARDWARE: uint32 = 1
const D3D11_USAGE_DEFAULT: uint32 = 0
const D3D11_BIND_VERTEX_BUFFER: uint32 = 1

extern cpp {
    namespace DirectX {
        opaque class IDXGIAdapter {}
        opaque class ID3D11DeviceContext {}
        opaque class ID3D11Texture2D {}
        
        func D3D11CreateDevice(
            *IDXGIAdapter,
            uint32,
            *void,
            uint32,
            *uint32,
            uint32,
            uint32,
            **ID3D11Device,
            *uint32,
            **ID3D11DeviceContext
        ) HRESULT
        
        class ID3D11Device {
            virtual func QueryInterface(self *ID3D11Device, *GUID, **void) HRESULT
            virtual func AddRef(self *ID3D11Device) uint32
            virtual func Release(self *ID3D11Device) uint32
            virtual func CreateBuffer(
                self *ID3D11Device,
                *D3D11_BUFFER_DESC,
                *D3D11_SUBRESOURCE_DATA,
                **ID3D11Buffer
            ) HRESULT
        }
        
        class ID3D11Buffer {
            virtual func QueryInterface(self *ID3D11Buffer, *GUID, **void) HRESULT
            virtual func AddRef(self *ID3D11Buffer) uint32
            virtual func Release(self *ID3D11Buffer) uint32
        }
    }
}

extern c {
    func printf(*byte, ...) int32
}

func main() {
    let device: *DirectX.ID3D11Device = null
    let context: *DirectX.ID3D11DeviceContext = null
    
    let hr = DirectX.D3D11CreateDevice(
        null,
        D3D_DRIVER_TYPE_HARDWARE,
        null,
        0,
        null,
        0,
        D3D11_SDK_VERSION,
        &device,
        null,
        &context
    )
    
    if hr != 0 {
        printf("Failed: %d\n", hr)
        return
    }
    defer device.Release()
    defer context.Release()
    
    let vertices = array<float32, 9>{
        0.0, 0.5, 0.0,
        0.5, -0.5, 0.0,
        -0.5, -0.5, 0.0
    }
    
    let desc = D3D11_BUFFER_DESC{
        byte_width: 36,
        usage: D3D11_USAGE_DEFAULT,
        bind_flags: D3D11_BIND_VERTEX_BUFFER,
        cpu_access_flags: 0,
        misc_flags: 0,
        structure_byte_stride: 0
    }
    
    let init_data = D3D11_SUBRESOURCE_DATA{
        sys_mem: &vertices as *void,
        sys_mem_pitch: 0,
        sys_mem_slice_pitch: 0
    }
    
    let buffer: *DirectX.ID3D11Buffer = null
    hr = device.CreateBuffer(&desc, &init_data, &buffer)
    
    if hr == 0 {
        defer buffer.Release()
        printf("Buffer created!\n")
    }
}
```

### SQLite (C)

```arc
namespace main

extern c {
    opaque struct sqlite3 {}
    opaque struct sqlite3_stmt {}
    
    const SQLITE_OK: int32 = 0
    const SQLITE_ROW: int32 = 100
    const SQLITE_DONE: int32 = 101
    
    func sqlite3_open(*byte, **sqlite3) int32
    func sqlite3_close(*sqlite3) int32
    func sqlite3_errmsg(*sqlite3) *byte
    
    func sqlite3_prepare_v2(*sqlite3, *byte, int32, **sqlite3_stmt, **byte) int32
    func sqlite3_step(*sqlite3_stmt) int32
    func sqlite3_finalize(*sqlite3_stmt) int32
    func sqlite3_reset(*sqlite3_stmt) int32
    
    func sqlite3_bind_int(*sqlite3_stmt, int32, int32) int32
    func sqlite3_bind_text(*sqlite3_stmt, int32, *byte, int32, *void) int32
    
    func sqlite3_column_int(*sqlite3_stmt, int32) int32
    func sqlite3_column_text(*sqlite3_stmt, int32) *byte
    
    func printf(*byte, ...) int32
}

func main() {
    let db: *sqlite3 = null
    
    if sqlite3_open("test.db", &db) != SQLITE_OK {
        printf("Failed to open: %s\n", sqlite3_errmsg(db))
        return
    }
    defer sqlite3_close(db)
    
    let stmt: *sqlite3_stmt = null
    let sql = "SELECT id, name FROM users"
    
    if sqlite3_prepare_v2(db, sql, -1, &stmt, null) != SQLITE_OK {
        printf("Failed to prepare: %s\n", sqlite3_errmsg(db))
        return
    }
    defer sqlite3_finalize(stmt)
    
    for sqlite3_step(stmt) == SQLITE_ROW {
        let id = sqlite3_column_int(stmt, 0)
        let name = sqlite3_column_text(stmt, 1)
        printf("User %d: %s\n", id, name)
    }
}
```

### Mixed C and C++ with Nested Namespaces

```arc
namespace main

extern c {
    func printf(*byte, ...) int32
    func malloc(usize) *void
    func free(*void) void
}

extern cpp {
    namespace Physics {
        class World {
            new() *World
            delete(self *World) void
            
            virtual func Step(self *World, float32) void
            virtual func GetGravity(self *const World) const Vec3
        }
        
        namespace Collision {
            abstract class Shape {}
            
            class BoxShape {
                new(float32, float32, float32) *BoxShape
                delete(self *BoxShape) void
            }
            
            class SphereShape {
                new(float32) *SphereShape
                delete(self *SphereShape) void
            }
            
            func TestCollision(*Shape, *Shape) bool
        }
    }
    
    namespace Audio.Effects {
        class Reverb {
            new() *Reverb
            delete(self *Reverb) void
            
            virtual func SetDecay(self *Reverb, float32) void
            virtual func Process(self *Reverb, *float32, usize) void
        }
    }
}

func main() {
    let world = Physics.World.new()
    defer world.delete()
    
    let box = Physics.Collision.BoxShape.new(1.0, 2.0, 1.0)
    defer box.delete()
    
    let sphere = Physics.Collision.SphereShape.new(0.5)
    defer sphere.delete()
    
    if Physics.Collision.TestCollision(box, sphere) {
        printf("Collision detected!\n")
    }
    
    let reverb = Audio.Effects.Reverb.new()
    defer reverb.delete()
    
    reverb.SetDecay(0.8)
    
    world.Step(0.016)
}
```

---

## Linkage

Library paths are specified in `build.arc`, not source files:

```arc
// build.arc
module graphics {
    sources: ["src/*.arc"]
    
    link: {
        windows: ["d3d11.lib", "dxgi.lib"]
        linux: ["libvulkan.so"]
        macos: ["Metal.framework"]
        
        sqlite: find("sqlite3")
    }
}
```

Source files just declare the interface:

```arc
// src/graphics.arc
extern cpp {
    namespace DirectX {
        func CreateDevice(...) HRESULT
    }
}
```

---

## Comparison: `extern c` vs `extern cpp`

| Feature | `extern c` | `extern cpp` |
|---------|------------|--------------|
| Name mangling | None | C++ ABI mangling |
| Namespaces | N/A | Nested access paths |
| Classes | N/A | Full support |
| Virtual methods | N/A | `virtual` keyword |
| Constructors | N/A | `new(...) *T` |
| Destructors | N/A | `delete(self *T)` |
| Static methods | N/A | `static func` |
| Opaque types | `opaque struct` | `opaque class` |
| Abstract types | N/A | `abstract class` |
| Overloading | N/A | Via mangling |
| Symbol override | `func name "symbol"` | `func name "symbol"` |
| Function pointers | ✅ | ✅ |

---

## Quick Reference

```arc
// C
extern c {
    func name(types...) ReturnType
    func arc_name "c_symbol" (types...) ReturnType
    type Callback = func(types...) ReturnType
    const NAME: Type = value
    opaque struct Handle {}
    
    // Calling conventions
    stdcall func WinApiFunc(...) int32
}

// C++
extern cpp {
    // Nested namespace blocks
    namespace Outer {
        func OuterFunc() void
        
        namespace Inner {
            func InnerFunc() void
        }
    }
    
    // Dot notation for namespaces
    namespace Outer.Inner.Deep {
        func DeepFunc() void
    }
    
    // Classes
    opaque class OpaqueHandle {}
    
    abstract class IInterface {
        virtual func Method(self *IInterface) void
    }
    
    class ClassName "optional::mangled::name" {
        new(types...) *ClassName
        delete(self *ClassName) void
        
        static func StaticMethod() *ClassName
        
        virtual func Method(self *ClassName, types...) ReturnType
        virtual func ConstMethod(self *const ClassName) const ReturnType
        virtual func RefParam(self *ClassName, &Type, &const Type) void
    }
    
    // Type aliases
    type Callback = func(*Event) void
    
    // Calling conventions
    vectorcall func SimdFunc(Vec4, Vec4) Vec4
}

// Struct attributes
@align(16)
struct Aligned { ... }

@packed
struct Packed { ... }
```