namespace main

extern c {
    func printf(*byte, ...) int32
}

func main() int32 {

    // For-in with vector
    let items: vector<int32> = {1, 2, 3, 4, 5}
    for item in items {
        c.printf("%d\n", item)
    }
   
    return 0
}