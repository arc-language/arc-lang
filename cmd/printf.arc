namespace main 

extern libc {
    func printf(*byte, ...) int32
}

struct Rectangle {
    width: int32
    height: int32
}

func main() int32 {

    let rect = Rectangle{width: 100, height: 50}
    libc.printf("width=%d\n", rect.width)

    return 0
}