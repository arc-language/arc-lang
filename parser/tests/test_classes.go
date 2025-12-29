package main

func init() {
	RegisterTest("Class Declaration", func() string {
		return `class Client {
    name: string
    port: int32
}`
	})

	RegisterTest("Class Inline Method", func() string {
		return `class Client {
    name: string
    port: int32
    
    func connect(self c: *Client, host: string) bool {
        return true
    }
    
    deinit(self c: *Client) {
    }
}`
	})

	RegisterTest("Class Async Method", func() string {
		return `class Client {
    name: string
    port: int32
    
    async func fetch_data(self c: *Client) string {
        let response = await http_get("https://example.com")
        return response
    }
}`
	})

	RegisterTest("Class Flat Method", func() string {
		return `class Client {
    name: string
    port: int32
}

func connect(self c: *Client, host: string) bool {
    return true
}

deinit(self c: *Client) {
}`
	})

	RegisterTest("Class Async Flat Method", func() string {
		return `class Client {
    name: string
    port: int32
}

async func fetch_data(self c: *Client) string {
    let response = await http_get("https://example.com")
    return response
}`
	})

	// Method usage tests
	RegisterTest("Method Declaration", func() string {
		return `struct Client {
    port: int32
}

func Connect(self c: *Client, host: string) bool {
    return true
}`
	})

	RegisterTest("Method Call", func() string {
		return `func test() {
    c.Connect("localhost")
}`
	})

	RegisterTest("Async Method Call", func() string {
		return `async func example() {
    let data = await c.fetch_data()
}`
	})
}