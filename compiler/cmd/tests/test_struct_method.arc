struct Point {
    x: int32
    y: int32
    
    func sum(self p: Point) int32 {
        return p.x + p.y
    }
}

func main() int32 {
    let p = Point{x: 10, y: 20}
    return p.sum()
}