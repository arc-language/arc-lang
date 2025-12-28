// test_async_simple.arc
namespace main

// Simple async function that returns immediately
async func get_value() int32 {
    return 42
}

// Async function that awaits another async function
async func fetch_data() int32 {
    let value = await get_value()
    return value
}

// Regular function
func add(a: int32, b: int32) int32 {
    return a + b
}

// Main function (not async)
func main() int32 {
    let result = add(5, 10)
    return result
}