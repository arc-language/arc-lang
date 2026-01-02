namespace main 

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {

    libc.printf("Ok\n")

    return 0
}