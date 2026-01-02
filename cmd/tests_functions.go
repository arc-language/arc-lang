package main

func init() {
	AllTests = append(AllTests, []TestCase{
		{
			Name: "Basic Function Call",
			Globals: `
func add(a: int32, b: int32) int32 {
    return a + b
}

func sub(a: int32, b: int32) int32 {
    return a - b
}
`,
			Body: `
    let sum = add(10, 20)
    let diff = sub(100, 50)
    libc.printf("Sum: %d, Diff: %d\n", sum, diff)
`,
			Expected: "Sum: 30, Diff: 50",
		},
		{
			Name: "Recursion (Fibonacci)",
			Globals: `
func fib(n: int32) int32 {
    if n <= 1 {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}
`,
			Body: `
    // 0, 1, 1, 2, 3, 5, 8
    let f = fib(6)
    libc.printf("Fib(6): %d\n", f)
`,
			Expected: "Fib(6): 8",
		},
		{
			Name: "Pass By Pointer (Swap)",
			Globals: `
func swap(a: *int32, b: *int32) {
    let temp = *a
    *a = *b
    *b = temp
}
`,
			Body: `
    let x = 10
    let y = 20
    libc.printf("Before: %d, %d\n", x, y)
    
    swap(&x, &y)
    
    libc.printf("After: %d, %d\n", x, y)
`,
			Expected: "After: 20, 10",
		},
		{
			Name: "Libc Malloc & Free",
			Globals: "", // libc externs are auto-injected by runner
			Body: `
    // Allocate 8 bytes
    let ptr = libc.malloc(8)
    
    // Cast void* to int64*
    let i_ptr = cast<*int64>(ptr)
    
    // Write value
    *i_ptr = 9999
    
    libc.printf("Heap Value: %d\n", *i_ptr)
    
    // Cleanup
    libc.free(ptr)
`,
			Expected: "Heap Value: 9999",
		},
		{
			Name: "Intrinsic: Sizeof",
			Globals: `
struct Small {
    a: int32
    b: int32
}
struct Padded {
    a: int64
    b: int8
}
`,
			Body: `
    let s32 = sizeof<int32>
    let s64 = sizeof<int64>
    let s_small = sizeof<Small>
    let s_padded = sizeof<Padded>
    
    // Padded should be 16 (8 + 1 + 7 padding)
    libc.printf("%d, %d, %d, %d\n", s32, s64, s_small, s_padded)
`,
			Expected: "4, 8, 8, 16",
		},
		{
			Name: "Intrinsic: Alloca (Stack Array)",
			Globals: "",
			Body: `
    // Allocate array of 10 integers on stack
    let count: usize = 10
    let buffer = alloca(int32, count)
    
    // Write to last element
    // Manual pointer arithmetic: buffer + 9
    let last_ptr = buffer + 9
    *last_ptr = 42
    
    // Read via index syntax
    buffer[0] = 1
    
    libc.printf("First: %d, Last: %d\n", buffer[0], buffer[9])
`,
			Expected: "First: 1, Last: 42",
		},
		{
			Name: "Intrinsic: Memset",
			Globals: "",
			Body: `
    let x: int32 = 123456
    let ptr = cast<*void>(&x)
    
    // Zero out memory
    libc.memset(ptr, 0, 4)
    
    libc.printf("Zeroed: %d\n", x)
`,
			Expected: "Zeroed: 0",
		},
		{
			Name: "Intrinsic: Memcpy",
			Globals: `
struct Point {
    x: int32
    y: int32
}
`,
			Body: `
    let p1 = Point{x: 100, y: 200}
    let p2 = Point{x: 0, y: 0}
    
    let src = cast<*void>(&p1)
    let dst = cast<*void>(&p2)
    
    libc.memcpy(dst, src, sizeof<Point>)
    
    libc.printf("P2: %d, %d\n", p2.x, p2.y)
`,
			Expected: "P2: 100, 200",
		},
	}...)
}