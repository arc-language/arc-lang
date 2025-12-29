func main() int32 {
    let a = 10
    let b = 5
    let sum = a + b
    let diff = a - b
    let prod = a * b
    let quot = a / b
    let rem = a % b
    let total = sum + diff + prod + quot + rem
    if total == 72 {
        return 0
    }
    return 1
}