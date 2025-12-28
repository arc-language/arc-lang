namespace main

import "io"
import "net"

func main() int32 {
    io.println("--- Standalone TCP Client Test ---")

    // 1. Instantiate the class
    // Note: Compiler handles heap allocation for 'class' types via struct literal syntax
    let client: net.Socket = net.Socket{fd: -1, connected: false}

    // 2. Open the socket file descriptor
    if !client.open() {
        io.eprintln("Error: Failed to open socket.")
        return 1
    }
    io.printf("Socket created (fd: %d)\n", client.fd)

    // Configuration
    let ip = "127.0.0.1"
    let port: uint16 = 8080

    io.printf("Connecting to %s:%d...\n", ip, cast<uint32>(port))
    io.println("(Ensure 'nc -l 8080' is running!)")

    // 3. Connect
    if client.connect(ip, port) {
        io.println("Connected!")

        // 4. Write
        let msg = "Hello from Arc!\n"
        let sent = client.write(msg)
        if sent > 0 {
            io.printf("Sent %d bytes.\n", cast<int32>(sent))
        }

        // 5. Read
        io.print("Waiting for response... ")
        let buf = alloca(byte, 64)
        memset(buf, 0, 64) // Zero buffer for safety

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

    // client.deinit() is called automatically here
    return 0
}