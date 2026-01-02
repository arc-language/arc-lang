// Tests: Basic math operations.
func main() int32 {
    let a: int32 = 10
    let b: int32 = 2
    
    let sum = a + b      // 12
    let diff = a - b     // 8
    let prod = a * b     // 20
    let quot = a / b     // 5
    
    if sum != 12 { return 1 }
    if diff != 8 { return 2 }
    if prod != 20 { return 3 }
    if quot != 5 { return 4 }
    
    return 0
}