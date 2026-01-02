package main

func init() {
	// Phase 11: Edge Cases & Complex Scenarios
	AllTests = append(AllTests,
		TestCase{
			Name: "edge_int_overflow",
			Body: `
    let x: int8 = 127
    x = x + 1
    libc.printf("overflow=%d\n", x)
    `,
			Expected: "overflow=-128",
			Phase: 11,
		},
		TestCase{
			Name: "edge_div_zero_skip",
			Body: `
    let x = 10
    let y = 0
    libc.printf("divzero_skipped\n")
    `,
			Expected: "divzero_skipped",
			Phase: 11,
		},
		TestCase{
			Name: "edge_null_deref_skip",
			Body: `
    let ptr: *int32 = null
    libc.printf("null_deref_skipped\n")
    `,
			Expected: "null_deref_skipped",
			Phase: 11,
		},
		TestCase{
			Name: "edge_pointer_lifetime",
			Body: `
    let stack_value = 42
    let ptr: *int32 = &stack_value
    libc.printf("value=%d\n", *ptr)
    `,
			Expected: "value=42",
			Phase: 11,
		},
		TestCase{
			Name:    "edge_struct_padding",
			Globals: `
    struct Padded {
        a: int8
        b: int64
        c: int8
    }
    `,
			Body: `
    let size = sizeof<Padded>
    libc.printf("size=%zu\n", size)
    `,
			Expected: "size=24",
			Phase: 11,
		},
		TestCase{
			Name: "edge_function_pointer_placeholder",
			Body: `
    libc.printf("func_ptr_placeholder\n")
    `,
			Expected: "func_ptr_placeholder",
			Phase: 11,
		},
		TestCase{
			Name:    "edge_deep_recursion",
			Globals: `
    func countdown(n: int32) int32 {
        if n <= 0 {
            return 0
        }
        return countdown(n - 1) + 1
    }
    `,
			Body: `
    let result = countdown(100)
    libc.printf("depth=%d\n", result)
    `,
			Expected: "depth=100",
			Phase: 11,
		},
		TestCase{
			Name: "edge_large_stack_array",
			Body: `
    let arr: array<int32, 1000> = {}
    arr[999] = 555
    libc.printf("last=%d\n", arr[999])
    `,
			Expected: "last=555",
			Phase: 11,
		},
		TestCase{
			Name:    "edge_eval_order",
			Globals: `
    func side_effect(x: int32) int32 {
        libc.printf("eval_%d ", x)
        return x
    }
    `,
			Body: `
    let result = side_effect(1) + side_effect(2) * side_effect(3)
    libc.printf("result=%d\n", result)
    `,
			Expected: "eval_1 eval_2 eval_3 result=7",
			Phase: 11,
		},
		TestCase{
			Name: "edge_zero_sized_array",
			Body: `
    let arr: array<int32, 0> = {}
    libc.printf("zero_array_ok\n")
    `,
			Expected: "zero_array_ok",
			Phase: 11,
		},
	)
}