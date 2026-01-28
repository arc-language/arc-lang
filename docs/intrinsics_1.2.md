# arc Compiler Intrinsics (Version 1.2)

> **Note**: All intrinsics are compiler-handled functions, not parser keywords.
> They are mapped in the compiler's intrinsic registry, not built into the grammar.

## Stack Allocation

### alloca - Allocate on Stack
```arc
// Allocate single item on stack (returns *T)
// Essential for creating kernel buffers without malloc
let ptr = alloca<int32>(int32)

// Allocate array/buffer on stack (returns *T, second arg is count)
let buffer = alloca<byte>(1024)
```

## Meta Information

### sizeof - Compile-Time Size
```arc
// Compile-time size in bytes (returns usize)
// Essential for passing buffer sizes to read/write/mmap
let sz = sizeof(int32)       // 4
let st_sz = sizeof(Stat)     // Struct size (with padding)
```

### alignof - Compile-Time Alignment
```arc
// Compile-time alignment requirement (returns usize)
let align = alignof(float64) // 8
let align_struct = alignof(Point) // Alignment of struct
```

## Memory Block Operations

### memset - Fill Memory
```arc
let buf = alloca(byte, 1024)

// memset(dest: *void, val: byte, count: usize)
// Essential for zeroing stack structs before syscalls
memset(buf, 0, 1024)

// Example: Zero-initialize a struct
let stat_buf = alloca(Stat)
memset(stat_buf, 0, sizeof(Stat))
```

### memcpy - Copy Memory (Non-Overlapping)
```arc
// memcpy(dest: *void, src: *void, count: usize)
// Essential for copying data to/from kernel buffers
// UNSAFE if regions overlap
memcpy(dest_ptr, src_ptr, 1024)

// Example: Copy struct data
let src = Point{x: 10, y: 20}
let dest = alloca(Point)
memcpy(dest, &src, sizeof(Point))
```

### memmove - Copy Memory (Safe for Overlap)
```arc
// memmove(dest: *void, src: *void, count: usize)
// Like memcpy but handles overlapping regions safely
// Essential for buffer manipulation where source/dest might overlap
memmove(dest_ptr, src_ptr, 1024)

// Example: Shift buffer contents
let buffer = alloca(byte, 1024)
memmove(buffer, buffer + 100, 900)  // Safe overlap
```

### memcmp - Compare Memory
```arc
// memcmp(ptr1: *void, ptr2: *void, count: usize) -> int32
// Returns 0 if equal, <0 if ptr1 < ptr2, >0 if ptr1 > ptr2
// Essential for implementing '==' for strings and structs
let diff = memcmp(ptr1, ptr2, 1024)

if memcmp(&struct1, &struct2, sizeof(Point)) == 0 {
    // Structs are equal
}
```

## String Operations

### strlen - C-String Length
```arc
// strlen(str: *byte) -> usize
// Calculate C-string length (stops at null terminator)
// Essential when interfacing with syscalls expecting null-terminated strings
let cstr: *byte = "hello\0"
let len = strlen(cstr)  // 5
```

### memchr - Find Byte in Memory
```arc
// memchr(ptr: *void, val: byte, count: usize) -> *void
// Find first occurrence of byte in memory region
// Returns pointer to found byte, or null if not found
// Essential for parsing, finding newlines in buffers, etc.
let buf: *byte = "hello\nworld"
let newline = memchr(buf, '\n', 11)  // Points to '\n'

if newline != null {
    let offset = cast<usize>(newline) - cast<usize>(buf)
}
```

## Variadic Arguments

### va_start - Initialize Argument List
```arc
// Essential for implementing printf without libc
func printf(fmt: string, ...) {
    // Initialize argument walker using the last known arg
    let args = va_start(fmt)
    defer va_end(args)
    
    // Process arguments...
}
```

### va_arg - Get Next Argument
```arc
func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)
    
    // Retrieve next argument as specific type
    let val = va_arg<int32>(args)
    let str = va_arg<string>(args)
    let ptr = va_arg<*byte>(args)
}
```

### va_end - Cleanup Argument List
```arc
func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)  // Always cleanup
    
    // Process arguments...
}
```

## Process Control

### raise - Abort with Message
```arc
// Immediately abort execution with a message
// Internally calls SYS_EXIT or triggers SIGABRT
// Used for unrecoverable errors/crashes
if ptr == null {
    raise("Memory corrupted")
}

if buffer_size > MAX_SIZE {
    raise("Buffer overflow detected")
}
```

## Bit Manipulation

### bit_cast - Reinterpret Bits
```arc
// bit_cast<TargetType>(SourceValue)
// Reinterprets the raw bits without conversion
// Example: getting the IEEE754 integer representation of a float
let f: float32 = 1.0
let bits = bit_cast<uint32>(f) // 0x3F800000, not 1

// Example: type punning for performance
let i: int32 = -1
let u = bit_cast<uint32>(i)  // 0xFFFFFFFF

// Example: pointer to integer for low-level operations
let ptr: *void = some_address
let addr = bit_cast<usize>(ptr)
```

## System Calls

### syscall - Direct System Call
```arc
// syscall(number, arg1, arg2, arg3, arg4, arg5, arg6)
// Direct interface to kernel syscalls (Linux, BSD, etc.)
// Up to 6 arguments supported

let msg = "Hello, Direct Syscall!\n"
let len = 23

// Example: SYS_WRITE
let result = syscall(SYS_WRITE, STDOUT, msg, len)

// Example: SYS_OPEN
let fd = syscall(SYS_OPEN, path_ptr, O_RDONLY, 0)

// Example: SYS_MMAP
let addr = syscall(SYS_MMAP, null, 4096, PROT_READ | PROT_WRITE, 
                   MAP_PRIVATE | MAP_ANONYMOUS, -1, 0)
```

## Slicing

### slice - Create Array/Vector View
```arc
// slice(collection, range) -> (*T, usize)
// Returns pointer and length for a range of elements
// Essential for zero-copy subarray operations

let arr: array<int32, 10> = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
let (ptr, len) = slice(arr, 1..4)  // ptr points to arr[1], len = 3

// Use with vectors
let vec: vector<byte> = {1, 2, 3, 4, 5}
let (data_ptr, data_len) = slice(vec, 0..3)

// Pass slice to syscall
let buffer: vector<byte> = load_file("data.bin")
let (chunk_ptr, chunk_len) = slice(buffer, 100..200)
syscall(SYS_WRITE, STDOUT, chunk_ptr, chunk_len)
```

## Usage Notes

**Memory Safety:**
- `memcpy` is faster but UNSAFE for overlapping regions
- `memmove` is safe for overlaps but slightly slower
- Always use `defer` with `va_end`

**Performance:**
- `alloca` is zero-cost (just stack pointer adjustment)
- `sizeof` and `alignof` are compile-time constants (zero runtime cost)
- `memset`, `memcpy`, `memmove` often compile to optimized CPU instructions

**Common Patterns:**
```arc
// Zero-init pattern
let buf = alloca(byte, 1024)
memset(buf, 0, 1024)

// Struct copy pattern
let copy = alloca(Point)
memcpy(copy, &original, sizeof(Point))

// Syscall buffer pattern
let buffer = alloca(byte, 4096)
let bytes_read = syscall(SYS_READ, fd, buffer, 4096)
```