func main() int32 {
    let status = 2
    let result = 0
    switch status {
        case 0:
            result = 10
        case 1:
            result = 20
        case 2:
            result = 30
        default:
            result = 40
    }
    if result == 30 {
        return 0
    }
    return 1
}