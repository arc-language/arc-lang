namespace net

// Syscall constants (Linux x86_64)
const SYS_READ: int32 = 0
const SYS_WRITE: int32 = 1
const SYS_CLOSE: int32 = 3
const SYS_SOCKET: int32 = 41
const SYS_CONNECT: int32 = 42

// Socket configuration
const AF_INET: int32 = 2        // IPv4
const SOCK_STREAM: int32 = 1    // TCP
const IPPROTO_TCP: int32 = 6

// ============================================================================
// Network Helpers
// ============================================================================

// Host to Network Short (Little Endian to Big Endian for x86)
func htons(v: uint16) uint16 {
    let high = (v >> 8) & 0xFF
    let low = v & 0xFF
    return (low << 8) | high
}

// Convert "127.0.0.1" string to uint32 for struct sockaddr_in
// This constructs the integer so that when written to memory on a Little Endian
// CPU, the bytes appear in network order (Big Endian): 127, 0, 0, 1
func inet_addr(ip: string) uint32 {
    let result: uint32 = 0
    let ip_ptr = cast<*byte>(ip)
    
    // Use intrinsic strlen directly
    let len = strlen(ip_ptr)
    
    let current_octet: uint32 = 0
    let shift: int32 = 0
    let i: usize = 0
    
    for i < len {
        let c = ip_ptr[i]
        if c == '.' {
            result = result | (current_octet << cast<uint32>(shift))
            shift += 8
            current_octet = 0
        } else {
            // ASCII '0' is 48
            let digit = cast<uint32>(c) - 48
            current_octet = (current_octet * 10) + digit
        }
        i++
    }
    // Pack last octet
    result = result | (current_octet << cast<uint32>(shift))
    
    return result
}

// ============================================================================
// Socket Class
// ============================================================================

class Socket {
    fd: int32
    connected: bool

    // Manually initialize the socket syscall
    // Returns true on success, false on failure
    func open(self s: *Socket) bool {
        let fd = cast<int32>(syscall(SYS_SOCKET, AF_INET, SOCK_STREAM, IPPROTO_TCP))
        if fd < 0 {
            return false
        }
        s.fd = fd
        s.connected = false
        return true
    }

    func connect(self s: *Socket, ip: string, port: uint16) bool {
        if s.fd < 0 {
            return false
        }

        // Construct sockaddr_in structure manually on the stack
        // struct sockaddr_in {
        //    short   sin_family;   // 2 bytes (Offset 0)
        //    u_short sin_port;     // 2 bytes (Offset 2)
        //    struct  in_addr;      // 4 bytes (Offset 4)
        //    char    sin_zero[8];  // 8 bytes (Offset 8)
        // };
        
        let addr_struct = alloca(byte, 16)
        // Zero out the structure (handles sin_zero padding)
        memset(addr_struct, 0, 16)

        // 1. Set Family (AF_INET = 2)
        let family_ptr = cast<*int16>(addr_struct)
        *family_ptr = cast<int16>(AF_INET)

        // 2. Set Port (Big Endian)
        let port_net = htons(port)
        let port_ptr = cast<*uint16>(addr_struct + 2)
        *port_ptr = port_net

        // 3. Set IP Address (Big Endian bytes)
        let ip_net = inet_addr(ip)
        let ip_ptr = cast<*uint32>(addr_struct + 4)
        *ip_ptr = ip_net

        // 4. Syscall Connect
        let res = cast<int32>(syscall(SYS_CONNECT, s.fd, cast<uint64>(addr_struct), 16))

        if res < 0 {
            return false
        }

        s.connected = true
        return true
    }

    func write(self s: *Socket, data: string) isize {
        if !s.connected { return -1 }
        
        // Use intrinsic strlen directly
        let ptr = cast<*byte>(data)
        let len = strlen(ptr)
        
        return cast<isize>(syscall(SYS_WRITE, s.fd, cast<uint64>(ptr), len))
    }

    func read(self s: *Socket, buffer: *byte, max: usize) isize {
        if !s.connected { return -1 }
        return cast<isize>(syscall(SYS_READ, s.fd, cast<uint64>(buffer), max))
    }

    func close(self s: *Socket) {
        if s.fd >= 0 {
            syscall(SYS_CLOSE, s.fd)
            s.fd = -1
            s.connected = false
        }
    }

    deinit(self s: *Socket) {
        s.close()
    }
}