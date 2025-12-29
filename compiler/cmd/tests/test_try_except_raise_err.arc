// NOTE: raise() is a low-level intrinsic that CANNOT be caught by except blocks
// This test demonstrates that raise() aborts the program
// For catchable exceptions, use 'throw' instead

func divide(a: int32, b: int32) int32 {
    if b == 0 {
        raise("Division by zero")  // This will abort the program!
    }
    return a / b
}

func main() int32 {
    try {
        let result = divide(10, 0)  // This triggers raise()
        return -1  // Won't reach here
    } except err {
        // raise() CANNOT be caught - this block will never execute
        return -1  // Won't reach here either
    }
    // Program will abort before reaching this
    return -1
}