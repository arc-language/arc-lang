func main() int32 {
    let x = 42
    let ptr: *int32 = &x
    let val = *ptr
    if val == 42 {
        return 0
    }
    return 1
}