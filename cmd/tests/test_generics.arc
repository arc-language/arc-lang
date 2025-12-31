func swap<T>(a: *T, b: *T) {
    let tmp: T = *a
    *a = *b
    *b = tmp
}

func main() int32 {
    let x = 10
    let y = 20
    swap<int32>(&x, &y)
    
    // After swap, x should be 20 and y should be 10
    if x == 20 {
        return 0  // Success!
    }
    return 1  // Failure - swap didn't work
}