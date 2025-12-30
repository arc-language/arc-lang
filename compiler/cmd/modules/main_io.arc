namespace main

import "io"

func main() int32 {

    io.println("=== Variadic Test Suite ===\n")
    
    // Test 1: Printf with NO variadic args (should always work)
    io.println("Test 1: Format string only")
    io.printf("Hello, World!\n")
    io.println("  ^ Should see 'Hello, World!' above\n")
    
    // Test 2: Single integer argument
    io.println("Test 2: Single %d")
    io.printf("Value: %d\n", 42)
    io.println("  ^ Should see 'Value: 42' above\n")
    
    // Test 3: Multiple integers
    io.println("Test 3: Multiple %d")
    io.printf("A=%d B=%d C=%d\n", 10, 20, 30)
    io.println("  ^ Should see 'A=10 B=20 C=30' above\n")
    
    // Test 4: Negative number
    io.println("Test 4: Negative %d")
    io.printf("Negative: %d\n", -999)
    io.println("  ^ Should see 'Negative: -999' above\n")
    
    // Test 5: Unsigned (%u)
    io.println("Test 5: Unsigned %u")
    let big: uint32 = 4000000000
    io.printf("Unsigned: %u\n", big)
    io.println("  ^ Should see 'Unsigned: 4000000000' above\n")
    
    // Test 6: Hex (%x)
    io.println("Test 6: Hex %x")
    io.printf("Hex: 0x%x\n", 255)
    io.println("  ^ Should see 'Hex: 0xff' above\n")
    
    // Test 7: Character (%c)
    io.println("Test 7: Character %c")
    io.printf("Char: '%c'\n", 'X')
    io.println("  ^ Should see \"Char: 'X'\" above\n")
    
    // Test 8: String (%s)
    io.println("Test 8: String %s")
    io.printf("Message: %s\n", "Testing")
    io.println("  ^ Should see 'Message: Testing' above\n")
    
    // Test 9: Pointer (%p)
    io.println("Test 9: Pointer %p")
    let x: int32 = 123
    io.printf("Address: %p\n", &x)
    io.println("  ^ Should see 'Address: 0x...' above\n")
    
    // Test 10: Mixed types
    io.println("Test 10: Mixed types")
    io.printf("str=%s int=%d char=%c hex=0x%x\n", "test", 99, 'Z', 255)
    io.println("  ^ Should see all values above\n")
    
    io.println("=== Manual Number Conversion Test ===\n")
    
    // Bypass printf completely to test int_to_str directly
    let buf = alloca(byte, 64)
    let len = io.int_to_str(buf, 42, 10)
    io.print("Direct int_to_str(42): ")
    io.write(io.STDOUT, buf, len)
    io.println("")
    
    io.println("\n=== End Tests ===")
    return 0
}