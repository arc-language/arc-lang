func main() int32 {
    let sz = sizeof<int32>
    let size_val = cast<int32>(sz)
    if size_val == 4 {
        return 0
    }
    return 1
}