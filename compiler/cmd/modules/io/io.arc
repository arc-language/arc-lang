namespace io

// ============================================================================
// System Constants (Linux x86_64)
// ============================================================================

const STDIN: int32 = 0
const STDOUT: int32 = 1
const STDERR: int32 = 2

const SYS_READ: int64 = 0
const SYS_WRITE: int64 = 1

// Buffer size for number conversions (sufficient for 64-bit int + sign + null)
const NUM_BUF_SIZE: usize = 64

// ============================================================================
// Low-Level System Calls
// ============================================================================

// write writes 'count' bytes from 'buf' to file descriptor 'fd'.
func write(fd: int32, buf: *byte, count: usize) isize {
    // syscall(number, arg1, arg2, arg3...)
    return cast<isize>(syscall(SYS_WRITE, fd, cast<uint64>(buf), count))
}

// read reads up to 'count' bytes from 'fd' into 'buf'.
func read(fd: int32, buf: *byte, count: usize) isize {
    return cast<isize>(syscall(SYS_READ, fd, cast<uint64>(buf), count))
}

// ============================================================================
// Primitive Output Helpers
// ============================================================================

// write_byte writes a single byte to the file descriptor.
func write_byte(fd: int32, c: byte) {
    let buf = alloca(byte, 1)
    buf[0] = c
    write(fd, buf, 1)
}

// write_str writes a null-terminated string to the file descriptor.
func write_str(fd: int32, s: string) {
    let ptr = cast<*byte>(s)
    let len = strlen(ptr)
    write(fd, ptr, len)
}

// print writes a string to STDOUT.
func print(s: string) {
    write_str(STDOUT, s)
}

// println writes a string followed by a newline to STDOUT.
func println(s: string) {
    write_str(STDOUT, s)
    write_byte(STDOUT, '\n')
}

// ============================================================================
// Number to String Conversion (No Libc)
// ============================================================================

// reverse_in_place reverses a byte buffer.
func reverse_in_place(buf: *byte, len: usize) {
    if len == 0 { return }
    
    let left: usize = 0
    let right: usize = len - 1
    
    for left < right {
        let temp = buf[left]
        buf[left] = buf[right]
        buf[right] = temp
        left++
        right--
    }
}

// int_to_str converts a signed integer to a string in the given base.
// Returns the number of bytes written.
func int_to_str(buf: *byte, val: int64, base: int64) usize {
    if val == 0 {
        buf[0] = '0'
        return 1
    }

    let i: usize = 0
    let is_neg = false
    let num = val

    // Handle negative numbers for base 10
    if val < 0 && base == 10 {
        is_neg = true
        num = -val
    } else if val < 0 {
        // For other bases, usually treat as unsigned, but here we just abs
        num = -val
    }

    // Extract digits
    for num > 0 {
        let rem = num % base
        if rem > 9 {
            // ASCII 'a' = 97. rem=10 -> 'a'
            buf[i] = cast<byte>((rem - 10) + 97) 
        } else {
            // ASCII '0' = 48
            buf[i] = cast<byte>(rem + 48)
        }
        num /= base
        i++
    }

    if is_neg {
        buf[i] = '-'
        i++
    }

    // Numbers are generated backward, reverse them
    reverse_in_place(buf, i)
    return i
}

// uint_to_str converts an unsigned integer to a string.
func uint_to_str(buf: *byte, val: uint64, base: uint64) usize {
    if val == 0 {
        buf[0] = '0'
        return 1
    }

    let i: usize = 0
    let num = val

    for num > 0 {
        let rem = num % base
        if rem > 9 {
            buf[i] = cast<byte>((rem - 10) + 97)
        } else {
            buf[i] = cast<byte>(rem + 48)
        }
        num /= base
        i++
    }

    reverse_in_place(buf, i)
    return i
}

// ============================================================================
// Printf Implementation
// ============================================================================

// printf formats and prints data to STDOUT.
// Supported specifiers:
// %s - string
// %c - char
// %d - int32
// %u - uint32
// %x - hex (uint32)
// %p - pointer
func printf(fmt: string, ...) {
    let args = va_start(fmt)
    
    // Manual defer simulation (since defer implies function scoping)
    // We will just call va_end at the end.

    let ptr = cast<*byte>(fmt)
    let len = strlen(ptr)
    let i: usize = 0

    // Stack buffer for converting numbers
    let num_buf = alloca(byte, NUM_BUF_SIZE)

    for i < len {
        let c = ptr[i]

        if c == '%' {
            i++
            if i >= len { break }
            
            let spec = ptr[i]
            
            switch spec {
                case 's':
                    let s_arg = va_arg<*byte>(args)
                    let s_len = strlen(s_arg)
                    write(STDOUT, s_arg, s_len)
                
                case 'c':
                    // In variadic calls, char/short promote to int
                    let c_arg = va_arg<int32>(args)
                    write_byte(STDOUT, cast<byte>(c_arg))
                
                case 'd':
                    let d_arg = va_arg<int32>(args)
                    let wrote = int_to_str(num_buf, cast<int64>(d_arg), 10)
                    write(STDOUT, num_buf, wrote)
                
                case 'u':
                    let u_arg = va_arg<uint32>(args)
                    let wrote = uint_to_str(num_buf, cast<uint64>(u_arg), 10)
                    write(STDOUT, num_buf, wrote)
                
                case 'x':
                    let x_arg = va_arg<uint32>(args)
                    let wrote = uint_to_str(num_buf, cast<uint64>(x_arg), 16)
                    write(STDOUT, num_buf, wrote)
                
                case 'p':
                    let p_arg = va_arg<*void>(args)
                    write_str(STDOUT, "0x")
                    let wrote = uint_to_str(num_buf, cast<uint64>(p_arg), 16)
                    write(STDOUT, num_buf, wrote)
                
                case '%':
                    write_byte(STDOUT, '%')
                
                default:
                    // Unknown specifier, print raw
                    write_byte(STDOUT, '%')
                    write_byte(STDOUT, spec)
            }
        } else {
            write_byte(STDOUT, c)
        }
        i++
    }

    va_end(args)
}

// ============================================================================
// Input Utilities
// ============================================================================

// read_line reads from STDIN until newline or buffer full.
// Adds null terminator. Returns bytes read.
func read_line(buf: *byte, max_size: usize) usize {
    let count: usize = 0
    let temp = alloca(byte, 1)

    // Leave room for null terminator
    for count < (max_size - 1) {
        let n = read(STDIN, temp, 1)
        
        // EOF or Error
        if n <= 0 {
            break
        }

        let c = temp[0]
        
        if c == '\n' {
            break
        }

        buf[count] = c
        count++
    }

    buf[count] = '\0'
    return count
}