func main() int32 {
    let ptr = alloca(int32)
    *ptr = 42
    return *ptr
}