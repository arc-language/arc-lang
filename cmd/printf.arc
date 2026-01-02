namespace main 

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let x = 3
    if x > 5 {
        libc.printf("greater\n")
    } else {
        libc.printf("not_greater\n")
    }

    return 0
}