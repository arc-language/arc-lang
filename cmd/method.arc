// example_v5.arc
namespace main

extern io {
    // printf is variadic, denoted by ...
    func printf(*byte, ...) int32
    func puts(*byte) int32
}


struct Logging {
    s: int32
}

mutating run_logging(self c: *Logging) {
    c.s = 1000
}


func main() int32 {

    let w = Logging{}

    w.run_logging()

    io.printf("it's working %d\n", w.s)

    return 0
}