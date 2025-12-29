extern libc {
    func abs(int32) int32
}

func main() int32 {
    let result = abs(-42)
    if result == 42 {
        return 0
    }
    return 1
}