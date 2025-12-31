func main() int32 {
    let ptr = alloca(int32)
    *ptr = 42
    let val = *ptr
    if val == 42 {
        return 0
    }
    return 1
}