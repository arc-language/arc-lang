package main

func init() {
	RegisterTest("Function Call", func() string {
		return `let result = add(5, 10)`
	})

	RegisterTest("Assignment Statement", func() string {
		return `func test() {
    x = 42
}`
	})

	RegisterTest("Complex Expression", func() string {
		return `func test() {
    let result = (a + b) * c - d / e
}`
	})

	RegisterTest("Nested Struct Access", func() string {
		return `let value = obj.field.subfield`
	})

	RegisterTest("Complete Program", func() string {
		return `namespace main

import "std/io"

struct Point {
    x: int32
    y: int32
}

func add(a: int32, b: int32) int32 {
    return a + b
}

func main() int32 {
    let p = Point{x: 10, y: 20}
    let sum = add(p.x, p.y)
    return sum
}`
	})

	RegisterTest("Complete Program With Class", func() string {
		return `namespace main

import "std/io"

class Client {
    name: string
    port: int32
    
    func connect(self c: *Client, host: string) bool {
        return true
    }
    
    deinit(self c: *Client) {
    }
}

func main() int32 {
    let c = Client{name: "test", port: 8080}
    c.connect("localhost")
    return 0
}`
	})

	RegisterTest("Complete Program With For-In", func() string {
		return `namespace main

import "std/io"

func main() int32 {
    let items: vector<int32> = {1, 2, 3, 4, 5}
    
    for item in items {
        x = item
    }
    
    for i in 0..10 {
        y = i
    }
    
    return 0
}`
	})

	RegisterTest("Complete Program Async", func() string {
		return `namespace main

import "std/io"

async func fetch_data(url: string) string {
    let response = await http_get(url)
    return response
}

async func main() int32 {
    let data = await fetch_data("https://api.example.com")
    return 0
}`
	})

	RegisterTest("Complete Program Intrinsics", func() string {
		return `namespace main

import "std/io"

func main() int32 {
    let buf = alloca(byte, 1024)
    memset(buf, 0, 1024)
    
    let sz = sizeof<int32>
    let align = alignof<float64>
    
    return 0
}`
	})

	RegisterTest("Complete Program Mutating", func() string {
		return `namespace main

struct Counter {
    count: int32
    
    mutating increment(self c: *Counter) {
        c.count++
    }
    
    func get_count(self c: Counter) int32 {
        return c.count
    }
}

func main() int32 {
    let counter = Counter{count: 0}
    counter.increment()
    let value = counter.get_count()
    return value
}`
	})
}