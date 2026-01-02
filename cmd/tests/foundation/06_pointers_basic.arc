// Tests: Pointer creation and dereferencing.
func main() int32 {
    let val: int32 = 123
    let ptr: *int32 = &val
    
    // Read via pointer
    if *ptr != 123 {
        return 1
    }
    
    // Write via pointer
    *ptr = 456
    
    // Check original variable modified
    if val != 456 {
        return 2
    }
    
    return 0
}