extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {
    let i: int32 = 0
    
    printf("Initial: %d\n", i)
    
    // This should print 0 (post-increment returns old value)
    // and side-effect should make i = 1
    printf("Post-inc expression: %d\n", i++)
    
    // This should print 1
    printf("After post-inc: %d\n", i)
    
    // Manual increment check
    i = i + 1
    printf("Manual add: %d\n", i)
    
    return 0
}