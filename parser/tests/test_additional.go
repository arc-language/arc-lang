package main

func init() {
	// Switch statements
	RegisterTest("Switch Statement", func() string {
		return `func test() {
    let status = 2
    switch status {
        case 0:
            x = 1
        case 1:
            x = 2
        case 2:
            x = 3
        default:
            x = 4
    }
}`
	})

	RegisterTest("Switch With Multiple Cases", func() string {
		return `func test() {
    switch value {
        case 0:
            x = 0
        case 1:
            x = 1
        case 2:
            x = 2
        default:
            x = -1
    }
}`
	})

	// Enum declarations
	RegisterTest("Enum Basic", func() string {
		return `enum Status {
    OK
    ERROR
    PENDING
}`
	})

	RegisterTest("Enum With Explicit Values", func() string {
		return `enum HttpCode {
    OK = 200
    NOT_FOUND = 404
    SERVER_ERROR = 500
}`
	})

	RegisterTest("Enum With Type", func() string {
		return `enum Color: uint8 {
    RED = 0xFF0000
    GREEN = 0x00FF00
    BLUE = 0x0000FF
}`
	})

	// Try-except-finally
	RegisterTest("Try-Except Basic", func() string {
		return `func test() {
    try {
        let result = divide(10, 0)
        x = result
    } except err {
        x = 0
    }
}`
	})

	RegisterTest("Try-Except Specific Error", func() string {
		return `func test() {
    try {
        let data = read_file("/tmp/config.txt")
        process(data)
    } except FileError.NotFound {
        x = 1
    } except FileError.PermissionDenied {
        x = 2
    } except err {
        x = 3
    }
}`
	})

	RegisterTest("Try-Except-Finally", func() string {
		return `func test() {
    try {
        let file = open("data.txt")
        process(file)
    } except err {
        x = 1
    } finally {
        cleanup()
    }
}`
	})

	// Function return tuples
	RegisterTest("Function Return Tuple", func() string {
		return `func divide(a: int32, b: int32) (int32, bool) {
    if b == 0 {
        return (0, false)
    }
    return (a / b, true)
}`
	})

	RegisterTest("Tuple Destructuring", func() string {
		return `func test() {
    let (result, ok) = divide(10, 2)
}`
	})

	RegisterTest("Tuple Assignment", func() string {
		return `func test() {
    let (x, y, z) = get_coordinates()
}`
	})

	// Fixed-size arrays
	RegisterTest("Array Declaration", func() string {
		return `let arr: array<int32, 5> = {1, 2, 3, 4, 5}`
	})

	RegisterTest("Array Float", func() string {
		return `let coords: array<float64, 3> = {1.0, 2.0, 3.0}`
	})

	RegisterTest("Array Zero Init", func() string {
		return `let zeros: array<int32, 100> = {}`
	})

	RegisterTest("Array Type Inference", func() string {
		return `let items: array<_, 3> = {10, 20, 30}`
	})

	RegisterTest("Array Access", func() string {
		return `func test() {
    let arr: array<int32, 5> = {1, 2, 3, 4, 5}
    let val = arr[2]
    arr[3] = 42
}`
	})

	RegisterTest("Array Pointer", func() string {
		return `func test() {
    let arr: array<int32, 5> = {1, 2, 3, 4, 5}
    let ptr: *int32 = &arr[0]
}`
	})

	// String interpolation
	RegisterTest("String Interpolation Basic", func() string {
		return `func test() {
    let name = "Alice"
    let msg = "Hello, \(name)!"
}`
	})

	RegisterTest("String Interpolation Expression", func() string {
		return `func test() {
    let age = 30
    let info = "You are \(age) years old, born in \(2025 - age)"
}`
	})

	RegisterTest("String Interpolation Complex", func() string {
		return `func test() {
    let result = "Sum: \(a + b), Product: \(a * b)"
}`
	})

	RegisterTest("String Interpolation Function Call", func() string {
		return `func test() {
    let upper = "Name: \(name.to_upper())"
}`
	})

	// Generics
	RegisterTest("Generic Struct", func() string {
		return `struct vector<T> {
    data: *T
    len: usize
    cap: usize
    
    func push(self v: *vector<T>, item: T) {
        if v.len >= v.cap {
            v.grow()
        }
        v.data[v.len] = item
        v.len++
    }
    
    func get(self v: *vector<T>, idx: usize) *T {
        return &v.data[idx]
    }
}`
	})

	RegisterTest("Generic Multiple Type Parameters", func() string {
		return `struct Entry<K, V> {
    key: K
    value: V
}

struct map<K, V> {
    buckets: *vector<Entry<K, V>>
    count: usize
    
    func insert(self m: *map<K, V>, key: K, val: V) {
        x = 1
    }
}`
	})

	RegisterTest("Generic Function", func() string {
		return `func swap<T>(a: *T, b: *T) {
    let tmp: T = *a
    *a = *b
    *b = tmp
}`
	})

	RegisterTest("Generic Function Find", func() string {
		return `func find<T>(arr: *T, len: usize, val: T) isize {
    for let i: usize = 0; i < len; i++ {
        if arr[i] == val {
            return cast<isize>(i)
        }
    }
    return -1
}`
	})

	// Slice intrinsic
	RegisterTest("Slice Array", func() string {
		return `func test() {
    let arr: array<int32, 10> = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    let (ptr, len) = slice(arr, 1..4)
}`
	})

	RegisterTest("Slice Vector", func() string {
		return `func test() {
    let vec: vector<byte> = {1, 2, 3, 4, 5}
    let (data_ptr, data_len) = slice(vec, 0..3)
}`
	})

	// Bitwise operators
	RegisterTest("Bitwise OR", func() string {
		return `func test() {
    let b_or = a | b
}`
	})

	RegisterTest("Bitwise XOR", func() string {
		return `func test() {
    let b_xor = a ^ b
}`
	})

	RegisterTest("Bitwise AND", func() string {
		return `func test() {
    let b_and = a & b
}`
	})

	RegisterTest("Bitwise Shift Left", func() string {
		return `func test() {
    let shl = a << 2
}`
	})

	RegisterTest("Bitwise Shift Right", func() string {
		return `func test() {
    let shr = a >> 1
}`
	})

	RegisterTest("Bitwise NOT", func() string {
		return `func test() {
    let b_not = ~a
}`
	})

	RegisterTest("Bitwise All Operators", func() string {
		return `func test() {
    let b_or = a | b
    let b_xor = a ^ b
    let b_and = a & b
    let shl = a << 2
    let shr = a >> 1
    let b_not = ~a
}`
	})

	// Compound bitwise assignment
	RegisterTest("Compound Bitwise Assignment", func() string {
		return `func test() {
    x |= 5
    x &= 3
    x ^= 2
    x <<= 1
    x >>= 1
}`
	})

	// Async callbacks (lambda-style)
	RegisterTest("Async Callback Single Param", func() string {
		return `func test() {
    onClick(args, async (item: string) => {
        await process(item)
    })
}`
	})

	RegisterTest("Async Callback Multiple Params", func() string {
		return `func test() {
    some.fetch(args, async (url: string, timeout: int32) => {
        let resp = await http_get(url, timeout)
        return resp
    })
}`
	})

	RegisterTest("Async Callback No Params", func() string {
		return `func test() {
    button.on_click(async () => {
        await save_state()
    })
}`
	})

	// Qualified types (namespace.Type)
	RegisterTest("Qualified Type Variable", func() string {
		return `let client: net.Socket = Socket{}`
	})

	RegisterTest("Qualified Type Multiple", func() string {
		return `func test() {
    let client: net.Socket = Socket{}
    let config: json.Config = Config{}
}`
	})

	RegisterTest("Qualified Struct Initialization", func() string {
		return `let client = net.Socket{fd: -1, connected: false}`
	})

	// Const with array size
	RegisterTest("Const Buffer Size", func() string {
		return `const BUFFER_SIZE: usize = 1024
let buffer: array<byte, BUFFER_SIZE> = {}`
	})

	// Variadic function parameters
	RegisterTest("Variadic Function", func() string {
		return `func printf(fmt: string, ...) int32 {
    let args = va_start(fmt)
    defer va_end(args)
    let val = va_arg<int32>(args)
    return 0
}`
	})

	// Complete program with everything
	RegisterTest("Complete Program With Switch And Enums", func() string {
		return `namespace main

import "std/io"

enum Status {
    OK = 0
    ERROR = 1
    PENDING = 2
}

func check_status(code: int32) string {
    switch code {
        case 0:
            return "OK"
        case 1:
            return "ERROR"
        case 2:
            return "PENDING"
        default:
            return "UNKNOWN"
    }
}

func main() int32 {
    let status = Status.OK
    let msg = check_status(cast<int32>(status))
    return 0
}`
	})

	RegisterTest("Complete Program With Try-Except", func() string {
		return `namespace main

import "std/io"

func divide(a: int32, b: int32) (int32, bool) {
    if b == 0 {
        return (0, false)
    }
    return (a / b, true)
}

func main() int32 {
    try {
        let (result, ok) = divide(10, 0)
        if !ok {
            raise("Division by zero")
        }
        return result
    } except err {
        return -1
    } finally {
        cleanup()
    }
}`
	})

	RegisterTest("Complete Program With Generics", func() string {
		return `namespace main

import "std/io"

func swap<T>(a: *T, b: *T) {
    let tmp: T = *a
    *a = *b
    *b = tmp
}

func main() int32 {
    let x = 10
    let y = 20
    swap<int32>(&x, &y)
    return 0
}`
	})

	RegisterTest("Complete Program With String Interpolation", func() string {
		return `namespace main

import "std/io"

func main() int32 {
    let name = "Alice"
    let age = 30
    let msg = "Hello, \(name)! You are \(age) years old."
    return 0
}`
	})

	RegisterTest("Complete Program With Arrays", func() string {
		return `namespace main

import "std/io"

func main() int32 {
    let arr: array<int32, 5> = {1, 2, 3, 4, 5}
    let sum = 0
    
    for let i: usize = 0; i < 5; i++ {
        sum = sum + arr[i]
    }
    
    return sum
}`
	})

	RegisterTest("Complete Program With Bitwise", func() string {
		return `namespace main

import "std/io"

func main() int32 {
    let flags = 0
    flags |= 1 << 0
    flags |= 1 << 2
    flags &= ~(1 << 1)
    
    let result = flags ^ 0xFF
    return result
}`
	})
}