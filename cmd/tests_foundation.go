package main

func init() {
	// Phase 1: Foundation (Literals & Variables)
	AllTests = append(AllTests,
		TestCase{
			Name: "literal_int_decimal",
			Body: `
    let x = 42
    printf("value=%d\n", x)
    `,
			Expected: "value=42",
			Phase: 1,
		},
		TestCase{
			Name: "literal_int_hex",
			Body: `
    let x = 0xFF
    printf("hex=%d\n", x)
    `,
			Expected: "hex=255",
			Phase: 1,
		},
		TestCase{
			Name: "literal_int_negative",
			Body: `
    let x = -100
    printf("neg=%d\n", x)
    `,
			Expected: "neg=-100",
			Phase: 1,
		},
		TestCase{
			Name: "literal_float_standard",
			Body: `
    let pi: float32 = 3.14
    printf("pi=%.2f\n", pi)
    `,
			Expected: "pi=3.14",
			Phase: 1,
		},
		TestCase{
			Name: "literal_float_scientific",
			Body: `
    let x: float64 = 1.5e-3
    printf("sci=%.4f\n", x)
    `,
			Expected: "sci=0.0015",
			Phase: 1,
		},
		TestCase{
			Name: "literal_bool",
			Body: `
    let t = true
    let f = false
    printf("true=%d false=%d\n", t, f)
    `,
			Expected: "true=1 false=0",
			Phase: 1,
		},
		TestCase{
			Name: "literal_char",
			Body: `
    let c: char = 'A'
    printf("char=%c code=%d\n", c, c)
    `,
			Expected: "char=A code=65",
			Phase: 1,
		},
		TestCase{
			Name: "literal_char_escape",
			Body: `
    let newline: char = '\n'
    let tab: char = '\t'
    printf("escape_test%cok\n", newline)
    `,
			Expected: "escape_test\nok",
			Phase: 1,
		},
		TestCase{
			Name: "literal_string_basic",
			Body: `
    let msg: string = "hello"
    printf("%s\n", msg)
    `,
			Expected: "hello",
			Phase: 1,
		},
		TestCase{
			Name: "literal_string_escape",
			Body: `
    let msg: string = "line1\nline2"
    printf("%s\n", msg)
    `,
			Expected: "line1\nline2",
			Phase: 1,
		},
		TestCase{
			Name: "literal_null",
			Body: `
    let ptr: *int32 = null
    if ptr == null {
        printf("ptr_is_null\n")
    }
    `,
			Expected: "ptr_is_null",
			Phase: 1,
		},
		TestCase{
			Name: "variable_mutable_typed",
			Body: `
    let x: int32 = 10
    x = 20
    printf("x=%d\n", x)
    `,
			Expected: "x=20",
			Phase: 1,
		},
		TestCase{
			Name: "variable_mutable_inferred",
			Body: `
    let x = 99
    x = 100
    printf("inferred=%d\n", x)
    `,
			Expected: "inferred=100",
			Phase: 1,
		},
		TestCase{
			Name: "constant_typed",
			Body: `
    const MAX: int32 = 255
    printf("max=%d\n", MAX)
    `,
			Expected: "max=255",
			Phase: 1,
		},
		TestCase{
			Name: "constant_inferred",
			Body: `
    const LIMIT = 1024
    printf("limit=%d\n", LIMIT)
    `,
			Expected: "limit=1024",
			Phase: 1,
		},
		TestCase{
			Name: "type_int8",
			Body: `
    let small: int8 = -128
    printf("int8=%d\n", small)
    `,
			Expected: "int8=-128",
			Phase: 1,
		},
		TestCase{
			Name: "type_int16",
			Body: `
    let medium: int16 = 32000
    printf("int16=%d\n", medium)
    `,
			Expected: "int16=32000",
			Phase: 1,
		},
		TestCase{
			Name: "type_int32",
			Body: `
    let x: int32 = 2147483647
    printf("int32=%d\n", x)
    `,
			Expected: "int32=2147483647",
			Phase: 1,
		},
		TestCase{
			Name: "type_int64",
			Body: `
    let big: int64 = 9223372036854775807
    printf("int64=%lld\n", big)
    `,
			Expected: "int64=9223372036854775807",
			Phase: 1,
		},
		TestCase{
			Name: "type_uint8",
			Body: `
    let byte_val: uint8 = 255
    printf("uint8=%u\n", byte_val)
    `,
			Expected: "uint8=255",
			Phase: 1,
		},
		TestCase{
			Name: "type_uint16",
			Body: `
    let port: uint16 = 65535
    printf("uint16=%u\n", port)
    `,
			Expected: "uint16=65535",
			Phase: 1,
		},
		TestCase{
			Name: "type_uint32",
			Body: `
    let x: uint32 = 4294967295
    printf("uint32=%u\n", x)
    `,
			Expected: "uint32=4294967295",
			Phase: 1,
		},
		TestCase{
			Name: "type_uint64",
			Body: `
    let huge: uint64 = 18446744073709551615
    printf("uint64=%llu\n", huge)
    `,
			Expected: "uint64=18446744073709551615",
			Phase: 1,
		},
		TestCase{
			Name: "type_usize",
			Body: `
    let size: usize = 1024
    printf("usize=%zu\n", size)
    `,
			Expected: "usize=1024",
			Phase: 1,
		},
		TestCase{
			Name: "type_isize",
			Body: `
    let offset: isize = -256
    printf("isize=%zd\n", offset)
    `,
			Expected: "isize=-256",
			Phase: 1,
		},
		TestCase{
			Name: "type_float32",
			Body: `
    let x: float32 = 3.14159
    printf("float32=%.5f\n", x)
    `,
			Expected: "float32=3.14159",
			Phase: 1,
		},
		TestCase{
			Name: "type_float64",
			Body: `
    let x: float64 = 2.718281828459045
    printf("float64=%.15f\n", x)
    `,
			Expected: "float64=2.718281828459045",
			Phase: 1,
		},
		TestCase{
			Name: "type_alias_byte",
			Body: `
    let b: byte = 0x42
    printf("byte=%u\n", b)
    `,
			Expected: "byte=66",
			Phase: 1,
		},
		TestCase{
			Name: "type_alias_bool",
			Body: `
    let flag: bool = true
    printf("bool=%d\n", flag)
    `,
			Expected: "bool=1",
			Phase: 1,
		},
		TestCase{
			Name: "type_alias_char",
			Body: `
    let ch: char = 'Z'
    printf("char=%c\n", ch)
    `,
			Expected: "char=Z",
			Phase: 1,
		},
	)
}