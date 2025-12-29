package main

func init() {
	RegisterTest("Alloca Single", func() string {
		return `let ptr = alloca(int32)`
	})

	RegisterTest("Alloca Array", func() string {
		return `let buffer = alloca(byte, 1024)`
	})

	RegisterTest("Indexed Pointer Read", func() string {
		return `func test() {
    let byte_val = buffer[5]
}`
	})

	RegisterTest("Indexed Pointer Write", func() string {
		return `func test() {
    buffer[10] = 0x42
}`
	})

	RegisterTest("Indexed Pointer Array", func() string {
		return `func test() {
    let ptr: *int32 = array_base
    let third_element = ptr[2]
    ptr[3] = 100
}`
	})

	RegisterTest("Type Cast", func() string {
		return `let result = cast<int64>(value)`
	})

	RegisterTest("Type Cast Pointer", func() string {
		return `let byte_ptr = cast<*byte>(int_ptr)`
	})

	RegisterTest("Type Cast Pointer To Int", func() string {
		return `let addr = cast<uint64>(ptr)`
	})

	RegisterTest("Type Cast Int To Pointer", func() string {
		return `let new_ptr = cast<*int32>(addr)`
	})

	RegisterTest("Type Cast Void Pointer", func() string {
		return `let generic = cast<*void>(typed_ptr)`
	})

	RegisterTest("Intrinsic sizeof", func() string {
		return `let sz = sizeof<int32>`
	})

	RegisterTest("Intrinsic sizeof Struct", func() string {
		return `let st_sz = sizeof<Stat>`
	})

	RegisterTest("Intrinsic alignof", func() string {
		return `let align = alignof<float64>`
	})

	RegisterTest("Intrinsic memset", func() string {
		return `func test() {
    let buf = alloca(byte, 1024)
    memset(buf, 0, 1024)
}`
	})

	RegisterTest("Intrinsic memcpy", func() string {
		return `func test() {
    memcpy(dest_ptr, src_ptr, 1024)
}`
	})

	RegisterTest("Intrinsic memmove", func() string {
		return `func test() {
    memmove(dest_ptr, src_ptr, 1024)
}`
	})

	RegisterTest("Intrinsic strlen", func() string {
		return `func test() {
    let cstr: *byte = "hello\0"
    let len = strlen(cstr)
}`
	})

	RegisterTest("Intrinsic memchr", func() string {
		return `func test() {
    let buf: *byte = "hello\nworld"
    let newline = memchr(buf, '\n', 11)
}`
	})

	RegisterTest("Intrinsic va_start", func() string {
		return `func printf(fmt: string, ...) {
    let args = va_start(fmt)
}`
	})

	RegisterTest("Intrinsic va_arg", func() string {
		return `func printf(fmt: string, ...) {
    let args = va_start(fmt)
    let val = va_arg<int32>(args)
}`
	})

	RegisterTest("Intrinsic va_end", func() string {
		return `func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)
}`
	})

	RegisterTest("Intrinsic raise", func() string {
		return `func test() {
    if ptr == null {
        raise("Memory corrupted")
    }
}`
	})

	RegisterTest("Intrinsic memcmp", func() string {
		return `func test() {
    let diff = memcmp(ptr1, ptr2, 1024)
}`
	})

	RegisterTest("Intrinsic bit_cast", func() string {
		return `func test() {
    let f: float32 = 1.0
    let bits = bit_cast<uint32>(f)
}`
	})

	RegisterTest("Syscall Write", func() string {
		return `func test() {
    let msg = "Hello, Direct Syscall!\n"
    let len = 23
    let result = syscall(SYS_WRITE, STDOUT, msg, len)
}`
	})
}