// --- START OF FILE tests/test_tuple_return.arc ---
func divide(a: int32, b: int32) (int32, bool) {
    if b == 0 {
        return (0, false)
    }
    return (a / b, true)
}

func main() int32 {
    let (result, ok) = divide(10, 2)
    if ok {
        if result == 5 {
            return 0
        }
    }
    return 1
}