namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let arr = alloca<int32>(3)
    arr[0] = 100
    arr[1] = 200
    arr[2] = 300
    libc.printf("arr[1]=%d\n", arr[1])

    return 0
}