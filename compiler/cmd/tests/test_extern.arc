extern libc {
    func abs(int32) int32
}

func main() int32 {
    return abs(-42)
}