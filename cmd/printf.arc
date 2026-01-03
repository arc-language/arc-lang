namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let arr = alloca<int32>(5)
    arr[0] = 10
    arr[1] = 20
    let ptr: *int32 = arr
    let next = ptr + 1
    libc.printf("value=%d\n", *next)

    return 0
}