func main() int32 {
    let x = 42
    let ptr: *int32 = &x
    let val = *ptr
    return val
}