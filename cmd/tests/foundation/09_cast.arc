func main() int32 {
    let large: int64 = 100
    
    // Cast down
    let small = cast<int32>(large)
    
    if small != 100 {
        return 1
    }
    
    // Cast up
    let back = cast<int64>(small)
    if back != 100 {
        return 2
    }
    
    return 0
}