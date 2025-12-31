func cleanup() void {
    // Cleanup code - we can't observe it easily without side effects
    // but it should execute when main returns
}

func main() int32 {
    defer cleanup()
    return 0
}