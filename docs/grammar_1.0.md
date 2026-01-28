# arc Language Grammar (Version 1.2 - Core Syntax)

Grammar rules:
 * type declaration can only have dot.

Grammar to not add to the parser files:
    * Empty initializer

# Empty initializer
You usually disallow "empty" initializer lists as statements. An array literal {1, 2, 3} on its own line does nothing anyway.
```arc
{ 
    print("data") 
}
```

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

const BUFFER_SIZE: usize = 1024
let buffer: array<byte, BUFFER_SIZE> = {}
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

## Functions, async gpu
```arc
async func process_gpu_memory<gpu>(gpu_arr: *float32, n: usize, gpu_result: *float32) {
    let idx = gpu.thread_id()
}
```

## Functions, async gpu await usage
```arc
// Await with gpu device index 1
let result = await(1) process_gpu_memory()

// Await with auto gpu
let result = await process_gpu_memory()
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

## Functions, async callback
```arc
// Async callback with single param
onClick(args, async (item: string) => {
    await process(item)
})

// Async callback with multiple params
some.fetch(args, async (url: string, timeout: int32) => {
    let resp = await http.get(url, timeout)
    return resp.body
})

// Async callback with no params
button.on_click(async () => {
    await save_state()
})
```

## Function Return Tuples
```arc
func divide(a: int32, b: int32) (int32, bool) {
    if b == 0 {
        return (0, false)
    }
    return (a / b, true)
}

let (result, ok) = divide(10, 2)
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

## Control Flow, switch
```arc
let status = 2

switch status {
    case 0:
        io.print("OK")
    case 1:
        io.print("Pending")
    case 2:
        io.print("Error")
    default:
        io.print("Unknown")
}

switch status {
    // Runs if status is 1 OR 2 OR 3
    case 1, 2, 3:
        io.print("Active/Pending")

    // Runs only if status is 4
    case 4:
        io.print("Completed")

    default:
        io.print("Unknown")
}
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
let buffer = alloca<byte>(1024)

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
extern c {
    // Maps Arc 'printf' to C symbol 'printf'
    // Uses *byte (C-String) instead of high-level string
    func printf "printf" (*byte, ...) int32
    
    // Maps Arc 'sleep' to C symbol 'usleep'
    func sleep "usleep" (int32) int32

    // Direct mapping (no alias needed if names match)
    func usleep(int32) int32
}

printf()
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

## Generics, struct, monomorphizes
```arc
struct Box<T> {
    value: T
    
    func get(self b: Box<T>) T {
        return b.value
    }
    
    func set(self b: *Box<T>, val: T) {
        b.value = val
    }
}
```

## Generics, multiple type parameters, monomorphizes
```arc
struct Pair<K, V> {
    key: K
    value: V
}

struct Result<T, E> {
    data: T
    error: E
    success: bool
}
```

## Generics, functions, monomorphizes
```arc
func swap<T>(a: *T, b: *T) {
    let tmp: T = *a
    *a = *b
    *b = tmp
}

func find<T>(arr: *T, len: usize, val: T) isize {
    for let i: usize = 0; i < len; i++ {
        if arr[i] == val {
            return cast<isize>(i)
        }
    }
    return -1
}
```

## Execution Context, process
```arc
// With args and return
let handle = process func(x: int) { 
    work(x) 
}(1000)

// Without args and return
process func() {  
    work(x)
}()
```

## Execution Context, async
```arc
// With args and return
async func(x: int) { 
    work(x) 
}(1000)

// Without args and return
async func() {  
    work(x)
}()
```

## Async Event Handlers (property assignment)
```arc
// Async handler - ergonomic shorthand, no await keyword allowed
handler.onEvent = (data: EventData) => { 
    process_immediate(data)
    update_state()
}

// Another example
stream.onData = (chunk: bytes) => {
    buffer.append(chunk)
    validate(chunk)
}
```

## Async Event Handlers (with await capability)
```arc
// Async handler with await - explicit async keyword required
handler.onEvent = async (data: EventData) => { 
    let result = await process_async(data)
    fmt.print(result.status)
    await store.save(result)
}

// Multiple awaits allowed
stream.onData = async (chunk: bytes) => {
    let validated = await validate_async(chunk)
    await buffer.write(validated)
    await notify_listeners(validated)
}
```

## Async Method Calls (callback parameter)
```arc
// Async callback - ergonomic shorthand, no await keyword allowed
service.request(args, (result: Result) => {
    fmt.print("Request completed")
    handle_result(result)
})

// Multiple arguments before callback
network.fetch(url, timeout, (response: Response) => {
    fmt.printf("Status: %d\n", response.status)
})
```

## Async Method Calls (callback with await capability)
```arc
// Async callback with await - explicit async keyword required
service.request(args, async (result: Result) => {
    let processed = await transform(result)
    fmt.print(processed.data)
})

// Async callback with multiple params
router.handle("/api/data", async (req: Request, res: Response) => {
    let data = await db.query("SELECT * FROM records")
    await res.send_json(data)
})
```

**Note:** Both forms are async (run on smart threads). The `async` keyword only determines whether `await` is allowed inside the lambda body. Omitting `async` is ergonomic shorthand for callbacks that don't need to suspend.


## Functions, async callback (indirect invocation)

```arc
// Define a simple event system with callback
class Event {
    // Store the async callback as a function pointer
    onTrigger: async func(string) void
    
    // Method to register the callback
    func register(self evt: *Event, handler: async func(string) void) {
        evt.onTrigger = handler
    }
    
    // Method that invokes the callback indirectly
    func send(self evt: *Event, message: string) {
        // Indirect call through function pointer
        // Returns Future<void>, but we fire-and-forget
        evt.onTrigger(message)
    }
}

// Usage
let evt = Event{}

// Register async callback
evt.register(async (msg: string) => {
    let processed = await process_message(msg)
    fmt.printf("Processed: %s\n", processed)
})

// Trigger the event - invokes callback indirectly
evt.send("Hello, World!")
```

## Functions, async callback (property assignment style)

```arc
class TcpServer {
    port: int32
    // Function pointer field for the callback
    onReceive: async func(bytes) void
    
    // Internal method that invokes the callback
    func handle_data(self s: *TcpServer, data: bytes) {
        // Check if callback is set
        if s.onReceive != null {
            // INDIRECT CALL - invoke whatever function was assigned
            s.onReceive(data)
        }
    }
}

// Usage
let server = TcpServer{port: 8080}

// Assign async callback directly to property
server.onReceive = async (data: bytes) => {
    let decoded = await decode_packet(data)
    await store_in_db(decoded)
}

// Server invokes it indirectly when data arrives
server.handle_data(received_bytes)
```

## Functions, sync callback (indirect invocation)

```arc
class Button {
    // Sync callback (no await allowed inside)
    onClick: func(int32, int32) void
    
    // Invoke the callback when button is pressed
    func press(self b: *Button, x: int32, y: int32) {
        if b.onClick != null {
            // INDIRECT CALL
            b.onClick(x, y)
        }
    }
}

// Usage
let button = Button{}

button.onClick = (x: int32, y: int32) => {
    fmt.printf("Clicked at (%d, %d)\n", x, y)
}

button.press(100, 200)  // Triggers the callback
```