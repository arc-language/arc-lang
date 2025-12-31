func main() int32 {
    let sum = 0
    for i in 0..10 {
        sum = sum + i
    }
    if sum == 45 {
        return 0
    }
    return 1
}