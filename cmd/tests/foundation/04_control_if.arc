// Tests: if, nested if, comparison operators.
func main() int32 {
    let x: int32 = 42
    
    if x == 42 {
        if x > 40 {
            return 0 // Success
        }
    }
    
    return 1 // Failure
}