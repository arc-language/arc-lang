package main

func init() {
	AllTests = append(AllTests, []TestCase{
		{
			Name:    "Variables & Assignments",
			Globals: "",
			Body: `
    let x: int32 = 10
    let y = 20
    let z = x
    
    x = 100
    
    libc.printf("x: %d, y: %d, z: %d\n", x, y, z)
`,
			Expected: "x: 100, y: 20, z: 10",
		},
		{
			Name:    "Integer Arithmetic",
			Globals: "",
			Body: `
    let a = 10
    let b = 3
    
    let sum = a + b
    let sub = a - b
    let mul = a * b
    let div = a / b
    let mod = a % b
    
    libc.printf("%d, %d, %d, %d, %d\n", sum, sub, mul, div, mod)
`,
			Expected: "13, 7, 30, 3, 1",
		},
		{
			Name:    "Float Casting & Logic",
			Globals: "",
			Body: `
    let pi: float64 = 3.14159
    let truncated = cast<int32>(pi)
    
    let large: int64 = 123456789
    let small = cast<int32>(large)
    
    libc.printf("Int: %d, Cast: %d\n", truncated, small)
`,
			Expected: "Int: 3, Cast: 123456789",
		},
		{
			Name:    "Control Flow: If/Else",
			Globals: "",
			Body: `
    let x = 50
    if x > 100 {
        libc.printf("Big")
    } else if x == 50 {
        libc.printf("Exact")
    } else {
        libc.printf("Small")
    }
    
    // Test nested if
    if x > 0 {
        if x < 60 {
            libc.printf(" Range OK")
        }
    }
`,
			Expected: "Exact Range OK",
		},
		{
			Name:    "Control Flow: While Loop",
			Globals: "",
			Body: `
    let i = 0
    let sum = 0
    
    // Simulate while loop using standard for syntax
    for i < 5 {
        sum = sum + i
        i = i + 1
    }
    libc.printf("While Sum: %d\n", sum)
`,
			Expected: "While Sum: 10",
		},
		{
			Name:    "Control Flow: C-Style Loop",
			Globals: "",
			Body: `
    let total = 0
    for let j = 0; j < 5; j = j + 1 {
        total = total + 10
    }
    libc.printf("Loop Total: %d\n", total)
`,
			Expected: "Loop Total: 50",
		},
		{
			Name:    "Control Flow: Break/Continue",
			Globals: "",
			Body: `
    for let i = 0; i < 10; i = i + 1 {
        if i == 2 {
            continue
        }
        if i == 5 {
            break
        }
        libc.printf("%d ", i)
    }
`,
			Expected: "0 1 3 4 ",
		},
		{
			Name:    "Pointers Basic",
			Globals: "",
			Body: `
    let val = 999
    let ptr = &val
    
    libc.printf("Original: %d\n", *ptr)
    
    *ptr = 111
    
    libc.printf("Modified: %d\n", val)
`,
			Expected: "Original: 999\nModified: 111",
		},
		{
			Name:    "Logical Operators",
			Globals: "",
			Body: `
    let t = true
    let f = false
    
    if t && t { libc.printf("AND ") }
    if f || t { libc.printf("OR ") }
    if !f { libc.printf("NOT") }
`,
			Expected: "AND OR NOT",
		},
	}...)
}