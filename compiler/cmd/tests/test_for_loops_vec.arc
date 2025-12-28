namespace main

extern c {
    func printf(*byte, ...) int32
}

func main() int32 {
    let items: vector<int32> = {1, 2, 3, 4, 5}
    
    c.printf("Before loop\n")
    for item in items {
        c.printf("In loop\n")  // Simple test - no format args
    }
    c.printf("After loop\n")
    
    return 0
}