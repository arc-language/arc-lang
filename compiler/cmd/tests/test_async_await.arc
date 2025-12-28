// test_async_chain.arc
namespace main

async func step1() int32 {
    return 10
}

async func step2(x: int32) int32 {
    return x + 20
}

async func step3(x: int32) int32 {
    return x + 30
}

async func process() int32 {
    let a = await step1()
    let b = await step2(a)
    let c = await step3(b)
    return c  // Should be 60: 10 + 20 + 30
}

func main() int32 {
    return 0
}