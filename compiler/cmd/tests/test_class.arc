class Client {
    name: string
    port: int32
    
    func get_port(self c: *Client) int32 {
        return c.port
    }
    
    deinit(self c: *Client) {
    }
}

func main() int32 {
    let c = Client{name: "test", port: 8080}
    let port = c.get_port()
    if port == 8080 {
        return 0
    }
    return 1
}