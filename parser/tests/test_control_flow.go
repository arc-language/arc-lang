package main

func init() {
	// If statements
	RegisterTest("If Statement", func() string {
		return `func test() {
    if condition {
        x = 1
    }
}`
	})

	RegisterTest("If-Else Statement", func() string {
		return `func test() {
    if condition {
        x = 1
    } else {
        x = 2
    }
}`
	})

	RegisterTest("If-Else-If Statement", func() string {
		return `func test() {
    if condition {
        x = 1
    } else if condition2 {
        x = 2
    } else {
        x = 3
    }
}`
	})

	// For loops
	RegisterTest("For Loop C-Style", func() string {
		return `func test() {
    for let i = 0; i < 10; i = i + 1 {
        x = i
    }
}`
	})

	RegisterTest("For Loop C-Style Increment", func() string {
		return `func test() {
    for let i = 0; i < 10; i++ {
        x = i
    }
}`
	})

	RegisterTest("For Loop Condition", func() string {
		return `func test() {
    let j = 5
    for j > 0 {
        j--
    }
}`
	})

	RegisterTest("For Loop Infinite", func() string {
		return `func test() {
    let counter = 0
    for {
        counter++
        if counter >= 10 {
            break
        }
    }
}`
	})

	RegisterTest("For-In Loop Vector", func() string {
		return `func test() {
    let items: vector<int32> = {1, 2, 3, 4, 5}
    for item in items {
        x = item
    }
}`
	})

	RegisterTest("For-In Loop Map", func() string {
		return `func test() {
    let scores: map<string, int32> = {"alice": 100, "bob": 95}
    for key, value in scores {
        x = value
    }
}`
	})

	RegisterTest("For-In Loop Range", func() string {
		return `func test() {
    for i in 0..10 {
        x = i
    }
}`
	})

	// Control flow statements
	RegisterTest("Break Statement", func() string {
		return `func test() {
    for let i = 0; i < 10; i = i + 1 {
        if i == 5 {
            break
        }
    }
}`
	})

	RegisterTest("Continue Statement", func() string {
		return `func test() {
    for let i = 0; i < 10; i = i + 1 {
        if i == 5 {
            continue
        }
        x = i
    }
}`
	})

	RegisterTest("Defer Statement", func() string {
		return `func test() {
    let ptr = alloca(int32)
    defer free(ptr)
}`
	})

	RegisterTest("Return Statement", func() string {
		return `func test() int32 {
    return 42
}`
	})

	RegisterTest("Return Void", func() string {
		return `func test() {
    return
}`
	})
}