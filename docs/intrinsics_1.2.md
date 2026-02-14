# arc Compiler Intrinsics (Version 1.4)

> **Note**: All intrinsics are compiler-handled functions, not parser keywords.
> They are mapped in the compiler's intrinsic registry, not built into the grammar.

## Meta Information

### sizeof - Compile-Time Size
```arc
// Compile-time size in bytes (returns usize)
let sz = sizeof(int32)       // 4
let st_sz = sizeof(Stat)     // Struct size (with padding)
```

### alignof - Compile-Time Alignment
```arc
// Compile-time alignment requirement (returns usize)
let align = alignof(float64)  // 8
let align_struct = alignof(Point)
```

## Memory Block Operations

### alloca - Stack Allocation
```arc
// alloca(type, count?) - stack allocate, zero cost
let buf = alloca(byte, 1024)
let stat_buf = alloca(Stat)
```

### memset - Fill Memory
```arc
let buf = alloca(byte, 1024)

// memset(dest, val, count)
memset(buf, 0, 1024)

// Zero-initialize a struct before syscalls
let stat_buf = alloca(Stat)
memset(stat_buf, 0, sizeof(Stat))
```

### memcpy - Copy Memory (Non-Overlapping)
```arc
// memcpy(dest, src, count)
// UNSAFE if regions overlap
memcpy(dest_ptr, src_ptr, 1024)

// Copy a struct
let src = Point{x: 10, y: 20}
let dest = alloca(Point)
memcpy(dest, src, sizeof(Point))
```

### memmove - Copy Memory (Safe for Overlap)
```arc
// memmove(dest, src, count)
// Safe for overlapping regions
memmove(dest_ptr, src_ptr, 1024)

// Shift buffer contents
let buffer = alloca(byte, 1024)
memmove(buffer, buffer + 100, 900)
```

### memcmp - Compare Memory
```arc
// memcmp(ptr1, ptr2, count) -> int32
// Returns 0 if equal, <0 if ptr1 < ptr2, >0 if ptr1 > ptr2
let diff = memcmp(ptr1, ptr2, 1024)

if memcmp(struct1, struct2, sizeof(Point)) == 0 {
    // Structs are equal
}
```

## String Operations

### strlen - C-String Length
```arc
// strlen(str) -> usize
// Used when interfacing with C APIs expecting null-terminated strings
let len = strlen(cstr)  // 5
```

### memchr - Find Byte in Memory
```arc
// memchr(ptr, val, count) -> rawptr
// Find first occurrence of byte, returns null if not found
let newline = memchr(buf, '\n', 11)

if newline != null {
    let offset = usize(newline) - usize(buf)
}
```

## Variadic Arguments

### va_start - Initialize Argument List
```arc
func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)
}
```

### va_arg - Get Next Argument
```arc
func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)

    let val = va_arg(args, int32)
    let str = va_arg(args, string)
}
```

### va_end - Cleanup Argument List
```arc
func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)  // Always cleanup
}
```

## Bit Manipulation

### bit_cast - Reinterpret Bits
```arc
// bit_cast(type, value)
// Reinterprets raw bits without conversion

// Get IEEE754 integer representation of a float
let f: float32 = 1.0
let bits = bit_cast(uint32, f)  // 0x3F800000

// Type punning
let i: int32 = -1
let u = bit_cast(uint32, i)  // 0xFFFFFFFF
```

## System Calls

### syscall - Direct System Call
```arc
// syscall(number, arg1..arg6)
// Direct kernel interface, up to 6 arguments

let msg = "Hello\n"
let len = 6

// SYS_WRITE
let result = syscall(SYS_WRITE, STDOUT, msg, len)

// SYS_OPEN
let fd = syscall(SYS_OPEN, path, O_RDONLY, 0)

// SYS_MMAP
let addr = syscall(SYS_MMAP, null, 4096,
                   PROT_READ | PROT_WRITE,
                   MAP_PRIVATE | MAP_ANONYMOUS, -1, 0)
```

## Type Casting

### Primitive casts - type(value) syntax
```arc
// Convert between numeric types
let x = int32(3.14)     // float to int
let y = float64(42)     // int to float
let z = uint8(flags)    // narrow cast
let n = usize(count)    // to pointer-sized int
```

### rawptr - Two forms
```arc
// rawptr(value) — cast integer value to raw untyped pointer
// rawptr(&val)  — get address of variable as raw pointer (like & in C)

// 1. Cast integer value to pointer
// SQLITE_TRANSIENT = -1 as pointer, tells sqlite to copy the string
let transient = rawptr(-1)
sqlite3_bind_text(stmt, 1, val, len, transient)

// 2. Get address of a variable
let x = 100
let ptr = rawptr(&x)    // ptr points TO x

// 3. Pointer arithmetic (via usize)
let addr = usize(some_ptr)
let offset_ptr = rawptr(addr + 16)

// 4. Pass address to extern output params
let db: sqlite3 = null
sqlite3_open("test.db", rawptr(&db))   // extern expects **sqlite3
```

## GPU Functions

### gpu func - Accelerator Kernels
```arc
// gpu func - all params are gpu bound
// compiler maps to build target (cuda, metal, xla)
gpu func kernel(data: float32, n: usize) {
    let idx = thread_id()
    data[idx] = data[idx] * 2.0
}

// Await a gpu func
async func main() {
    let result = await kernel(data, n)
}
```

## Usage Notes

**Memory Safety:**
- `memcpy` is faster but UNSAFE for overlapping regions
- `memmove` handles overlaps safely but is slightly slower
- Always use `defer` with `va_end`

**Performance:**
- `alloca` is zero-cost (stack pointer adjustment only)
- `sizeof` and `alignof` are compile-time constants
- `memset`, `memcpy`, `memmove` compile to optimized CPU instructions

**Common Patterns:**
```arc
// Zero-init stack buffer
let buf = alloca(byte, 1024)
memset(buf, 0, 1024)

// Struct copy
let copy = alloca(Point)
memcpy(copy, original, sizeof(Point))

// Syscall buffer pattern
let buffer = alloca(byte, 4096)
let bytes_read = syscall(SYS_READ, fd, buffer, 4096)

// Pass variable address to extern
let result: SomeType = null
extern_func(rawptr(&result))
```