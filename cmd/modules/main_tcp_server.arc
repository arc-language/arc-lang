namespace main

import "io"
import "tcp"

func main() {

    let port: uint16 = 8080
    
    let server = tcp.NewServer(port)
    server.listen(10) // Backlog of 10 connections
    
    io.printf("HTTP Server listening on http://localhost:%d\n", cast<int32>(port))
    
    // Simple Loop
    for {
        // Block until a client connects
        let client = server.accept()

        if client > 0 {
            io.printf("got new client\n")
        }
    }
    
}