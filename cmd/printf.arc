namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let f: float32 = 3.9
    let i: int32 = cast<int32>(f)
    libc.printf("int=%d\n", i)

    return 0
}