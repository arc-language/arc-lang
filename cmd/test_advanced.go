package main

func init() {
	// Phase 10: Advanced Features (Requires Runtime)
	AllTests = append(AllTests,
		TestCase{
			Name:    "generic_struct_def",
			Globals: `
    struct Box<T> {
        value: T
    }
    `,
			Body: `
    let b: Box<int32> = Box<int32>{value: 42}
    printf("value=%d\n", b.value)
    `,
			Expected: "value=42",
			Phase: 10,
		},
		TestCase{
			Name:    "generic_func_def",
			Globals: `
    func identity<T>(x: T) T {
        return x
    }
    `,
			Body: `
    let result = identity<int32>(99)
    printf("result=%d\n", result)
    `,
			Expected: "result=99",
			Phase: 10,
		},
		TestCase{
			Name:    "generic_swap",
			Globals: `
    func swap<T>(a: *T, b: *T) {
        let tmp: T = *a
        *a = *b
        *b = tmp
    }
    `,
			Body: `
    let x = 10
    let y = 20
    swap<int32>(&x, &y)
    printf("x=%d y=%d\n", x, y)
    `,
			Expected: "x=20 y=10",
			Phase: 10,
		},
		TestCase{
			Name:    "enum_basic",
			Globals: `
    enum Status {
        OK
        ERROR
        PENDING
    }
    `,
			Body: `
    let s: Status = Status.OK
    printf("status=%d\n", s)
    `,
			Expected: "status=0",
			Phase: 10,
		},
		TestCase{
			Name:    "enum_explicit_values",
			Globals: `
    enum HttpCode {
        OK = 200
        NOT_FOUND = 404
        SERVER_ERROR = 500
    }
    `,
			Body: `
    let code: HttpCode = HttpCode.NOT_FOUND
    printf("code=%d\n", code)
    `,
			Expected: "code=404",
			Phase: 10,
		},
		TestCase{
			Name: "exec_thread_placeholder",
			Body: `
    printf("thread_placeholder\n")
    `,
			Expected: "thread_placeholder",
			Phase: 10,
		},
		TestCase{
			Globals: `
    async func exec_async_placeholder() int32 {
		sleep(1)
		let results = 100
		return results
    }
    `,
			Name: "exec_async_placeholder",
			Body: `
	let results = await exec_async_placeholder()
    printf("%d\n", results)
	sleep(4)
    `,
			Expected: "100",
			Phase: 10,
		},
		TestCase{
			Name: "for_in_range",
			Body: `
    let sum = 0
    for i in 0..5 {
        sum += i
    }
    printf("sum=%d\n", sum)
    `,
			Expected: "sum=10",
			Phase: 10,
		},
		TestCase{
			Name: "try_except_placeholder",
			Body: `
    printf("exception_placeholder\n")
    `,
			Expected: "exception_placeholder",
			Phase: 10,
		},
	)
}