namespace main

extern libc {
    func printf(*byte, ...) int32
}


struct Counter {
    value: int32
    
    func get(self c: Counter) int32 {
        return c.value
    }
}


func main() int32 {

    let counter = Counter{value: 42}
    let val = counter.get()
    libc.printf("value=%d\n", val)

    return 0
}