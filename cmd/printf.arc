namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let value: int32 = 77
    let ptr: *int32 = &value
    let pptr: **int32 = &ptr
    libc.printf("value=%d\n", **pptr)

    return 0
}