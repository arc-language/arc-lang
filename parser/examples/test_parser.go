package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
)

// CustomErrorListener implements ANTLR error listener for better error reporting
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]string, 0),
	}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)
	c.Errors = append(c.Errors, errorMsg)
}

// ParseArcSource parses Arc source code from a string
func ParseArcSource(source string) error {
	// Create input stream from string
	input := antlr.NewInputStream(source)

	// Create the lexer
	lexer := parser.NewArcLexer(input)
	
	// Remove default error listeners and add custom one
	lexer.RemoveErrorListeners()
	errorListener := NewCustomErrorListener()
	lexer.AddErrorListener(errorListener)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewArcParser(stream)
	
	// Remove default error listeners and add custom one
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the compilation unit (starting rule)
	_ = p.CompilationUnit()

	// Check for parsing errors
	if len(errorListener.Errors) > 0 {
		return fmt.Errorf("parsing failed: %s", strings.Join(errorListener.Errors, "; "))
	}

	return nil
}

// Test functions that return Arc code snippets

func test_import_single() string {
	return `import "some/path/package"`
}

func test_import_multiple() string {
	return `import (
    "std/io"
    "std/os"
    "github.com/user/physics"
    "github.com/user/graphics"
    "gitlab.com/company/lib"
)`
}

func test_namespace() string {
	return `namespace main`
}

func test_variable_mutable_typed() string {
	return `let x: int32 = 42`
}

func test_variable_mutable_inferred() string {
	return `let x = 42`
}

func test_constant_typed() string {
	return `const x: int32 = 42`
}

func test_constant_inferred() string {
	return `const x = 42`
}

func test_basic_types_signed_integers() string {
	return `let i: int32 = -500`
}

func test_basic_types_unsigned_integers() string {
	return `let u: uint64 = 10000`
}

func test_basic_types_usize() string {
	return `let len: usize = 100`
}

func test_basic_types_isize() string {
	return `let offset: isize = -4`
}

func test_basic_types_float32() string {
	return `let f32: float32 = 3.14`
}

func test_basic_types_float64() string {
	return `let f64: float64 = 2.71828`
}

func test_basic_types_byte() string {
	return `let b: byte = 255`
}

func test_basic_types_bool() string {
	return `let flag: bool = true`
}

func test_basic_types_char() string {
	return `let r: char = 'a'`
}

func test_basic_types_string() string {
	return `let s: string = "hello"`
}

func test_null_literal() string {
	return `let ptr: *int32 = null`
}

func test_null_check() string {
	return `func test() {
    if ptr == null {
        x = 1
    }
}`
}

func test_pointer_basic() string {
	return `let ptr: *int32 = &value`
}

func test_pointer_void() string {
	return `let handle: *void = alloca(void, 64)`
}

func test_reference_basic() string {
	return `let ref: &int32 = value`
}

func test_vector_type() string {
	return `let v: vector<int32> = {}`
}

func test_vector_literal_empty() string {
	return `let empty: vector<int32> = {}`
}

func test_vector_literal_with_values() string {
	return `let nums: vector<int32> = {1, 2, 3, 4, 5}`
}

func test_vector_literal_inferred() string {
	return `let items = {10, 20, 30}`
}

func test_map_type() string {
	return `let m: map<string, int32> = {}`
}

func test_map_literal_empty() string {
	return `let empty: map<string, int32> = {}`
}

func test_map_literal_with_values() string {
	return `let scores: map<string, int32> = {"alice": 100, "bob": 95}`
}

func test_map_literal_inferred() string {
	return `let config = {"host": "localhost", "port": "8080"}`
}

func test_function_basic() string {
	return `func add(a: int32, b: int32) int32 {
    return a + b
}`
}

func test_function_no_return() string {
	return `func print(msg: string) {
}`
}

func test_function_void_return() string {
	return `func main() {
    let x = 42
}`
}

func test_function_async() string {
	return `async func fetch_data(url: string) string {
    let response = await http_get(url)
    return response
}`
}

func test_function_async_no_return() string {
	return `async func process_items(items: vector<string>) {
    for item in items {
        await process(item)
    }
}`
}

func test_await_expression() string {
	return `async func main() {
    let data = await fetch_data("https://api.example.com")
}`
}

func test_await_multiple() string {
	return `async func main() {
    let result1 = await task1()
    let result2 = await task2()
}`
}

func test_await_in_if() string {
	return `async func main() {
    if await check_status() {
        x = 1
    }
}`
}

func test_struct_declaration() string {
	return `struct Point {
    x: int32
    y: int32
}`
}

func test_struct_literal() string {
	return `let p = Point{x: 10, y: 20}`
}

func test_struct_literal_type_inference() string {
	return `let p2 = Point{x: 5, y: 15}`
}

func test_struct_literal_empty() string {
	return `let p3: Point = Point{}`
}

func test_struct_field_access() string {
	return `let x = p.x`
}

func test_struct_field_assignment() string {
	return `func test() {
    p.y = 30
}`
}

func test_struct_inline_method() string {
	return `struct Point {
    x: int32
    y: int32
    
    func distance(self p: Point) float64 {
        return cast<float64>(p.x * p.x + p.y * p.y)
    }
    
    func move(self p: *Point, dx: int32, dy: int32) {
        p.x = p.x + dx
        p.y = p.y + dy
    }
}`
}

func test_struct_mutating_method_inline() string {
	return `struct Counter {
    count: int32
    
    mutating increment(self c: *Counter) {
        c.count = c.count + 1
    }
    
    mutating add(self c: *Counter, value: int32) {
        c.count = c.count + value
    }
    
    func get_count(self c: Counter) int32 {
        return c.count
    }
}`
}

func test_struct_mutating_method_usage() string {
	return `func test() {
    let counter = Counter{count: 0}
    counter.increment()
    counter.add(5)
    let value = counter.get_count()
}`
}

func test_struct_flat_method() string {
	return `struct Point {
    x: int32
    y: int32
}

func distance(self p: Point) float64 {
    return cast<float64>(p.x * p.x + p.y * p.y)
}

func move(self p: *Point, dx: int32, dy: int32) {
    p.x = p.x + dx
    p.y = p.y + dy
}`
}

func test_struct_mutating_method_flat() string {
	return `struct Point {
    x: int32
    y: int32
}

mutating reset(self p: *Point) {
    p.x = 0
    p.y = 0
}`
}

func test_class_declaration() string {
	return `class Client {
    name: string
    port: int32
}`
}

func test_class_inline_method() string {
	return `class Client {
    name: string
    port: int32
    
    func connect(self c: *Client, host: string) bool {
        return true
    }
    
    deinit(self c: *Client) {
    }
}`
}

func test_class_async_method() string {
	return `class Client {
    name: string
    port: int32
    
    async func fetch_data(self c: *Client) string {
        let response = await http_get("https://example.com")
        return response
    }
}`
}

func test_class_flat_method() string {
	return `class Client {
    name: string
    port: int32
}

func connect(self c: *Client, host: string) bool {
    return true
}

deinit(self c: *Client) {
}`
}

func test_class_async_flat_method() string {
	return `class Client {
    name: string
    port: int32
}

async func fetch_data(self c: *Client) string {
    let response = await http_get("https://example.com")
    return response
}`
}

func test_method_declaration() string {
	return `struct Client {
    port: int32
}

func Connect(self c: *Client, host: string) bool {
    return true
}`
}

func test_method_call() string {
	return `func test() {
    c.Connect("localhost")
}`
}

func test_async_method_call() string {
	return `async func example() {
    let data = await c.fetch_data()
}`
}

func test_if_statement() string {
	return `func test() {
    if condition {
        x = 1
    }
}`
}

func test_if_else_statement() string {
	return `func test() {
    if condition {
        x = 1
    } else {
        x = 2
    }
}`
}

func test_if_else_if_statement() string {
	return `func test() {
    if condition {
        x = 1
    } else if condition2 {
        x = 2
    } else {
        x = 3
    }
}`
}

func test_for_loop_c_style() string {
	return `func test() {
    for let i = 0; i < 10; i = i + 1 {
        x = i
    }
}`
}

func test_for_loop_c_style_increment() string {
	return `func test() {
    for let i = 0; i < 10; i++ {
        x = i
    }
}`
}

func test_for_loop_condition() string {
	return `func test() {
    let j = 5
    for j > 0 {
        j--
    }
}`
}

func test_for_loop_infinite() string {
	return `func test() {
    let counter = 0
    for {
        counter++
        if counter >= 10 {
            break
        }
    }
}`
}

func test_for_in_loop_vector() string {
	return `func test() {
    let items: vector<int32> = {1, 2, 3, 4, 5}
    for item in items {
        x = item
    }
}`
}

func test_for_in_loop_map() string {
	return `func test() {
    let scores: map<string, int32> = {"alice": 100, "bob": 95}
    for key, value in scores {
        x = value
    }
}`
}

func test_for_in_loop_range() string {
	return `func test() {
    for i in 0..10 {
        x = i
    }
}`
}

func test_break_statement() string {
	return `func test() {
    for let i = 0; i < 10; i = i + 1 {
        if i == 5 {
            break
        }
    }
}`
}

func test_continue_statement() string {
	return `func test() {
    for let i = 0; i < 10; i = i + 1 {
        if i == 5 {
            continue
        }
        x = i
    }
}`
}

func test_defer_statement() string {
	return `func test() {
    let ptr = alloca(int32)
    defer free(ptr)
}`
}

func test_return_statement() string {
	return `func test() int32 {
    return 42
}`
}

func test_return_void() string {
	return `func test() {
    return
}`
}

func test_arithmetic_operators() string {
	return `func test() {
    let sum = a + b
    let diff = a - b
    let prod = a * b
    let quot = a / b
    let rem = a % b
}`
}

func test_compound_assignment_all() string {
	return `func test() {
    x += 5
    x -= 3
    x *= 2
    x /= 4
    x %= 3
}`
}

func test_increment_decrement_operators() string {
	return `func test() {
    i++
    pos++
    i--
    pos--
}`
}

func test_pre_increment_decrement() string {
	return `func test() {
    let x = ++i
    let y = --j
}`
}

func test_post_increment_decrement() string {
	return `func test() {
    let x = i++
    let y = j--
}`
}

func test_pointer_arithmetic() string {
	return `func test() {
    let next = ptr + 1
    let prev = ptr - 2
}`
}

func test_comparison_operators() string {
	return `func test() {
    let eq = a == b
    let ne = a != b
    let lt = a < b
    let le = a <= b
    let gt = a > b
    let ge = a >= b
}`
}

func test_logical_operators() string {
	return `func test() {
    let and = a && b
    let or = a || b
}`
}

func test_unary_operators() string {
	return `func test() {
    let neg = -value
    let not = !flag
}`
}

func test_address_of_operator() string {
	return `let ptr: *int32 = &value`
}

func test_dereference_operator_read() string {
	return `let x = *ptr`
}

func test_dereference_operator_write() string {
	return `func test() {
    *ptr = 42
}`
}

func test_range_expression() string {
	return `func test() {
    let r = 0..10
}`
}

func test_alloca_single() string {
	return `let ptr = alloca(int32)`
}

func test_alloca_array() string {
	return `let buffer = alloca(byte, 1024)`
}

func test_indexed_pointer_read() string {
	return `func test() {
    let byte_val = buffer[5]
}`
}

func test_indexed_pointer_write() string {
	return `func test() {
    buffer[10] = 0x42
}`
}

func test_indexed_pointer_array() string {
	return `func test() {
    let ptr: *int32 = array_base
    let third_element = ptr[2]
    ptr[3] = 100
}`
}

func test_type_cast() string {
	return `let result = cast<int64>(value)`
}

func test_type_cast_pointer() string {
	return `let byte_ptr = cast<*byte>(int_ptr)`
}

func test_type_cast_pointer_to_int() string {
	return `let addr = cast<uint64>(ptr)`
}

func test_type_cast_int_to_pointer() string {
	return `let new_ptr = cast<*int32>(addr)`
}

func test_type_cast_void_pointer() string {
	return `let generic = cast<*void>(typed_ptr)`
}

func test_function_call() string {
	return `let result = add(5, 10)`
}

func test_intrinsic_sizeof() string {
	return `let sz = sizeof<int32>`
}

func test_intrinsic_sizeof_struct() string {
	return `let st_sz = sizeof<Stat>`
}

func test_intrinsic_alignof() string {
	return `let align = alignof<float64>`
}

func test_intrinsic_memset() string {
	return `func test() {
    let buf = alloca(byte, 1024)
    memset(buf, 0, 1024)
}`
}

func test_intrinsic_memcpy() string {
	return `func test() {
    memcpy(dest_ptr, src_ptr, 1024)
}`
}

func test_intrinsic_memmove() string {
	return `func test() {
    memmove(dest_ptr, src_ptr, 1024)
}`
}

func test_intrinsic_strlen() string {
	return `func test() {
    let cstr: *byte = "hello\0"
    let len = strlen(cstr)
}`
}

func test_intrinsic_memchr() string {
	return `func test() {
    let buf: *byte = "hello\nworld"
    let newline = memchr(buf, '\n', 11)
}`
}

func test_intrinsic_va_start() string {
	return `func printf(fmt: string, ...) {
    let args = va_start(fmt)
}`
}

func test_intrinsic_va_arg() string {
	return `func printf(fmt: string, ...) {
    let args = va_start(fmt)
    let val = va_arg<int32>(args)
}`
}

func test_intrinsic_va_end() string {
	return `func printf(fmt: string, ...) {
    let args = va_start(fmt)
    defer va_end(args)
}`
}

func test_intrinsic_raise() string {
	return `func test() {
    if ptr == null {
        raise("Memory corrupted")
    }
}`
}

func test_intrinsic_memcmp() string {
	return `func test() {
    let diff = memcmp(ptr1, ptr2, 1024)
}`
}

func test_intrinsic_bit_cast() string {
	return `func test() {
    let f: float32 = 1.0
    let bits = bit_cast<uint32>(f)
}`
}

func test_syscall_write() string {
	return `func test() {
    let msg = "Hello, Direct Syscall!\n"
    let len = 23
    let result = syscall(SYS_WRITE, STDOUT, msg, len)
}`
}

func test_char_literal_escape_newline() string {
	return `let newline: char = '\n'`
}

func test_char_literal_escape_tab() string {
	return `let tab: char = '\t'`
}

func test_char_literal_escape_backslash() string {
	return `let backslash: char = '\\'`
}

func test_char_literal_escape_quote() string {
	return `let quote: char = '\''`
}

func test_char_literal_escape_null() string {
	return `let null_char: char = '\0'`
}

func test_string_literal_escapes() string {
	return `func test() {
    let msg: string = "Hello\nWorld"
    let path: string = "C:\\Users\\file"
    let quote: string = "He said \"hello\""
    let tab: string = "Column1\tColumn2"
}`
}

func test_extern_declaration() string {
	return `extern os {
    func printf(*byte, ...) int32
    func sleep(int32) int32
    func usleep(int32) int32
}`
}

func test_extern_with_alias() string {
	return `extern libc {
    func printf "printf" (*byte, ...) int32
    func sleep "usleep" (int32) int32
}`
}

func test_extern_direct_mapping() string {
	return `extern libc {
    func usleep(int32) int32
}`
}

func test_assignment_statement() string {
	return `func test() {
    x = 42
}`
}

func test_complex_expression() string {
	return `func test() {
    let result = (a + b) * c - d / e
}`
}

func test_nested_struct_access() string {
	return `let value = obj.field.subfield`
}

func test_complete_program() string {
	return `namespace main

import "std/io"

struct Point {
    x: int32
    y: int32
}

func add(a: int32, b: int32) int32 {
    return a + b
}

func main() int32 {
    let p = Point{x: 10, y: 20}
    let sum = add(p.x, p.y)
    return sum
}`
}

func test_complete_program_with_class() string {
	return `namespace main

import "std/io"

class Client {
    name: string
    port: int32
    
    func connect(self c: *Client, host: string) bool {
        return true
    }
    
    deinit(self c: *Client) {
    }
}

func main() int32 {
    let c = Client{name: "test", port: 8080}
    c.connect("localhost")
    return 0
}`
}

func test_complete_program_with_for_in() string {
	return `namespace main

import "std/io"

func main() int32 {
    let items: vector<int32> = {1, 2, 3, 4, 5}
    
    for item in items {
        x = item
    }
    
    for i in 0..10 {
        y = i
    }
    
    return 0
}`
}

func test_complete_program_async() string {
	return `namespace main

import "std/io"

async func fetch_data(url: string) string {
    let response = await http_get(url)
    return response
}

async func main() int32 {
    let data = await fetch_data("https://api.example.com")
    return 0
}`
}

func test_complete_program_intrinsics() string {
	return `namespace main

import "std/io"

func main() int32 {
    let buf = alloca(byte, 1024)
    memset(buf, 0, 1024)
    
    let sz = sizeof<int32>
    let align = alignof<float64>
    
    return 0
}`
}

func test_complete_program_mutating() string {
	return `namespace main

struct Counter {
    count: int32
    
    mutating increment(self c: *Counter) {
        c.count++
    }
    
    func get_count(self c: Counter) int32 {
        return c.count
    }
}

func main() int32 {
    let counter = Counter{count: 0}
    counter.increment()
    let value = counter.get_count()
    return value
}`
}

// TestCase represents a single test case
type TestCase struct {
	Name string
	Code func() string
}

func main() {
	// Define all test cases
	testCases := []TestCase{
		// Basic Language Features
		{"Import Single", test_import_single},
		{"Import Multiple", test_import_multiple},
		{"Namespace", test_namespace},
		
		// Variables and Constants
		{"Variable Mutable Typed", test_variable_mutable_typed},
		{"Variable Mutable Inferred", test_variable_mutable_inferred},
		{"Constant Typed", test_constant_typed},
		{"Constant Inferred", test_constant_inferred},
		
		// Basic Types
		{"Signed Integers", test_basic_types_signed_integers},
		{"Unsigned Integers", test_basic_types_unsigned_integers},
		{"USize Type", test_basic_types_usize},
		{"ISize Type", test_basic_types_isize},
		{"Float32 Type", test_basic_types_float32},
		{"Float64 Type", test_basic_types_float64},
		{"Byte Type", test_basic_types_byte},
		{"Bool Type", test_basic_types_bool},
		{"Char Type", test_basic_types_char},
		{"String Type", test_basic_types_string},
		
		// Null Literal
		{"Null Literal", test_null_literal},
		{"Null Check", test_null_check},
		
		// Pointers and References
		{"Pointer Basic", test_pointer_basic},
		{"Pointer Void", test_pointer_void},
		{"Reference Basic", test_reference_basic},
		
		// Collections
		{"Vector Type", test_vector_type},
		{"Vector Literal Empty", test_vector_literal_empty},
		{"Vector Literal With Values", test_vector_literal_with_values},
		{"Vector Literal Inferred", test_vector_literal_inferred},
		{"Map Type", test_map_type},
		{"Map Literal Empty", test_map_literal_empty},
		{"Map Literal With Values", test_map_literal_with_values},
		{"Map Literal Inferred", test_map_literal_inferred},
		
		// Functions
		{"Function Basic", test_function_basic},
		{"Function No Return", test_function_no_return},
		{"Function Void Return", test_function_void_return},
		{"Function Async", test_function_async},
		{"Function Async No Return", test_function_async_no_return},
		
		// Async/Await
		{"Await Expression", test_await_expression},
		{"Await Multiple", test_await_multiple},
		{"Await In If", test_await_in_if},
		
		// Structs
		{"Struct Declaration", test_struct_declaration},
		{"Struct Literal", test_struct_literal},
		{"Struct Literal Type Inference", test_struct_literal_type_inference},
		{"Struct Literal Empty", test_struct_literal_empty},
		{"Struct Field Access", test_struct_field_access},
		{"Struct Field Assignment", test_struct_field_assignment},
		{"Struct Inline Method", test_struct_inline_method},
		{"Struct Mutating Method Inline", test_struct_mutating_method_inline},
		{"Struct Mutating Method Usage", test_struct_mutating_method_usage},
		{"Struct Flat Method", test_struct_flat_method},
		{"Struct Mutating Method Flat", test_struct_mutating_method_flat},
		
		// Classes
		{"Class Declaration", test_class_declaration},
		{"Class Inline Method", test_class_inline_method},
		{"Class Async Method", test_class_async_method},
		{"Class Flat Method", test_class_flat_method},
		{"Class Async Flat Method", test_class_async_flat_method},
		
		// Methods
		{"Method Declaration", test_method_declaration},
		{"Method Call", test_method_call},
		{"Async Method Call", test_async_method_call},
		
		// Control Flow
		{"If Statement", test_if_statement},
		{"If-Else Statement", test_if_else_statement},
		{"If-Else-If Statement", test_if_else_if_statement},
		{"For Loop C-Style", test_for_loop_c_style},
		{"For Loop C-Style Increment", test_for_loop_c_style_increment},
		{"For Loop Condition", test_for_loop_condition},
		{"For Loop Infinite", test_for_loop_infinite},
		{"For-In Loop Vector", test_for_in_loop_vector},
		{"For-In Loop Map", test_for_in_loop_map},
		{"For-In Loop Range", test_for_in_loop_range},
		{"Break Statement", test_break_statement},
		{"Continue Statement", test_continue_statement},
		{"Defer Statement", test_defer_statement},
		{"Return Statement", test_return_statement},
		{"Return Void", test_return_void},
		
		// Operators
		{"Arithmetic Operators", test_arithmetic_operators},
		{"Compound Assignment All", test_compound_assignment_all},
		{"Increment Decrement Operators", test_increment_decrement_operators},
		{"Pre-Increment Decrement", test_pre_increment_decrement},
		{"Post-Increment Decrement", test_post_increment_decrement},
		{"Pointer Arithmetic", test_pointer_arithmetic},
		{"Comparison Operators", test_comparison_operators},
		{"Logical Operators", test_logical_operators},
		{"Unary Operators", test_unary_operators},
		{"Address-Of Operator", test_address_of_operator},
		{"Dereference Operator Read", test_dereference_operator_read},
		{"Dereference Operator Write", test_dereference_operator_write},
		{"Range Expression", test_range_expression},
		
		// Memory and Intrinsics
		{"Alloca Single", test_alloca_single},
		{"Alloca Array", test_alloca_array},
		{"Indexed Pointer Read", test_indexed_pointer_read},
		{"Indexed Pointer Write", test_indexed_pointer_write},
		{"Indexed Pointer Array", test_indexed_pointer_array},
		{"Type Cast", test_type_cast},
		{"Type Cast Pointer", test_type_cast_pointer},
		{"Type Cast Pointer To Int", test_type_cast_pointer_to_int},
		{"Type Cast Int To Pointer", test_type_cast_int_to_pointer},
		{"Type Cast Void Pointer", test_type_cast_void_pointer},
		
		// Intrinsic Functions
		{"Intrinsic sizeof", test_intrinsic_sizeof},
		{"Intrinsic sizeof Struct", test_intrinsic_sizeof_struct},
		{"Intrinsic alignof", test_intrinsic_alignof},
		{"Intrinsic memset", test_intrinsic_memset},
		{"Intrinsic memcpy", test_intrinsic_memcpy},
		{"Intrinsic memmove", test_intrinsic_memmove},
		{"Intrinsic strlen", test_intrinsic_strlen},
		{"Intrinsic memchr", test_intrinsic_memchr},
		{"Intrinsic va_start", test_intrinsic_va_start},
		{"Intrinsic va_arg", test_intrinsic_va_arg},
		{"Intrinsic va_end", test_intrinsic_va_end},
		{"Intrinsic raise", test_intrinsic_raise},
		{"Intrinsic memcmp", test_intrinsic_memcmp},
		{"Intrinsic bit_cast", test_intrinsic_bit_cast},
		
		// Syscalls
		{"Syscall Write", test_syscall_write},
		
		// Character Literals
		{"Char Literal Escape Newline", test_char_literal_escape_newline},
		{"Char Literal Escape Tab", test_char_literal_escape_tab},
		{"Char Literal Escape Backslash", test_char_literal_escape_backslash},
		{"Char Literal Escape Quote", test_char_literal_escape_quote},
		{"Char Literal Escape Null", test_char_literal_escape_null},
		{"String Literal Escapes", test_string_literal_escapes},
		
		// Extern
		{"Extern Declaration", test_extern_declaration},
		{"Extern With Alias", test_extern_with_alias},
		{"Extern Direct Mapping", test_extern_direct_mapping},
		
		// Complex Tests
		{"Function Call", test_function_call},
		{"Assignment Statement", test_assignment_statement},
		{"Complex Expression", test_complex_expression},
		{"Nested Struct Access", test_nested_struct_access},
		
		// Complete Programs
		{"Complete Program", test_complete_program},
		{"Complete Program With Class", test_complete_program_with_class},
		{"Complete Program With For-In", test_complete_program_with_for_in},
		{"Complete Program Async", test_complete_program_async},
		{"Complete Program Intrinsics", test_complete_program_intrinsics},
		{"Complete Program Mutating", test_complete_program_mutating},
	}

	// Run all tests
	passed := 0
	failed := 0

	fmt.Println("Running Arc Language Parser Tests")
	fmt.Println(strings.Repeat("=", 70))

	for i, tc := range testCases {
		code := tc.Code()
		err := ParseArcSource(code)
		
		if err != nil {
			failed++
			fmt.Printf("❌ Test %2d FAILED: %-35s\n", i+1, tc.Name)
			fmt.Printf("   Error: %s\n", err)
			if len(code) < 200 {
				fmt.Printf("   Code: %s\n", strings.ReplaceAll(code, "\n", "\\n"))
			}
		} else {
			passed++
			fmt.Printf("✅ Test %2d PASSED: %-35s\n", i+1, tc.Name)
		}
	}

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("\nResults: %d passed, %d failed out of %d tests\n", 
		passed, failed, len(testCases))
	
	if failed > 0 {
		os.Exit(1)
	}
}