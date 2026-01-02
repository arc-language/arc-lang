// Tests: Basic counting loop.
func main() int32 {
    let sum: int32 = 0
    let i: int32 = 0
    
    // Standard C-style loop logic (using assignment if ++ not ready)
    for i = 0; i < 10; i = i + 1 {
        sum = sum + 1
    }
    
    if sum == 10 {
        return 0
    }
    return 1
}