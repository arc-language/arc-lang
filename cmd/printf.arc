namespace main 

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    let score = 75
    if score >= 90 {
        libc.printf("A\n")
    } else if score >= 80 {
        libc.printf("B\n")
    } else if score >= 70 {
        libc.printf("C\n")
    } else {
        libc.printf("F\n")
    }

    return 0
}