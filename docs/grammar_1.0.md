# arc Language Grammar Snippets (Version 1.0 - Core Features & Intrinsics)

## Comments
```arc
// Single-line comment

/*
   Multi-line comment
   spanning multiple lines
*/
```

## Import, declaration
```arc
import "some/path/package"

import (
    // Standard Lib
    "std/io"
    "std/net"

    // Third Party (Explicit URLs)
    "github.com/user/physics"
    "github.com/user/graphics"
    
    // Other Hosts
    "gitlab.com/company/lib"
)
```

## Namespaces, declaration
```arc
namespace main
```

## Variables, mutable with type
```arc
let x: int32 = 42
x = 100
```

## Variables, mutable with inference
```arc
let x = 42
x = 100
```

## Constants, immutable with type
```arc
const x: int32 = 42
```

## Constants, immutable with inference
```arc
const x = 42
```

## Basic Types, Fixed-Width Integers
```arc
// Signed: int8, int16, int32, int64
let i: int32 = -500

// Unsigned: uint8, uint16, uint32, uint64
let u: uint64 = 10000
```

## Basic Types, Architecture Dependent
```arc
// Unsigned pointer-sized integer (x64 = uint64, x86 = uint32)
// Use for: Array indexing, Memory sizes, Loop counters
let len: usize = 100

// Signed pointer-sized integer (x64 = int64, x86 = int32)
// Use for: Pointer offsets, C functions returning ssize_t
let offset: isize = -4
```

## Basic Types, Floating Point
```arc
let f32: float32 = 3.14
let f64: float64 = 2.71828
```

## Basic Types, Aliases (Semantic)
```arc
// 'byte' is an alias for 'uint8' (Raw data)
let b: byte = 255

// 'bool' is an alias for 'uint8' (1=true, 0=false)
let flag: bool = true

// 'char' is an alias for 'uint32' (Unicode Code Point)
let r: char = 'a'
```

## Basic Types, String (Composite)
```arc
// High-level string (ptr + length)
let s: string = "hello"
```

## Basic Types, Qualified (Namespaced)
```arc
// Type from a specific namespace
let client: net.Socket = ...
let config: json.Config = ...
```

## Literals, boolean
```arc
let flag: bool = true
let enabled: bool = false
```

## Literals, null pointer
```arc
let ptr: *int32 = null

if ptr == null {
    // handle null case
}
```

## Literals, character
```arc
let ch: char = 'a'
let digit: char = '5'
```

## Literals, character escapes
```arc
let newline: char = '\n'
let tab: char = '\t'
let backslash: char = '\\'
let quote: char = '\''
let null_char: char = '\0'
```

## Literals, string
```arc
let msg: string = "Hello, World!"
let empty: string = ""
```

## Literals, string escapes
```arc
let msg: string = "Hello\nWorld"
let path: string = "C:\\Users\\file"
let quote: string = "He said \"hello\""
let tab: string = "Column1\tColumn2"
```

## Pointers, Basic
```arc
let ptr: *int32 = &value
```

## Pointers, Void (Opaque)
```arc
// Generic pointer to unknown memory (equivalent to C void*)
let handle: *void = malloc(64)
```

## References, Basic
```arc
let ref: &int32 = value
```

## Collections, vector
```arc
let v: vector<int32> = {}
```

## Collections, vector initialization
```arc
// Empty vector
let empty: vector<int32> = {}

// Initialized with values
let nums: vector<int32> = {1, 2, 3, 4, 5}

// Type inference
let items = {10, 20, 30}
```

## Collections, map
```arc
let m: map<string, int32> = {}
```

## Collections, map initialization
```arc
// Empty map
let empty: map<string, int32> = {}

// Initialized with key-value pairs
let scores: map<string, int32> = {"alice": 100, "bob": 95}

// Type inference
let config = {"host": "localhost", "port": "8080"}
```

## Functions, basic
```arc
func add(a: int32, b: int32) int32 {
    return a + b
}
```

## Functions, no return
```arc
func print(msg: string) {
    
}
```

## Functions, async
```arc
// Async function that returns a value
async func fetch_data(url: string) string {
    let response = await http.get(url)
    return response.body
}

// Async function with no return
async func process_items(items: vector<string>) {
    for item in items {
        await process(item)
    }
}
```

## Functions, await usage
```arc
async func main() {
    // Await async function call
    let data = await fetch_data("https://api.example.com")
    
    // Multiple awaits
    let result1 = await task1()
    let result2 = await task2()
    
    // Await in expressions
    if await check_status() {
        // do something
    }
}
```

## Structs, basic (value type - stack allocated, copied)
```arc
struct Point {
    x: int32
    y: int32
}
```

## Structs, initialization
```arc
// Named field initialization
let p1: Point = Point{x: 10, y: 20}

// Qualified struct initialization (Namespace.Type)
let client = net.Socket{fd: -1, connected: false}

// Type inference
let p2 = Point{x: 5, y: 15}

// Default/zero initialization
let p3: Point = Point{}
```

## Structs, field access
```arc
let p: Point = Point{x: 10, y: 20}
let x = p.x
p.y = 30
```

## Structs, inline methods
```arc
struct Point {
    x: int32
    y: int32
    
    func distance(self p: Point) float64 {
        return cast<float64>(p.x * p.x + p.y * p.y)
    }
    
    func move(self p: *Point, dx: int32, dy: int32) {
        p.x += dx
        p.y += dy
    }
}
```

## Structs, mutating methods
```arc
struct Counter {
    count: int32
    
    // Mutating method - modifies the struct in-place
    // Only callable on mutable instances (let, not const)
    mutating increment(self c: *Counter) {
        c.count++
    }
    
    mutating add(self c: *Counter, value: int32) {
        c.count += value
    }
    
    // Non-mutating method - read-only access
    func get_count(self c: Counter) int32 {
        return c.count
    }
}

// Usage
let counter = Counter{count: 0}  // Mutable
counter.increment()               // OK - counter is mutable
counter.add(5)                    // OK

const frozen = Counter{count: 10} // Immutable
// frozen.increment()              // ERROR - cannot call mutating method on const
let value = frozen.get_count()    // OK - non-mutating method
```

## Structs, flat methods (alternative style)
```arc
struct Point {
    x: int32
    y: int32
}

// Methods can be declared outside the struct body
func distance(self p: Point) float64 {
    return cast<float64>(p.x * p.x + p.y * p.y)
}

func move(self p: *Point, dx: int32, dy: int32) {
    p.x += dx
    p.y += dy
}

// Mutating method declared outside
mutating reset(self p: *Point) {
    p.x = 0
    p.y = 0
}
```

## Classes, basic (reference type - heap allocated, ref counted)
```arc
class Client {
    name: string
    port: int32
}
```

## Classes, inline methods
```arc
class Client {
    name: string
    port: int32
    
    func connect(self c: *Client, host: string) bool {
        return true
    }
    
    // Async method
    async func fetch_data(self c: *Client) string {
        let response = await http.get("https://example.com")
        return response.body
    }
    
    deinit(self c: *Client) {
        // cleanup when ref count hits 0
    }
}
```

## Classes, flat methods (alternative style)
```arc
class Client {
    name: string
    port: int32
}

// Methods can be declared outside the class body
func connect(self c: *Client, host: string) bool {
    return true
}

async func fetch_data(self c: *Client) string {
    let response = await http.get("https://example.com")
    return response.body
}

deinit(self c: *Client) {
    // cleanup when ref count hits 0
}
```

## Methods, self pattern declaration
```arc
struct Client {
    port: int32
}

// 'self' keyword allows dot notation on the instance
// Colon used for consistency: self name: type
func Connect(self c: *Client, host: string) bool {
    return true
}
```

## Methods, usage
```arc
let c: Client = Client{port: 8080}
// Method call using dot notation
c.Connect("localhost")

// Async method call
async func example() {
    let data = await c.fetch_data()
}
```

## Type Differences

**class vs struct:**
- `class` = Reference type (heap allocated, ref counted, shared via pointers)
- `struct` = Value type (stack allocated, copied on assignment, no ref counting)
- Both support methods (inline or flat declaration style)
- Only `class` supports `deinit` (called when ref count reaches 0)
- `struct` methods can be marked `mutating` to modify the struct in-place
- `mutating` keyword only applies to structs (classes already use pointers)

## Control Flow, if-else
```arc
if condition {
    
} else if condition2 {
    
} else {
    
}
```

## Control Flow, for loop (C-style)
```arc
// Standard three-part for loop
for let i = 0; i < 10; i = i + 1 {
    // loop body
}

// With increment operator
for let i = 0; i < 10; i++ {
    // loop body
}
```

## Control Flow, for loop (while-style)
```arc
// Condition-only for loop (like while)
let j = 5
for j > 0 {
    j--
}
```

## Control Flow, for loop (infinite)
```arc
// Infinite loop with break
let counter = 0
for {
    counter++
    
    if counter >= 10 {
        break
    }
}
```

## Control Flow, for-in loop (iterators)
```arc
// Iterate over vector
let items: vector<int32> = {1, 2, 3, 4, 5}
for item in items {
    // use item
}

// Iterate over map (key, value)
let scores: map<string, int32> = {"alice": 100, "bob": 95}
for key, value in scores {
    // use key and value
}

// Iterate over range
for i in 0..10 {
    // i goes from 0 to 9
}

// Iterate over custom iterable (implements iterator interface)
for chunk in custom_iterator {
    // use chunk
}
```

## Control Flow, break
```arc
for let i = 0; i < 10; i++ {
    if i == 5 {
        break  // Exit loop
    }
}
```

## Control Flow, continue
```arc
for let i = 0; i < 10; i++ {
    if i == 5 {
        continue  // Skip this iteration
    }
}
```

## Control Flow, defer
```arc
let ptr = malloc(64)
// Execution is guaranteed at scope exit (LIFO order)
defer free(ptr)
```

## Control Flow, return
```arc
return value
```

## Operators, arithmetic
```arc
let sum = a + b
let diff = a - b
let prod = a * b
let quot = a / b
let rem = a % b
```

## Operators, bitwise
```arc
let b_or = a | b    // Bitwise OR
let b_xor = a ^ b   // Bitwise XOR
let b_and = a & b   // Bitwise AND
let shl = a << 2    // Left Shift
let shr = a >> 1    // Right Shift (Arithmetic for signed, Logical for unsigned)
let b_not = ~a      // Bitwise NOT
```

## Operators, compound assignment
```arc
x += 5   // x = x + 5
x -= 3   // x = x - 3
x *= 2   // x = x * 2
x /= 4   // x = x / 4
x %= 3   // x = x % 3
```

## Operators, increment/decrement
```arc
// Post-increment (returns old value, then increments)
i++
pos++

// Pre-increment (increments, then returns new value)
++i
++pos

// Post-decrement (returns old value, then decrements)
i--
pos--

// Pre-decrement (decrements, then returns new value)
--i
--pos
```

## Operators, pointer arithmetic
```arc
let ptr: *int32 = ...

// Advances pointer by 1 * sizeof(int32) bytes
let next = ptr + 1 

// Moves back by 2 * sizeof(int32) bytes
let prev = ptr - 2
```

## Operators, comparison
```arc
let eq = a == b
let ne = a != b
let lt = a < b
let le = a <= b
let gt = a > b
let ge = a >= b
```

## Operators, logical
```arc
let and = a && b
let or = a || b
```

## Operators, unary
```arc
let neg = -value
let not = !flag
```

## Operators, address-of and dereference
```arc
// Get address of a variable (address-of operator)
let ptr: *int32 = &value

// Dereference pointer to read value
let x = *ptr

// Dereference pointer to write value
*ptr = 42
```

## Intrinsics, stack allocation (alloca)
```arc
// Allocate single item on stack (returns *T)
// Essential for creating kernel buffers without malloc
let ptr = alloca(int32)

// Allocate array/buffer on stack (returns *T, second arg is count)
let buffer = alloca(byte, 1024)
```

## Intrinsics, meta (sizeof, alignof)
```arc
// Compile-time size in bytes (returns usize)
// Essential for passing buffer sizes to read/write/mmap
let sz = sizeof<int32>       // 4
let st_sz = sizeof<Stat>     // Struct size (with padding)

// Compile-time alignment requirement (returns usize)
let align = alignof<float64> // 8
```

## Intrinsics, memory block (memset, memcpy, memmove)
```arc
let buf = alloca(byte, 1024)

// memset(dest: *void, val: byte, count: usize)
// Essential for zeroing stack structs before syscalls
memset(buf, 0, 1024)

// memcpy(dest: *void, src: *void, count: usize)
// Essential for copying data to/from kernel buffers
// UNSAFE if regions overlap
memcpy(dest_ptr, src_ptr, 1024)

// memmove(dest: *void, src: *void, count: usize)
// Like memcpy but handles overlapping regions safely
// Essential for buffer manipulation where source/dest might overlap
memmove(dest_ptr, src_ptr, 1024)
```

## Intrinsics, string operations (strlen, memchr)
```arc
// strlen(str: *byte) -> usize
// Calculate C-string length (stops at null terminator)
// Essential when interfacing with syscalls expecting null-terminated strings
let cstr: *byte = "hello\0"
let len = strlen(cstr)  // 5

// memchr(ptr: *void, val: byte, count: usize) -> *void
// Find first occurrence of byte in memory region
// Returns pointer to found byte, or null if not found
// Essential for parsing, finding newlines in buffers, etc.
let buf: *byte = "hello\nworld"
let newline = memchr(buf, '\n', 11)  // Points to '\n'
```

## Intrinsics, variadic (va_start, va_arg, va_end)
```arc
// Essential for implementing printf without libc
func printf(fmt: string, ...) {
    // Initialize argument walker using the last known arg
    let args = va_start(fmt)
    defer va_end(args)

    // Retrieve next argument as specific type
    let val = va_arg<int32>(args)
}
```

## Intrinsics, process (raise)
```arc
// Immediately abort execution with a message
// Internally calls SYS_EXIT or triggers SIGABRT
// Used for unrecoverable errors/crashes.
if ptr == null {
    raise("Memory corrupted")
}
```

### Intrinsics, memory compare (memcmp)
```arc
// memcmp(ptr1: *void, ptr2: *void, count: usize) -> int32
// Returns 0 if equal, <0 or >0 if different.
// Essential for implementing '==' for strings and structs.
let diff = memcmp(ptr1, ptr2, 1024)
```

### Intrinsics, bit manipulation
```arc
// bit_cast<TargetType>(SourceValue)
// Reinterprets the raw bits without conversion.
// Example: getting the IEEE754 integer representation of a float.
let f: float32 = 1.0
let bits = bit_cast<uint32>(f) // 0x3F800000, not 1
```

## Intrinsics, syscall
```arc
let msg = "Hello, Direct Syscall!\n"
let len = 23

// syscall(number, arg1, arg2, arg3, arg4, arg5, arg6)
let result = syscall(SYS_WRITE, STDOUT, msg, len)
```

## Memory, load (dereference to read)
```arc
let value = *ptr
```

## Memory, store (dereference to write)
```arc
*ptr = value
```

## Memory, indexed pointer access
```arc
let buffer = alloca(byte, 1024)

// Read byte at offset 5
let byte_val = buffer[5]

// Write byte at offset 10
buffer[10] = 0x42

// Works with any pointer type
let ptr: *int32 = array_base
let third_element = ptr[2]  // Read array[2]
ptr[3] = 100                // Write array[3] = 100
```

## Type Casting, basic
```arc
let result = cast<int64>(value)
```

## Type Casting, pointer conversions
```arc
// Cast between pointer types
let byte_ptr = cast<*byte>(int_ptr)

// Cast pointer to integer (for address arithmetic)
let addr = cast<uint64>(ptr)

// Cast integer to pointer
let new_ptr = cast<*int32>(addr)

// Cast to void pointer (generic)
let generic = cast<*void>(typed_ptr)
```

## Function Calls
```arc
let result = add(5, 10)
```

## Extern, C interoperability
```arc
extern libc {
    // Maps Arc 'printf' to C symbol 'printf'
    // Uses *byte (C-String) instead of high-level string
    func printf "printf" (*byte, ...) int32
    
    // Maps Arc 'sleep' to C symbol 'usleep'
    func sleep "usleep" (int32) int32

    // Direct mapping (no alias needed if names match)
    func usleep(int32) int32
}

libc.printf()
```

## Enums
```arc
// Basic enumeration
enum Status {
    OK
    ERROR
    PENDING
}

// With explicit values
enum HttpCode {
    OK = 200
    NOT_FOUND = 404
    SERVER_ERROR = 500
}

// Underlying type (default is int32)
enum Color: uint8 {
    RED = 0xFF0000
    GREEN = 0x00FF00
    BLUE = 0x0000FF
}
```


## Control Flow, try-except
```arc
// Basic try-except block
try {
    let result = divide(10, 0)
    io.printf("Result: %d\n", result)
} except err {
    io.printf("Error: %s\n", err)
}

// Try-except with typed errors
enum FileError {
    NotFound
    PermissionDenied
    IOError
}

func read_file(path: string) string throws FileError {
    if !exists(path) {
        throw FileError.NotFound
    }
    return contents
}

try {
    let data = read_file("/tmp/config.txt")
    process(data)
} except FileError.NotFound {
    io.printf("File not found\n")
} except FileError.PermissionDenied {
    io.printf("Permission denied\n")
} except err {
    // Catch-all for other errors
    io.printf("Unexpected error\n")
}

// Try-except with finally (always executes)
try {
    let file = open("data.txt")
    process(file)
} except err {
    io.printf("Error: %s\n", err)
} finally {
    // Cleanup code always runs
    cleanup()
}

// Try-except with ref-counted classes (automatic cleanup)
try {
    let client = HttpClient{}  // Ref count = 1
    let data = client.fetch("https://api.example.com")
    process(data)
} except err {
    // If exception thrown, client is automatically released
    io.printf("Failed to fetch: %s\n", err)
}
// client ref count decremented here (or during exception unwinding)
```


## Functions Return Tuples
```arc
func divide(a: int32, b: int32) -> (int32, bool) {
    if b == 0 {
        return (0, false)
    }
    return (a / b, true)
}

let (result, ok) = divide(10, 2)
```

