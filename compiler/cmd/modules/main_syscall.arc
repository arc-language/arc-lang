import "syscalls"

func main() {
    let msg = "Hello from syscall!\n"
    let len = 20
    
    // fd 1 is STDOUT
    let (n, err) = syscalls.Write(1, msg, len)
    
    if err != 0 {
        // Handle error (errno is set)
        return
    }
    
    // n contains bytes written
}