package main

func init() {
	// Phase 5: Pointers & Memory
	AllTests = append(AllTests,
		TestCase{
			Name: "pointer_basic_type",
			Body: `
    let x = 42
    let ptr: *int32 = &x
    libc.printf("pointer_created\n")
    `,
			Expected: "pointer_created",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_address_of",
			Body: `
    let value: int32 = 100
    let ptr: *int32 = &value
    libc.printf("got_address\n")
    `,
			Expected: "got_address",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_deref_read",
			Body: `
    let value = 50
    let ptr: *int32 = &value
    let retrieved = *ptr
    libc.printf("value=%d\n", retrieved)
    `,
			Expected: "value=50",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_deref_write",
			Body: `
    let value = 10
    let ptr: *int32 = &value
    *ptr = 99
    libc.printf("value=%d\n", value)
    `,
			Expected: "value=99",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_null_check",
			Body: `
    let ptr: *int32 = null
    if ptr == null {
        libc.printf("is_null\n")
    }
    `,
			Expected: "is_null",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_non_null",
			Body: `
    let value = 42
    let ptr: *int32 = &value
    if ptr != null {
        libc.printf("not_null\n")
    }
    `,
			Expected: "not_null",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_arithmetic_add",
			Body: `
    let arr = alloca(int32, 5)
    arr[0] = 10
    arr[1] = 20
    let ptr: *int32 = arr
    let next = ptr + 1
    libc.printf("value=%d\n", *next)
    `,
			Expected: "value=20",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_arithmetic_sub",
			Body: `
    let arr = alloca(int32, 5)
    arr[0] = 10
    arr[1] = 20
    let ptr: *int32 = arr + 1
    let prev = ptr - 1
    libc.printf("value=%d\n", *prev)
    `,
			Expected: "value=10",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_array_indexing",
			Body: `
    let arr = alloca(int32, 3)
    arr[0] = 100
    arr[1] = 200
    arr[2] = 300
    libc.printf("arr[1]=%d\n", arr[1])
    `,
			Expected: "arr[1]=200",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_void_type",
			Body: `
    let x = 42
    let vptr: *void = cast<*void>(&x)
    let iptr: *int32 = cast<*int32>(vptr)
    libc.printf("value=%d\n", *iptr)
    `,
			Expected: "value=42",
			Phase: 5,
		},
		TestCase{
			Name: "pointer_double_pointer",
			Body: `
    let value = 77
    let ptr: *int32 = &value
    let pptr: **int32 = &ptr
    libc.printf("value=%d\n", **pptr)
    `,
			Expected: "value=77",
			Phase: 5,
		},
		TestCase{
			Name:    "pointer_pass_to_function",
			Globals: `
    func modify_value(ptr: *int32) {
        *ptr = 888
    }
    `,
			Body: `
    let x = 10
    modify_value(&x)
    libc.printf("x=%d\n", x)
    `,
			Expected: "x=888",
			Phase: 5,
		},
	)
}