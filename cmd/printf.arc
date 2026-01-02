namespace main 

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let x = 2
    switch x {
        case 1:
            libc.printf("one\n")
        case 2:
            libc.printf("two\n")
        case 3:
            libc.printf("three\n")
    }

    return 0
}