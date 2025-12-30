

func main() int32 {
    let arr: array<int32, 5> = {10, 20, 30, 40, 50}
    let sum: int32 = 0
    
    // Test for-in loop over array
    for x in arr {
        sum += x
    }
    
    // Expected sum: 10+20+30+40+50 = 150
    if sum == 150 {
        return 0
    }
    return 1
}