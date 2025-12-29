func main() int32 {
    let ptr = alloca(int32)
    defer {
        *ptr = 0
    }
    *ptr = 42
    return *ptr
}