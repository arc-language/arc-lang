namespace main

import "io"

func main() int32 {
    // 1. Basic String Output
    io.println("=== IO Library Test ===")
    io.print("1. Basic Print: ")
    io.println("Success")

    // 2. Integers (Signed)
    io.println("\n2. Signed Integers:")
    io.printf("   Positive: %d\n", 42)
    io.printf("   Negative: %d\n", -123456)
    io.printf("   Zero:     %d\n", 0)

    // 3. Unsigned & Hex
    io.println("\n3. Unsigned & Hex:")
    let u_val: uint32 = 3000000000
    io.printf("   Unsigned: %u\n", u_val)
    io.printf("   Hex (255): %x\n", 255)
    io.printf("   Hex (Large): %x\n", 3735928559) // 0xDEADBEEF

    // 4. Characters & Strings
    io.println("\n4. Chars & Strings:")
    io.printf("   Char: %c\n", 'A')
    io.printf("   String: %s\n", "Hello inside printf")

    // 5. Pointers
    io.println("\n5. Pointers:")
    let ptr: *void = null
    io.printf("   Null: %p\n", ptr)
    
    let local_var: int32 = 100
    io.printf("   Stack Addr: %p\n", &local_var)

    // 6. Direct File Descriptor Access (Stderr)
    io.println("\n6. Testing STDERR:")
    let err_msg = "   [STDERR] This is a direct write to fd 2\n"
    // Note: We cast string to *byte manually for the low-level write
    io.write(io.STDERR, cast<*byte>(err_msg), 43)

    io.println("\n=== All Tests Passed ===")
    return 0
}