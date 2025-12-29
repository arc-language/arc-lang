func divide(a: int32, b: int32) int32 {
    if b == 0 {
        throw "Division by zero"  // Throws catchable exception
    }
    return a / b
}

func main() int32 {
    try {
        let result = divide(10, 0)  // This triggers throw
        return -1  // Should not reach here - exception thrown
    } except err {
        // Successfully caught the exception!
        return 0  // Test passes - exception was caught
    }
}