func main() int32 {
    let sum = 0
    for let i = 0; i < 10; i++ {
        if i == 5 {
            continue
        }
        if i == 8 {
            break
        }
        sum = sum + i
    }
    return sum
}