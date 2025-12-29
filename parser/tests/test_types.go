package main

func init() {
	// Basic type tests
	RegisterTest("Signed Integers", func() string {
		return `let i: int32 = -500`
	})

	RegisterTest("Unsigned Integers", func() string {
		return `let u: uint64 = 10000`
	})

	RegisterTest("USize Type", func() string {
		return `let len: usize = 100`
	})

	RegisterTest("ISize Type", func() string {
		return `let offset: isize = -4`
	})

	RegisterTest("Float32 Type", func() string {
		return `let f32: float32 = 3.14`
	})

	RegisterTest("Float64 Type", func() string {
		return `let f64: float64 = 2.71828`
	})

	RegisterTest("Byte Type", func() string {
		return `let b: byte = 255`
	})

	RegisterTest("Bool Type", func() string {
		return `let flag: bool = true`
	})

	RegisterTest("Char Type", func() string {
		return `let r: char = 'a'`
	})

	RegisterTest("String Type", func() string {
		return `let s: string = "hello"`
	})

	// Null tests
	RegisterTest("Null Literal", func() string {
		return `let ptr: *int32 = null`
	})

	RegisterTest("Null Check", func() string {
		return `func test() {
    if ptr == null {
        x = 1
    }
}`
	})

	// Pointer and reference tests
	RegisterTest("Pointer Basic", func() string {
		return `let ptr: *int32 = &value`
	})

	RegisterTest("Pointer Void", func() string {
		return `let handle: *void = alloca(void, 64)`
	})

	RegisterTest("Reference Basic", func() string {
		return `let ref: &int32 = value`
	})
}