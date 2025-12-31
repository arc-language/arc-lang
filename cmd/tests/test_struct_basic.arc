struct Point {
    x: int32
    y: int32
}

func main() int32 {
    let p = Point{x: 10, y: 20}
    let sum = p.x + p.y
    if sum == 30 {
        return 0
    }
    return 1
}