enum Status {
    OK = 0
    ERROR = 1
    PENDING = 2
}

func main() int32 {
    let s = Status.OK
    return cast<int32>(s)
}