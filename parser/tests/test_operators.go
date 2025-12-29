package main

func init() {
	RegisterTest("Arithmetic Operators", func() string {
		return `func test() {
    let sum = a + b
    let diff = a - b
    let prod = a * b
    let quot = a / b
    let rem = a % b
}`
	})

	RegisterTest("Compound Assignment All", func() string {
		return `func test() {
    x += 5
    x -= 3
    x *= 2
    x /= 4
    x %= 3
}`
	})

	RegisterTest("Increment Decrement Operators", func() string {
		return `func test() {
    i++
    pos++
    i--
    pos--
}`
	})

	RegisterTest("Pre-Increment Decrement", func() string {
		return `func test() {
    let x = ++i
    let y = --j
}`
	})

	RegisterTest("Post-Increment Decrement", func() string {
		return `func test() {
    let x = i++
    let y = j--
}`
	})

	RegisterTest("Pointer Arithmetic", func() string {
		return `func test() {
    let next = ptr + 1
    let prev = ptr - 2
}`
	})

	RegisterTest("Comparison Operators", func() string {
		return `func test() {
    let eq = a == b
    let ne = a != b
    let lt = a < b
    let le = a <= b
    let gt = a > b
    let ge = a >= b
}`
	})

	RegisterTest("Logical Operators", func() string {
		return `func test() {
    let and = a && b
    let or = a || b
}`
	})

	RegisterTest("Unary Operators", func() string {
		return `func test() {
    let neg = -value
    let not = !flag
}`
	})

	RegisterTest("Address-Of Operator", func() string {
		return `let ptr: *int32 = &value`
	})

	RegisterTest("Dereference Operator Read", func() string {
		return `let x = *ptr`
	})

	RegisterTest("Dereference Operator Write", func() string {
		return `func test() {
    *ptr = 42
}`
	})

	RegisterTest("Range Expression", func() string {
		return `func test() {
    let r = 0..10
}`
	})
}