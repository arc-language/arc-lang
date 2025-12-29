func divide(a: int32, b: int32) int32 {
    if b == 0 {
        raise("Division by zero")  // Won't be hit in this test
    }
    return a / b
}

func main() int32 {
    try {
        let result = divide(10, 2)  // Normal execution, no error
        return 0  // Success
    } except err {
        return -1  // Should not reach here
    }
}