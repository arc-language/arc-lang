func main() int32 {
    let ptr = alloca(int32)
    *ptr = 42
    defer {
        *ptr = 0
    }
    let val = *ptr
    if val == 42 {
        return 0
    }
    return 1
}