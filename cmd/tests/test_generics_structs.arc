
struct Pair<T, U> {
    first: T
    second: U

    func set_first(self p: *Pair<T, U>, val: T) {
        p.first = val
    }

    func get_first(self p: *Pair<T, U>) T {
        return p.first
    }
    
    func set_second(self p: *Pair<T, U>, val: U) {
        p.second = val
    }

    func get_second(self p: *Pair<T, U>) U {
        return p.second
    }
}

func main() int32 {
    // Test 1: Instantiation with identical types (int32, int32)
    let p1: Pair<int32, int32>
    p1.first = 10
    p1.second = 20

    if p1.first != 10 {
        return 1
    }
    if p1.second != 20 {
        return 2
    }

    // Test 2: Method calls on p1
    p1.set_first(50)
    if p1.get_first() != 50 {
        return 3
    }
    
    // Test 3: Instantiation with mixed types (int32, bool)
    let p2: Pair<int32, bool>
    p2.first = 100
    p2.second = true

    if p2.first != 100 {
        return 4
    }
    if !p2.second {
        return 5
    }

    // Test 4: Method calls on p2
    p2.set_second(false)
    if p2.get_second() {
        return 6
    }

    // Final Comprehensive Check
    // We only return 0 if ALL states are exactly as expected.
    // p1.first should be 50 (modified)
    // p1.second should be 20 (original)
    // p2.first should be 100 (original)
    // p2.second should be false (modified)
    if p1.first == 50 && p1.second == 20 && p2.first == 100 && !p2.second {
        return 0
    }

    return 7
}