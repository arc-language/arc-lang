package main

func init() {
	// Function tests
	RegisterTest("Function Basic", func() string {
		return `func add(a: int32, b: int32) int32 {
    return a + b
}`
	})

	RegisterTest("Function No Return", func() string {
		return `func print(msg: string) {
}`
	})

	RegisterTest("Function Void Return", func() string {
		return `func main() {
    let x = 42
}`
	})

	RegisterTest("Function Async", func() string {
		return `async func fetch_data(url: string) string {
    let response = await http_get(url)
    return response
}`
	})

	RegisterTest("Function Async No Return", func() string {
		return `async func process_items(items: vector<string>) {
    for item in items {
        await process(item)
    }
}`
	})

	// Async/await tests
	RegisterTest("Await Expression", func() string {
		return `async func main() {
    let data = await fetch_data("https://api.example.com")
}`
	})

	RegisterTest("Await Multiple", func() string {
		return `async func main() {
    let result1 = await task1()
    let result2 = await task2()
}`
	})

	RegisterTest("Await In If", func() string {
		return `async func main() {
    if await check_status() {
        x = 1
    }
}`
	})
}