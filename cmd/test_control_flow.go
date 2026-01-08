package main

func init() {
	// Phase 3: Control Flow
	AllTests = append(AllTests,
		TestCase{
			Name: "control_if_basic",
			Body: `
    let x = 10
    if x > 5 {
        printf("condition_true\n")
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
        printf("greater\n")
    } else {
        printf("not_greater\n")
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
        printf("A\n")
    } else if score >= 80 {
        printf("B\n")
    } else if score >= 70 {
        printf("C\n")
    } else {
        printf("F\n")
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
            printf("both_true\n")
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
    printf("sum=%d\n", sum)
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
    printf("count=%d\n", count)
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
    printf("i=%d\n", i)
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
    printf("sum_odd=%d\n", sum)
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
    printf("product=%d\n", product)
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
            printf("one\n")
        case 2:
            printf("two\n")
        case 3:
            printf("three\n")
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
            printf("one\n")
        case 2:
            printf("two\n")
        default:
            printf("other\n")
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
            printf("low\n")
        case 4, 5, 6:
            printf("mid\n")
        default:
            printf("high\n")
    }
    `,
			Expected: "low",
			Phase: 3,
		},
		TestCase{
			Name:    "control_return_void",
			Globals: `
    func test_return() {
        printf("before_return\n")
        return
    }
    `,
			Body: `
    test_return()
    printf("done\n")
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
    printf("value=%d\n", x)
    `,
			Expected: "value=42",
			Phase: 3,
		},
		TestCase{
			Name: "control_defer_single",
			Body: `
    printf("before\n")
    defer printf("deferred\n")
    printf("after\n")
    `,
			Expected: "before\nafter\ndeferred",
			Phase: 3,
		},
		TestCase{
			Name: "control_defer_multiple",
			Body: `
    defer printf("first\n")
    defer printf("second\n")
    defer printf("third\n")
    printf("main\n")
    `,
			Expected: "main\nthird\nsecond\nfirst",
			Phase: 3,
		},
		TestCase{
			Name:    "control_defer_function",
			Globals: `
    func cleanup() {
        printf("cleanup_called\n")
    }
    `,
			Body: `
    printf("start\n")
    defer cleanup()
    printf("end\n")
    `,
			Expected: "start\nend\ncleanup_called",
			Phase: 3,
		},
	)
}