package main

func init() {
	// Phase 7: Arrays & Collections
	AllTests = append(AllTests,
		TestCase{
			Name: "array_fixed_size",
			Body: `
    let arr: array<int32, 3> = {10, 20, 30}
    printf("arr[1]=%d\n", arr[1])
    `,
			Expected: "arr[1]=20",
			Phase: 7,
		},
		TestCase{
			Name: "array_zero_init",
			Body: `
    let arr: array<int32, 5> = {}
    printf("arr[0]=%d\n", arr[0])
    `,
			Expected: "arr[0]=0",
			Phase: 7,
		},
		TestCase{
			Name: "array_element_assign",
			Body: `
    let arr: array<int32, 3> = {1, 2, 3}
    arr[1] = 99
    printf("arr[1]=%d\n", arr[1])
    `,
			Expected: "arr[1]=99",
			Phase: 7,
		},
		TestCase{
			Name: "array_ptr_first",
			Body: `
    let arr: array<int32, 3> = {100, 200, 300}
    let ptr: *int32 = &arr[0]
    printf("first=%d\n", *ptr)
    `,
			Expected: "first=100",
			Phase: 7,
		},
		TestCase{
			Name: "array_iteration_cstyle",
			Body: `
    let arr: array<int32, 4> = {1, 2, 3, 4}
    let sum = 0
    for let i: usize = 0; i < 4; i++ {
        sum += arr[i]
    }
    printf("sum=%d\n", sum)
    `,
			Expected: "sum=10",
			Phase: 7,
		},
		TestCase{
			Name: "array_2d",
			Body: `
    let matrix: array<array<int32, 2>, 2> = {
        {1, 2},
        {3, 4}
    }
    printf("matrix[1][1]=%d\n", matrix[1][1])
    `,
			Expected: "matrix[1][1]=4",
			Phase: 7,
		},
		TestCase{
			Name: "vector_placeholder",
			Body: `
    printf("vector_placeholder\n")
    `,
			Expected: "vector_placeholder",
			Phase: 7,
		},
		TestCase{
			Name: "map_placeholder",
			Body: `
    printf("map_placeholder\n")
    `,
			Expected: "map_placeholder",
			Phase: 7,
		},
	)
}