namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let large: int32 = 1000
    let small: int8 = cast<int8>(large)
    libc.printf("small=%d\n", small)

    return 0
}