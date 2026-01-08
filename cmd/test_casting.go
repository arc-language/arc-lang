package main

func init() {
	// Phase 8: Type Casting
	AllTests = append(AllTests,
		TestCase{
			Name: "cast_int_widen",
			Body: `
    let small: int8 = 100
    let large: int32 = cast<int32>(small)
    printf("large=%d\n", large)
    `,
			Expected: "large=100",
			Phase: 8,
		},
		TestCase{
			Name: "cast_int_truncate",
			Body: `
    let large: int32 = 1000
    let small: int8 = cast<int8>(large)
    printf("small=%d\n", small)
    `,
			Expected: "small=-24",
			Phase: 8,
		},
		TestCase{
			Name: "cast_int_to_float",
			Body: `
    let i: int32 = 42
    let f: float32 = cast<float32>(i)
    printf("float=%.1f\n", f)
    `,
			Expected: "float=42.0",
			Phase: 8,
		},
		TestCase{
			Name: "cast_float_to_int",
			Body: `
    let f: float32 = 3.9
    let i: int32 = cast<int32>(f)
    printf("int=%d\n", i)
    `,
			Expected: "int=3",
			Phase: 8,
		},
		TestCase{
			Name: "cast_ptr_to_ptr",
			Body: `
    let x: int32 = 42
    let iptr: *int32 = &x
    let vptr: *void = cast<*void>(iptr)
    let back: *int32 = cast<*int32>(vptr)
    printf("value=%d\n", *back)
    `,
			Expected: "value=42",
			Phase: 8,
		},
		TestCase{
			Name: "cast_ptr_to_int",
			Body: `
    let x: int32 = 99
    let ptr: *int32 = &x
    let addr: uint64 = cast<uint64>(ptr)
    printf("addr_non_zero=%d\n", addr != 0)
    `,
			Expected: "addr_non_zero=1",
			Phase: 8,
		},
		TestCase{
			Name: "cast_int_to_ptr",
			Body: `
    let addr: uint64 = 0x1000
    let ptr: *int32 = cast<*int32>(addr)
    printf("ptr_created\n")
    `,
			Expected: "ptr_created",
			Phase: 8,
		},
		TestCase{
			Name: "cast_signed_to_unsigned",
			Body: `
    let neg: int32 = -1
    let pos: uint32 = cast<uint32>(neg)
    printf("unsigned=%u\n", pos)
    `,
			Expected: "unsigned=4294967295",
			Phase: 8,
		},
		TestCase{
			Name: "cast_unsigned_to_signed",
			Body: `
    let big: uint32 = 4294967295
    let neg: int32 = cast<int32>(big)
    printf("signed=%d\n", neg)
    `,
			Expected: "signed=-1",
			Phase: 8,
		},
	)
}