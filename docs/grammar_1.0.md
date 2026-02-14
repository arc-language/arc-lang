# arc Language Grammar (Version 1.3 - Core Syntax)

Grammar rules:
 * type declaration can only have dot
 * no stars in regular arc code, stars only inside extern blocks
 * null valid for class types only, compiler enforces this
 * cast uses type(value) syntax, not cast<T>(value)
 * gpu func for accelerator kernels, target set in build config

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

import yourpackage "some/path/package"

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

## Literals, null
```arc
// null is valid for class types (reference types)
// compiler error if assigned to struct or primitive
let client: Socket = null
let server: net.Server = null

if client == null {
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

## Functions, gpu
```arc
// All params are gpu bound, compiler maps to build target
gpu func kernel(data: float32, n: usize) {
    let idx = thread_id()
    data[idx] = data[idx] * 2.0
}

// Await a gpu func
async func main() {
    let result = await kernel(data, n)
}
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
        return float64(p.x * p.x + p.y * p.y)
    }
    
    func move(self p: Point, dx: int32, dy: int32) {
        p.x += dx
        p.y += dy
    }
}
```

## Structs, flat methods (alternative style)
```arc
struct Point {
    x: int32
    y: int32
}

// Methods can be declared outside the struct body
func distance(self p: Point) float64 {
    return float64(p.x * p.x + p.y * p.y)
}

func move(self p: Point, dx: int32, dy: int32) {
    p.x += dx
    p.y += dy
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
    
    func connect(self c: Client, host: string) bool {
        return true
    }
    
    // Async method
    async func fetch_data(self c: Client) string {
        let response = await http.get("https://example.com")
        return response.body
    }
    
    deinit(self c: Client) {
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
func connect(self c: Client, host: string) bool {
    return true
}

async func fetch_data(self c: Client) string {
    let response = await http.get("https://example.com")
    return response.body
}

deinit(self c: Client) {
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
func connect(self c: Client, host: string) bool {
    return true
}
```

## Methods, usage
```arc
let c = Client{port: 8080}
// Method call using dot notation
c.connect("localhost")

// Async method call
async func example() {
    let data = await c.fetch_data()
}
```

## Type Differences

**class vs struct:**
- `class` = Reference type (heap allocated, ref counted, null allowed)
- `struct` = Value type (stack allocated, copied on assignment, null not allowed)
- Both support methods (inline or flat declaration style)
- Only `class` supports `deinit` (called when ref count reaches 0)

## Type Casting
```arc
// Primitive casts - type(value) syntax
let x = int32(3.14)     // float to int
let y = float64(42)     // int to float
let z = uint8(flags)    // narrow cast
let n = usize(count)    // to pointer-sized int
```

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
let items: vector[int32] = {1, 2, 3, 4, 5}
for item in items {
    // use item
}

// Iterate over map (key, value)
let scores: map[string]int32 = {"alice": 100, "bob": 95}
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

## Function Calls
```arc
let result = add(5, 10)
```

## Extern, C interoperability
```arc
// Stars are required inside extern blocks
// This is the C/C++ boundary, all pointer details live here
extern c {
    func printf(*byte, ...) int32
    func sleep "usleep" (int32) int32
    func usleep(int32) int32
}

printf("hello\n")
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
    
    func set(self b: Box<T>, val: T) {
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
func swap<T>(a: T, b: T) {
    let tmp: T = a
    a = b
    b = tmp
}

func find<T>(arr: vector<T>, val: T) isize {
    for let i: usize = 0; i < arr.len; i++ {
        if arr[i] == val {
            return isize(i)
        }
    }
    return -1
}
```

## Execution Context, process
```arc
// With args and return
let handle = process func(x: int32) { 
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
async func(x: int32) { 
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
class Event {
    onTrigger: async func(string) void
    
    func register(self evt: Event, handler: async func(string) void) {
        evt.onTrigger = handler
    }
    
    func send(self evt: Event, message: string) {
        evt.onTrigger(message)
    }
}

let evt = Event{}

evt.register(async (msg: string) => {
    let processed = await process_message(msg)
    fmt.printf("Processed: %s\n", processed)
})

evt.send("Hello, World!")
```

## Functions, async callback (property assignment style)

```arc
class TcpServer {
    port: int32
    onReceive: async func(bytes) void
    
    func handle_data(self s: TcpServer, data: bytes) {
        if s.onReceive != null {
            s.onReceive(data)
        }
    }
}

let server = TcpServer{port: 8080}

server.onReceive = async (data: bytes) => {
    let decoded = await decode_packet(data)
    await store_in_db(decoded)
}

server.handle_data(received_bytes)
```

## Functions, sync callback (indirect invocation)

```arc
class Button {
    onClick: func(int32, int32) void
    
    func press(self b: Button, x: int32, y: int32) {
        if b.onClick != null {
            b.onClick(x, y)
        }
    }
}

let button = Button{}

button.onClick = (x: int32, y: int32) => {
    fmt.printf("Clicked at (%d, %d)\n", x, y)
}

button.press(100, 200)
```