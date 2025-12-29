func swap<T>(a: *T, b: *T) {
    let tmp: T = *a
    *a = *b
    *b = tmp
}

func main() int32 {
    let x = 10
    let y = 20
    swap<int32>(&x, &y)
    return x
}