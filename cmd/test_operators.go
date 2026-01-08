package main

func init() {
	// Phase 2: Operators & Expressions
	AllTests = append(AllTests,
		TestCase{
			Name: "op_addition",
			Body: `
    let a = 10
    let b = 20
    let sum = a + b
    printf("sum=%d\n", sum)
    `,
			Expected: "sum=30",
			Phase: 2,
		},
		TestCase{
			Name: "op_subtraction",
			Body: `
    let diff = 50 - 15
    printf("diff=%d\n", diff)
    `,
			Expected: "diff=35",
			Phase: 2,
		},
		TestCase{
			Name: "op_multiplication",
			Body: `
    let prod = 7 * 8
    printf("prod=%d\n", prod)
    `,
			Expected: "prod=56",
			Phase: 2,
		},
		TestCase{
			Name: "op_division_signed",
			Body: `
    let quot: int32 = -20 / 4
    printf("quot=%d\n", quot)
    `,
			Expected: "quot=-5",
			Phase: 2,
		},
		TestCase{
			Name: "op_division_unsigned",
			Body: `
    let quot: uint32 = 20 / 4
    printf("quot=%u\n", quot)
    `,
			Expected: "quot=5",
			Phase: 2,
		},
		TestCase{
			Name: "op_modulo",
			Body: `
    let rem = 17 % 5
    printf("rem=%d\n", rem)
    `,
			Expected: "rem=2",
			Phase: 2,
		},
		TestCase{
			Name: "op_unary_negation",
			Body: `
    let x = 100
    let neg = -x
    printf("neg=%d\n", neg)
    `,
			Expected: "neg=-100",
			Phase: 2,
		},
		TestCase{
			Name: "op_bitwise_and",
			Body: `
    let result = 0xF0 & 0x0F
    printf("and=%d\n", result)
    `,
			Expected: "and=0",
			Phase: 2,
		},
		TestCase{
			Name: "op_bitwise_or",
			Body: `
    let result = 0xF0 | 0x0F
    printf("or=%d\n", result)
    `,
			Expected: "or=255",
			Phase: 2,
		},
		TestCase{
			Name: "op_bitwise_xor",
			Body: `
    let result = 0xFF ^ 0x0F
    printf("xor=%d\n", result)
    `,
			Expected: "xor=240",
			Phase: 2,
		},
		TestCase{
			Name: "op_bitwise_not",
			Body: `
    let x: uint8 = 0x0F
    let result = ~x
    printf("not=%u\n", result)
    `,
			Expected: "not=240",
			Phase: 2,
		},
		TestCase{
			Name: "op_shift_left",
			Body: `
    let result = 1 << 4
    printf("shl=%d\n", result)
    `,
			Expected: "shl=16",
			Phase: 2,
		},
		TestCase{
			Name: "op_shift_right_logical",
			Body: `
    let x: uint32 = 64
    let result = x >> 2
    printf("shr=%u\n", result)
    `,
			Expected: "shr=16",
			Phase: 2,
		},
		TestCase{
			Name: "op_shift_right_arithmetic",
			Body: `
    let x: int32 = -64
    let result = x >> 2
    printf("shr=%d\n", result)
    `,
			Expected: "shr=-16",
			Phase: 2,
		},
		TestCase{
			Name: "op_logical_and",
			Body: `
    let result = true && true
    printf("and=%d\n", result)
    `,
			Expected: "and=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_logical_or",
			Body: `
    let result = false || true
    printf("or=%d\n", result)
    `,
			Expected: "or=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_logical_not",
			Body: `
    let result = !false
    printf("not=%d\n", result)
    `,
			Expected: "not=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_equality",
			Body: `
    let result = (42 == 42)
    printf("eq=%d\n", result)
    `,
			Expected: "eq=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_inequality",
			Body: `
    let result = (10 != 20)
    printf("ne=%d\n", result)
    `,
			Expected: "ne=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_less_than",
			Body: `
    let result = (5 < 10)
    printf("lt=%d\n", result)
    `,
			Expected: "lt=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_less_equal",
			Body: `
    let result = (10 <= 10)
    printf("le=%d\n", result)
    `,
			Expected: "le=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_greater_than",
			Body: `
    let result = (20 > 10)
    printf("gt=%d\n", result)
    `,
			Expected: "gt=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_greater_equal",
			Body: `
    let result = (15 >= 15)
    printf("ge=%d\n", result)
    `,
			Expected: "ge=1",
			Phase: 2,
		},
		TestCase{
			Name: "op_compound_add",
			Body: `
    let x = 10
    x += 5
    printf("x=%d\n", x)
    `,
			Expected: "x=15",
			Phase: 2,
		},
		TestCase{
			Name: "op_compound_sub",
			Body: `
    let x = 20
    x -= 7
    printf("x=%d\n", x)
    `,
			Expected: "x=13",
			Phase: 2,
		},
		TestCase{
			Name: "op_compound_mul",
			Body: `
    let x = 5
    x *= 3
    printf("x=%d\n", x)
    `,
			Expected: "x=15",
			Phase: 2,
		},
		TestCase{
			Name: "op_compound_div",
			Body: `
    let x = 100
    x /= 4
    printf("x=%d\n", x)
    `,
			Expected: "x=25",
			Phase: 2,
		},
		TestCase{
			Name: "op_compound_mod",
			Body: `
    let x = 17
    x %= 5
    printf("x=%d\n", x)
    `,
			Expected: "x=2",
			Phase: 2,
		},
		TestCase{
			Name: "op_increment_postfix",
			Body: `
    let x = 10
    let old = x++
    printf("old=%d new=%d\n", old, x)
    `,
			Expected: "old=10 new=11",
			Phase: 2,
		},
		TestCase{
			Name: "op_increment_prefix",
			Body: `
    let x = 10
    let new_val = ++x
    printf("new=%d\n", new_val)
    `,
			Expected: "new=11",
			Phase: 2,
		},
		TestCase{
			Name: "op_decrement_postfix",
			Body: `
    let x = 10
    let old = x--
    printf("old=%d new=%d\n", old, x)
    `,
			Expected: "old=10 new=9",
			Phase: 2,
		},
		TestCase{
			Name: "op_decrement_prefix",
			Body: `
    let x = 10
    let new_val = --x
    printf("new=%d\n", new_val)
    `,
			Expected: "new=9",
			Phase: 2,
		},
		TestCase{
			Name: "op_precedence",
			Body: `
    let result = 2 + 3 * 4
    printf("result=%d\n", result)
    `,
			Expected: "result=14",
			Phase: 2,
		},
		TestCase{
			Name: "op_parentheses",
			Body: `
    let result = (2 + 3) * 4
    printf("result=%d\n", result)
    `,
			Expected: "result=20",
			Phase: 2,
		},
	)
}