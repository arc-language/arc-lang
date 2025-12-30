namespace syscall

// Standard syscall number for write on Linux x86_64
const SYS_WRITE: usize = 1

// Write attempts to write 'count' bytes from 'buf' to file descriptor 'fd'.
// Returns a tuple: (bytes_written, errno).
// If errno is 0, the operation was successful.
func Write(fd: int32, buf: *byte, count: usize) (isize, int32) {
    // syscall() is a compiler intrinsic
    let ret = syscall(SYS_WRITE, fd, buf, count)
    
    // Linux kernel returns errors as negative values between -4095 and -1
    if ret < 0 && ret > -4096 {
        // Return (-1, errno)
        return (-1, cast<int32>(-ret))
    }
    
    // Success: Return (bytes_written, 0)
    return (ret, 0)
}