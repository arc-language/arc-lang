// --- START OF FILE tests/test_struct_method.arc ---
struct Point {
    x: int32
    y: int32
    
    func sum(self p: Point) int32 {
        return p.x + p.y
    }
}

func main() int32 {
    let p = Point{x: 10, y: 20}
    let s = p.sum()
    if s == 30 {
        return 0
    }
    return 1
}