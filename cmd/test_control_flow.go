package main

func init() {
	// Phase 3: Control Flow
	AllTests = append(AllTests,
		TestCase{
			Name: "control_if_basic",
			Body: `
    let x = 10
    if x > 5 {
        libc.printf("condition_true\n")
    }
    `,
			Expected: "condition_true",
			Phase: 3,
		},
		TestCase{
			Name: "control_if_else",
			Body: `
    let x = 3
    if x > 5 {
        libc.printf("greater\n")
    } else {
        libc.printf("not_greater\n")
    }
    `,
			Expected: "not_greater",
			Phase: 3,
		},
		TestCase{
			Name: "control_if_elseif",
			Body: `
    let score = 75
    if score >= 90 {
        libc.printf("A\n")
    } else if score >= 80 {
        libc.printf("B\n")
    } else if score >= 70 {
        libc.printf("C\n")
    } else {
        libc.printf("F\n")
    }
    `,
			Expected: "C",
			Phase: 3,
		},
		TestCase{
			Name: "control_if_nested",
			Body: `
    let x = 10
    let y = 20
    if x > 5 {
        if y > 15 {
            libc.printf("both_true\n")
        }
    }
    `,
			Expected: "both_true",
			Phase: 3,
		},
		TestCase{
			Name: "control_for_cstyle",
			Body: `
    let sum = 0
    for let i = 0; i < 5; i++ {
        sum += i
    }
    libc.printf("sum=%d\n", sum)
    `,
			Expected: "sum=10",
			Phase: 3,
		},
		TestCase{
			Name: "control_for_while",
			Body: `
    let count = 0
    for count < 3 {
        count++
    }
    libc.printf("count=%d\n", count)
    `,
			Expected: "count=3",
			Phase: 3,
		},
		TestCase{
			Name: "control_for_infinite_break",
			Body: `
    let i = 0
    for {
        i++
        if i >= 5 {
            break
        }
    }
    libc.printf("i=%d\n", i)
    `,
			Expected: "i=5",
			Phase: 3,
		},
		TestCase{
			Name: "control_for_continue",
			Body: `
    let sum = 0
    for let i = 0; i < 10; i++ {
        if i % 2 == 0 {
            continue
        }
        sum += i
    }
    libc.printf("sum_odd=%d\n", sum)
    `,
			Expected: "sum_odd=25",
			Phase: 3,
		},
		TestCase{
			Name: "control_for_nested",
			Body: `
    let product = 0
    for let i = 1; i <= 3; i++ {
        for let j = 1; j <= 2; j++ {
            product = i * j
        }
    }
    libc.printf("product=%d\n", product)
    `,
			Expected: "product=6",
			Phase: 3,
		},
		TestCase{
			Name: "control_switch_simple",
			Body: `
    let x = 2
    switch x {
        case 1:
            libc.printf("one\n")
        case 2:
            libc.printf("two\n")
        case 3:
            libc.printf("three\n")
    }
    `,
			Expected: "two",
			Phase: 3,
		},
		TestCase{
			Name: "control_switch_default",
			Body: `
    let x = 99
    switch x {
        case 1:
            libc.printf("one\n")
        case 2:
            libc.printf("two\n")
        default:
            libc.printf("other\n")
    }
    `,
			Expected: "other",
			Phase: 3,
		},
		TestCase{
			Name: "control_switch_multi",
			Body: `
    let x = 3
    switch x {
        case 1, 2, 3:
            libc.printf("low\n")
        case 4, 5, 6:
            libc.printf("mid\n")
        default:
            libc.printf("high\n")
    }
    `,
			Expected: "low",
			Phase: 3,
		},
		TestCase{
			Name:    "control_return_void",
			Globals: `
    func test_return() {
        libc.printf("before_return\n")
        return
    }
    `,
			Body: `
    test_return()
    libc.printf("done\n")
    `,
			Expected: "before_return\ndone",
			Phase: 3,
		},
		TestCase{
			Name:    "control_return_value",
			Globals: `
    func get_value() int32 {
        return 42
    }
    `,
			Body: `
    let x = get_value()
    libc.printf("value=%d\n", x)
    `,
			Expected: "value=42",
			Phase: 3,
		},
		TestCase{
			Name: "control_defer_single",
			Body: `
    libc.printf("before\n")
    defer libc.printf("deferred\n")
    libc.printf("after\n")
    `,
			Expected: "before\nafter\ndeferred",
			Phase: 3,
		},
		TestCase{
			Name: "control_defer_multiple",
			Body: `
    defer libc.printf("first\n")
    defer libc.printf("second\n")
    defer libc.printf("third\n")
    libc.printf("main\n")
    `,
			Expected: "main\nthird\nsecond\nfirst",
			Phase: 3,
		},
		TestCase{
			Name:    "control_defer_function",
			Globals: `
    func cleanup() {
        libc.printf("cleanup_called\n")
    }
    `,
			Body: `
    libc.printf("start\n")
    defer cleanup()
    libc.printf("end\n")
    `,
			Expected: "start\nend\ncleanup_called",
			Phase: 3,
		},
	)
}