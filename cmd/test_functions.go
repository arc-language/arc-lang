package main

func init() {
	// Phase 4: Functions
	AllTests = append(AllTests,
		TestCase{
			Name:    "func_no_params_void",
			Globals: `
    func greet() {
        printf("hello\n")
    }
    `,
			Body: `
    greet()
    `,
			Expected: "hello",
			Phase: 4,
		},
		TestCase{
			Name:    "func_with_params_void",
			Globals: `
    func print_sum(a: int32, b: int32) {
        printf("sum=%d\n", a + b)
    }
    `,
			Body: `
    print_sum(10, 20)
    `,
			Expected: "sum=30",
			Phase: 4,
		},
		TestCase{
			Name:    "func_with_return",
			Globals: `
    func add(a: int32, b: int32) int32 {
        return a + b
    }
    `,
			Body: `
    let result = add(15, 25)
    printf("result=%d\n", result)
    `,
			Expected: "result=40",
			Phase: 4,
		},
		TestCase{
			Name:    "func_multiple_return",
			Globals: `
    func divide(a: int32, b: int32) (int32, bool) {
        if b == 0 {
            return (0, false)
        }
        return (a / b, true)
    }
    `,
			Body: `
    let (result, ok) = divide(10, 2)
    printf("result=%d ok=%d\n", result, ok)
    `,
			Expected: "result=5 ok=1",
			Phase: 4,
		},
		TestCase{
			Name:    "func_recursive_factorial",
			Globals: `
    func factorial(n: int32) int32 {
        if n <= 1 {
            return 1
        }
        return n * factorial(n - 1)
    }
    `,
			Body: `
    let result = factorial(5)
    printf("factorial=%d\n", result)
    `,
			Expected: "factorial=120",
			Phase: 4,
		},
		TestCase{
			Name:    "func_recursive_fibonacci",
			Globals: `
    func fib(n: int32) int32 {
        if n <= 1 {
            return n
        }
        return fib(n - 1) + fib(n - 2)
    }
    `,
			Body: `
    let result = fib(7)
    printf("fib=%d\n", result)
    `,
			Expected: "fib=13",
			Phase: 4,
		},
		TestCase{
			Name:    "func_nested_calls",
			Globals: `
    func double(x: int32) int32 {
        return x * 2
    }
    func triple(x: int32) int32 {
        return x * 3
    }
    `,
			Body: `
    let result = double(triple(5))
    printf("result=%d\n", result)
    `,
			Expected: "result=30",
			Phase: 4,
		},
		TestCase{
			Name:    "func_early_return",
			Globals: `
    func check_positive(x: int32) int32 {
        if x < 0 {
            return 0
        }
        return x
    }
    `,
			Body: `
    let result = check_positive(-5)
    printf("result=%d\n", result)
    `,
			Expected: "result=0",
			Phase: 4,
		},
		TestCase{
			Name:    "func_pass_by_value",
			Globals: `
    func modify(x: int32) {
        x = 99
    }
    `,
			Body: `
    let original = 10
    modify(original)
    printf("original=%d\n", original)
    `,
			Expected: "original=10",
			Phase: 4,
		},
		TestCase{
			Name: "extern_libc_printf",
			Body: `
    printf("extern_test\n")
    `,
			Expected: "extern_test",
			Phase: 4,
		},
		TestCase{
			Name: "extern_libc_malloc",
			Body: `
    let ptr = malloc(64)
    if ptr != null {
        printf("malloc_success\n")
    }
    free(ptr)
    `,
			Expected: "malloc_success",
			Phase: 4,
		},
	)
}