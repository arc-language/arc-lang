import "io"

namespace main

func main() int32 {
    // Basic printing
    io.print("Hello, ")
    io.println("World!")
    
    // Printf with various format specifiers
    io.printf("Integer: %d\n", 42)
    io.printf("Unsigned: %u\n", 4294967295)
    io.printf("Hex: 0x%x\n", 255)
    io.printf("String: %s\n", "Hello")
    io.printf("Char: %c\n", 'A')
    
    let value: int32 = 100
    let ptr: *int32 = &value
    io.printf("Pointer: %p\n", ptr)
    
    // Error output
    io.eprintln("This goes to stderr")
    
    // Reading input
    let buffer = alloca(byte, 256)
    io.print("Enter text: ")
    io.read_line(buffer, 256)
    io.printf("You entered: %s\n", buffer)
    
    return 0
}