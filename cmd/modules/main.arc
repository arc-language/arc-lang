namespace main

import "io"

func main() int32 {

    for x in 1..10 {

        if x == 5 {
            io.printf("breaking\n")
            break
        }

        io.printf("%d\n", x)
    }

    return 0
}