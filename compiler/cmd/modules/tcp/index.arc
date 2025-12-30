namespace tcp

import "../io"

// ============================================================================
// Constants (Linux x86_64)
// ============================================================================

// Syscall Numbers
const SYS_READ: int32 = 0
const SYS_WRITE: int32 = 1
const SYS_CLOSE: int32 = 3
const SYS_SOCKET: int32 = 41
const SYS_CONNECT: int32 = 42

// Socket Configuration
const AF_INET: int32 = 2        // IPv4
const SOCK_STREAM: int32 = 1    // TCP
const IPPROTO_TCP: int32 = 6

// ============================================================================
// Internal Helpers
// ============================================================================

// Host to Network Short (Little Endian to Big Endian swap)
func htons(v: uint16) uint16 {
    // Extract high and low bytes
    let high = (v >> 8) & 0xFF
    let low = v & 0xFF
    // Swap them
    return (low << 8) | high
}

// Parse IPv4 string "127.0.0.1" to raw uint32 (Network Byte Order/Big Endian)
// Note: This constructs the integer so that when written to memory on a 
// Little Endian CPU, the bytes map correctly to the network structure.
func inet_addr(ip: string) uint32 {
    let result: uint32 = 0
    let ip_ptr = cast<*byte>(ip)
    let len = strlen(ip_ptr)
    
    let current_octet: uint32 = 0
    let shift: int32 = 0
    let i: usize = 0
    
    for i < len {
        let c = ip_ptr[i]
        
        if c == '.' {
            // Pack current octet into result
            result = result | (current_octet << cast<uint32>(shift))
            shift += 8
            current_octet = 0
        } else {
            // Parse digit (ASCII '0' is 48)
            let digit = cast<uint32>(c) - 48
            current_octet = (current_octet * 10) + digit
        }
        i++
    }
    
    // Pack the final octet
    result = result | (current_octet << cast<uint32>(shift))
    
    return result
}

// ============================================================================
// TCP Client Class
// ============================================================================

class Client {
    fd: int32
    connected: bool
    host: string
    port: uint16

    // Send data to the connected server
    func write(self c: *Client, data: string) isize {

        io.printf("writing to socket\n")

        //if !c.connected { 
        //    return -1 
        //}
        
        let ptr = cast<*byte>(data)
        let len = strlen(ptr)

        io.printf("writing to socket len %d\n", len)

        
        // syscall(SYS_WRITE, fd, buf, count)
        let res = syscall(SYS_WRITE, c.fd, cast<uint64>(ptr), len)
        return cast<isize>(res)
    }

    // Read data into a buffer
    func read(self c: *Client, buffer: *byte, max_len: usize) isize {
        if !c.connected { 
            return -1 
        }

        // syscall(SYS_READ, fd, buf, count)
        let res = syscall(SYS_READ, c.fd, cast<uint64>(buffer), max_len)
        return cast<isize>(res)
    }

    // Close the connection manually
    func close(self c: *Client) {
        if c.fd >= 0 {
            syscall(SYS_CLOSE, c.fd)
            c.fd = -1
            c.connected = false
        }
    }

    // Cleanup when reference count hits 0
    deinit(self c: *Client) {
        c.close()
    }
}

// ============================================================================
// Factory Function
// ============================================================================

func NewClient(host: string, port: uint16) Client {
    // 1. Create the socket file descriptor
    // syscall(SYS_SOCKET, domain, type, protocol)
    let raw_fd = syscall(SYS_SOCKET, AF_INET, SOCK_STREAM, IPPROTO_TCP)
    let fd = cast<int32>(raw_fd)

    if fd < 0 {
        throw "Failed to create socket"
    }

    // 2. Prepare sockaddr_in structure on the stack
    // struct sockaddr_in {
    //    short   sin_family;   // 2 bytes
    //    u_short sin_port;     // 2 bytes
    //    struct  in_addr;      // 4 bytes
    //    char    sin_zero[8];  // 8 bytes (padding)
    // };
    let addr_struct = alloca(byte, 16)
    memset(addr_struct, 0, 16)

    // Set Family (AF_INET = 2)
    let family_ptr = cast<*int16>(addr_struct)
    *family_ptr = cast<int16>(AF_INET)

    // Set Port (Big Endian)
    let port_net = htons(port)
    let port_ptr = cast<*uint16>(addr_struct + 2)
    *port_ptr = port_net

    // Set IP Address (Big Endian)
    let ip_net = inet_addr(host)
    let ip_ptr = cast<*uint32>(addr_struct + 4)
    *ip_ptr = ip_net

    // 3. Connect
    // syscall(SYS_CONNECT, fd, addr, addrlen)
    let res = syscall(SYS_CONNECT, fd, cast<uint64>(addr_struct), 16)
    
    if cast<int32>(res) < 0 {
        syscall(SYS_CLOSE, fd)
        throw "Failed to connect to host"
    }

    // 4. Return initialized class instance
    // Note: 'Client' is a reference type, so this returns the reference.
    return Client{
        fd: fd, 
        connected: true,
        host: host,
        port: port
    }
}



// ============================================================================
// Syscall Constants (Linux x86_64) - Server Additions
// ============================================================================

const SYS_ACCEPT: int32 = 43
const SYS_BIND: int32 = 49
const SYS_LISTEN: int32 = 50
const SYS_SETSOCKOPT: int32 = 54

const SOL_SOCKET: int32 = 1
const SO_REUSEADDR: int32 = 2

// ============================================================================
// Server Class
// ============================================================================

class Server {
    fd: int32
    port: uint16
    running: bool

    // Enable SO_REUSEADDR so you don't get "Address already in use" 
    // errors when restarting the server quickly.
    func set_reuse_addr(self s: *Server) {
        let val: int32 = 1
        // setsockopt(fd, level, optname, optval, optlen)
        syscall(SYS_SETSOCKOPT, s.fd, SOL_SOCKET, SO_REUSEADDR, &val, 4)
    }

    func bind(self s: *Server, port: uint16) {
        s.port = port
        
        // Prepare sockaddr_in (16 bytes)
        let addr = alloca(byte, 16)
        memset(addr, 0, 16)

        // Family: AF_INET (2)
        let family_ptr = cast<*int16>(addr)
        *family_ptr = cast<int16>(AF_INET)

        // Port: Network Byte Order
        let port_net = htons(port)
        let port_ptr = cast<*uint16>(addr + 2)
        *port_ptr = port_net

        // IP: 0.0.0.0 (INADDR_ANY) - Listen on all interfaces
        // We leave the 4 bytes at offset 4 as 0 (memset handled this)

        // syscall(SYS_BIND, fd, addr, addrlen)
        let res = cast<int32>(syscall(SYS_BIND, s.fd, cast<uint64>(addr), 16))
        
        if res < 0 {
            throw "Failed to bind to port"
        }
    }

    func listen(self s: *Server, backlog: int32) {
        // syscall(SYS_LISTEN, fd, backlog)
        let res = cast<int32>(syscall(SYS_LISTEN, s.fd, backlog))
        if res < 0 {
            throw "Failed to listen on socket"
        }
        s.running = true
    }

    // Returns a new Client instance representing the incoming connection
    func accept(self s: *Server) int32 {
        if !s.running {
            throw "Server is not running"
        }

        // We need buffers to store the connecting client's info
        let cli_addr = alloca(byte, 16)
        let cli_len = alloca(uint32)
        *cli_len = 16

        // syscall(SYS_ACCEPT, fd, addr, addrlen_ptr)
        // Returns the File Descriptor for the NEW connection
        let client_fd = cast<int32>(syscall(SYS_ACCEPT, s.fd, cast<uint64>(cli_addr), cast<uint64>(cli_len)))
        return client_fd
    }

    func close(self s: *Server) {
        if s.fd >= 0 {
            syscall(SYS_CLOSE, s.fd)
            s.fd = -1
            s.running = false
        }
    }

    deinit(self s: *Server) {
        s.close()
    }
}

// ============================================================================
// Factory Function
// ============================================================================

func NewServer(port: uint16) Server {
    let raw_fd = syscall(SYS_SOCKET, AF_INET, SOCK_STREAM, IPPROTO_TCP)
    let fd = cast<int32>(raw_fd)

    if fd < 0 {
        throw "Failed to create socket"
    }

    let server = Server{
        fd: fd,
        port: port,
        running: false
    }

    // Setup default behavior
    server.set_reuse_addr()
    server.bind(port)
    
    return server
}