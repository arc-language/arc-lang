namespace main

extern libc {
    func printf(*byte, ...) int32
}

struct Counter {
    count: int32
    
    mutating increment(self c: *Counter) {
        c.count++
    }
}

func main() int32 {

    let counter = Counter{count: 10}
    counter.increment()
    libc.printf("count=%d\n", counter.count)

    return 0
}