func main() int32 {
    let x: int64 = 42
    let y = cast<int32>(x)
    if y == 42 {
        return 0
    }
    return 1
}