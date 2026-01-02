// Tests: Variable declaration, type annotation, assignment.
func main() int32 {
    let a: int32 = 10
    let b: int32 = 20
    let c: int32 = 30
    
    // Check if assignments worked
    if a != 10 { return 1 }
    if b != 20 { return 2 }
    if c != 30 { return 3 }
    
    return 0
}