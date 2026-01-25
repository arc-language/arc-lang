package main

func init() {
	// Phase 9: Intrinsics
	AllTests = append(AllTests,
		TestCase{
			Name: "intrinsic_sizeof_int32",
			Body: `
    let size = sizeof(int32)
    printf("size=%zu\n", size)
    `,
			Expected: "size=4",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_sizeof_int64",
			Body: `
    let size = sizeof(int64)
    printf("size=%zu\n", size)
    `,
			Expected: "size=8",
			Phase: 9,
		},
		TestCase{
			Name:    "intrinsic_sizeof_struct",
			Globals: `
    struct Pair {
        a: int32
        b: int32
    }
    `,
			Body: `
    let size = sizeof(Pair)
    printf("size=%zu\n", size)
    `,
			Expected: "size=8",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_alignof",
			Body: `
    let align = alignof(int64)
    printf("align=%zu\n", align)
    `,
			Expected: "align=8",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_alloca_single",
			Body: `
    let ptr = alloca<int32>()
    *ptr = 777
    printf("value=%d\n", *ptr)
    `,
			Expected: "value=777",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_alloca_array",
			Body: `
    let buffer = alloca<int32>(5)
    buffer[0] = 10
    buffer[1] = 20
    buffer[2] = 30
    printf("buffer[1]=%d\n", buffer[1])
    `,
			Expected: "buffer[1]=20",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_memset",
			Body: `
    let buffer = alloca<byte>(10)
    memset(buffer, 0x42, 10)
    printf("buffer[5]=%d\n", buffer[5])
    `,
			Expected: "buffer[5]=66",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_memcpy",
			Body: `
    let src = alloca<int32>(3)
    src[0] = 100
    src[1] = 200
    src[2] = 300
    
    let dest = alloca<int32>(3)
    memcpy(dest, src, 12)
    
    printf("dest[1]=%d\n", dest[1])
    `,
			Expected: "dest[1]=200",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_strlen",
			Body: `
    let str: *byte = "hello"
    let len = strlen(str)
    printf("len=%zu\n", len)
    `,
			Expected: "len=5",
			Phase: 9,
		},
		TestCase{
			Name: "intrinsic_bitcast",
			Body: `
    let f: float32 = 1.0
    let bits = bit_cast<uint32>(f)
    printf("bits=%u\n", bits)
    `,
			Expected: "bits=1065353216",
			Phase: 9,
		},
	)
}