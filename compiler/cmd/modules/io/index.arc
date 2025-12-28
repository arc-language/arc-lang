namespace io

// File descriptors
const STDIN: int32 = 0
const STDOUT: int32 = 1
const STDERR: int32 = 2

// Syscall numbers (Linux x86_64)
const SYS_READ: int32 = 0
const SYS_WRITE: int32 = 1

// Internal buffer size for number conversions
const NUM_BUFFER_SIZE: usize = 32

// ============================================================================
// Core Write Functions
// ============================================================================

func write(fd: int32, data: *byte, len: usize) isize {
    return cast<isize>(syscall(SYS_WRITE, fd, cast<uint64>(data), len))
}

func write_string(fd: int32, s: string) isize {
    return write(fd, cast<*byte>(s), cast<usize>(strlen(cast<*byte>(s))))
}

func write_byte(fd: int32, b: byte) isize {
    let buf = alloca(byte, 1)
    buf[0] = b
    return write(fd, buf, 1)
}

// ============================================================================
// Number to String Conversion Helpers
// ============================================================================

func int_to_string(value: int64, buffer: *byte, base: uint32) usize {
    if base < 2 || base > 36 {
        return 0
    }
    
    let digits = "0123456789abcdefghijklmnopqrstuvwxyz"
    let is_negative = value < 0
    let abs_value = if is_negative { -value } else { value }
    let pos: usize = 0
    
    // Convert digits in reverse
    let temp = alloca(byte, NUM_BUFFER_SIZE)
    let temp_pos: usize = 0
    
    if abs_value == 0 {
        temp[temp_pos] = '0'
        temp_pos++
    } else {
        let n = cast<uint64>(abs_value)
        for n > 0 {
            temp[temp_pos] = digits[n % cast<uint64>(base)]
            temp_pos++
            n /= cast<uint64>(base)
        }
    }
    
    // Add negative sign if needed
    if is_negative {
        buffer[pos] = '-'
        pos++
    }
    
    // Reverse the digits into output buffer
    for temp_pos > 0 {
        temp_pos--
        buffer[pos] = temp[temp_pos]
        pos++
    }
    
    return pos
}

func uint_to_string(value: uint64, buffer: *byte, base: uint32) usize {
    if base < 2 || base > 36 {
        return 0
    }
    
    let digits = "0123456789abcdefghijklmnopqrstuvwxyz"
    let pos: usize = 0
    
    // Convert digits in reverse
    let temp = alloca(byte, NUM_BUFFER_SIZE)
    let temp_pos: usize = 0
    
    if value == 0 {
        temp[temp_pos] = '0'
        temp_pos++
    } else {
        let n = value
        for n > 0 {
            temp[temp_pos] = digits[n % cast<uint64>(base)]
            temp_pos++
            n /= cast<uint64>(base)
        }
    }
    
    // Reverse the digits into output buffer
    for temp_pos > 0 {
        temp_pos--
        buffer[pos] = temp[temp_pos]
        pos++
    }
    
    return pos
}

func ptr_to_string(ptr: *void, buffer: *byte) usize {
    let addr = cast<uint64>(ptr)
    
    // Write "0x" prefix
    buffer[0] = '0'
    buffer[1] = 'x'
    let pos: usize = 2
    
    // Convert to hex
    let len = uint_to_string(addr, buffer + pos, 16)
    return pos + len
}

// ============================================================================
// Printf Implementation
// ============================================================================

func printf(fmt: string, ...) int32 {
    let args = va_start(fmt)
    defer va_end(args)
    
    let fmt_ptr = cast<*byte>(fmt)
    let fmt_len = strlen(fmt_ptr)
    let i: usize = 0
    let written: isize = 0
    let num_buffer = alloca(byte, NUM_BUFFER_SIZE)
    
    for i < fmt_len {
        if fmt_ptr[i] == '%' && i + 1 < fmt_len {
            i++
            let specifier = fmt_ptr[i]
            
            if specifier == 'd' {
                // Signed decimal integer
                let val = va_arg<int32>(args)
                let len = int_to_string(cast<int64>(val), num_buffer, 10)
                write(STDOUT, num_buffer, len)
                written += cast<isize>(len)
            } else if specifier == 'i' {
                // Signed decimal integer (same as %d)
                let val = va_arg<int32>(args)
                let len = int_to_string(cast<int64>(val), num_buffer, 10)
                write(STDOUT, num_buffer, len)
                written += cast<isize>(len)
            } else if specifier == 'u' {
                // Unsigned decimal integer
                let val = va_arg<uint32>(args)
                let len = uint_to_string(cast<uint64>(val), num_buffer, 10)
                write(STDOUT, num_buffer, len)
                written += cast<isize>(len)
            } else if specifier == 'x' {
                // Unsigned hexadecimal (lowercase)
                let val = va_arg<uint32>(args)
                let len = uint_to_string(cast<uint64>(val), num_buffer, 16)
                write(STDOUT, num_buffer, len)
                written += cast<isize>(len)
            } else if specifier == 'X' {
                // Unsigned hexadecimal (uppercase) - simplified version
                let val = va_arg<uint32>(args)
                let len = uint_to_string(cast<uint64>(val), num_buffer, 16)
                write(STDOUT, num_buffer, len)
                written += cast<isize>(len)
            } else if specifier == 'p' {
                // Pointer address
                let val = va_arg<*void>(args)
                let len = ptr_to_string(val, num_buffer)
                write(STDOUT, num_buffer, len)
                written += cast<isize>(len)
            } else if specifier == 's' {
                // C-string (null-terminated)
                let val = va_arg<*byte>(args)
                let len = strlen(val)
                write(STDOUT, val, len)
                written += cast<isize>(len)
            } else if specifier == 'c' {
                // Character
                let val = va_arg<int32>(args)
                write_byte(STDOUT, cast<byte>(val))
                written += 1
            } else if specifier == '%' {
                // Literal %
                write_byte(STDOUT, '%')
                written += 1
            }
            
            i++
        } else {
            // Regular character
            write_byte(STDOUT, fmt_ptr[i])
            written += 1
            i++
        }
    }
    
    return cast<int32>(written)
}

// ============================================================================
// Convenience Functions
// ============================================================================

func print(s: string) {
    write_string(STDOUT, s)
}

func println(s: string) {
    write_string(STDOUT, s)
    write_byte(STDOUT, '\n')
}

func eprint(s: string) {
    write_string(STDERR, s)
}

func eprintln(s: string) {
    write_string(STDERR, s)
    write_byte(STDERR, '\n')
}

// ============================================================================
// Read Functions
// ============================================================================

func read(fd: int32, buffer: *byte, count: usize) isize {
    return cast<isize>(syscall(SYS_READ, fd, cast<uint64>(buffer), count))
}

func read_line(buffer: *byte, max_len: usize) isize {
    let total_read: isize = 0
    
    for cast<usize>(total_read) < max_len - 1 {
        let byte_buf = alloca(byte, 1)
        let n = read(STDIN, byte_buf, 1)
        
        if n <= 0 {
            break
        }
        
        if byte_buf[0] == '\n' {
            break
        }
        
        buffer[total_read] = byte_buf[0]
        total_read++
    }
    
    buffer[total_read] = '\0'
    return total_read
}