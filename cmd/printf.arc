namespace main

extern libc {
    func printf(*byte, ...) int32
}


struct Person {
    age: int32
    height: int32
    weight: int32
}


func main() int32 {

    let person = Person{age: 30, height: 180, weight: 75}
    libc.printf("age=%d height=%d\n", person.age, person.height)

    return 0
}