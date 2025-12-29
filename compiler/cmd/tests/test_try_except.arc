func divide(a: int32, b: int32) int32 {
    if b == 0 {
        raise("Division by zero")
    }
    return a / b
}

func main() int32 {
    try {
        let result = divide(10, 2)
        return result
    } except err {
        return -1
    }
}