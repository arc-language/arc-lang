namespace main

import "io"
import "net"

func main() int32 {
    io.println("--- Standalone TCP Client Test ---")

    // 1. Instantiate
    let client: net.Socket = net.Socket{fd: -1, connected: false}

    // 2. Open
    if !client.open() {
        io.eprintln("Error: Failed to open socket.")
        return 1
    }
    
    // Manual integer printing to avoid printf issues
    io.print("Socket created fd: ")
    let num_buf = alloca(byte, 32)
    let len = io.int_to_string(cast<int64>(client.fd), num_buf, 10)
    io.write(io.STDOUT, num_buf, len)
    io.println("") // Newline

    // Configuration
    let ip = "127.0.0.1"
    let port: uint16 = 8080

    io.print("Connecting to ")
    io.print(ip)
    io.print(":")
    len = io.uint_to_string(cast<uint64>(port), num_buf, 10)
    io.write(io.STDOUT, num_buf, len)
    io.println("...")
    
    io.println("(Ensure python server is running!)")

    // 3. Connect
    if client.connect(ip, port) {
        io.println("Connected!")

        // 4. Write
        let msg = "Hello from Arc!\n"
        let sent = client.write(msg)
        
        if sent > 0 {
            io.print("Sent bytes: ")
            len = io.int_to_string(cast<int64>(sent), num_buf, 10)
            io.write(io.STDOUT, num_buf, len)
            io.println("")
        }

        // 5. Read
        io.print("Waiting for response... ")
        let buf = alloca(byte, 64)
        memset(buf, 0, 64) 

        let n = client.read(buf, 63)
        if n > 0 {
            io.println("\nReceived:")
            io.println(cast<string>(buf))
        } else {
            io.println("\nConnection closed or empty read.")
        }
    } else {
        io.eprintln("Connection failed.")
        return 1
    }

    return 0
}