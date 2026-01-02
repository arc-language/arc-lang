struct Point {
    x: int32
    y: int32
}

func main() int32 {
    // Test initialization
    let p = Point{x: 10, y: 20}
    
    // Test read
    if p.x != 10 { return 1 }
    
    // Test write
    p.y = 30
    if p.y != 30 { return 2 }
    
    return 0
}