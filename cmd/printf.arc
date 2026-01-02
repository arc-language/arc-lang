namespace main 

extern libc {
    func printf(*byte, ...) int32
}

func get_value() int32 {
    return 42
}

func main() int32 {

    let x = get_value()
    libc.printf("value=%d\n", x)

    return 0
}