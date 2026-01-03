namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let i: int32 = 42
    let f: float32 = cast<float32>(i)
    libc.printf("float=%.1f\n", f)

    return 0
}