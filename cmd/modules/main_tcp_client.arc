namespace main

import "io"
import "tcp" 


func run_client() {

    // Factory pattern usage
    let client = tcp.NewClient("127.0.0.1", 8080)
    
    io.printf("Connected to server! \n")
    
    for x in 1..10 {
        // Write data
        client.write("GET / HTTP/1.1\r\n\r\n")
    }
    
    // Read response
    let buffer = alloca(byte, 4096)
    let bytes_read = client.read(buffer, 4096)
    
    if bytes_read > 0 {
        io.printf("Received data\n")
    }
    
    // Client is automatically closed via deinit when it goes out of scope,
    // or can be closed manually:
    client.close()
    io.printf("disconnect\n")
}


func main() {

    run_client()
    
}