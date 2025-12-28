namespace main

extern libc {
    func printf(*byte, ...) int32
}

func main() int32 {
    let i: int32 = 10
    
    printf("Initial: %d\n", i) // 10
    
    // 1. Post-Increment (i++)
    // Should print 10, then i becomes 11
    printf("Post-inc expression (i++): %d\n", i++)
    printf("After post-inc: %d\n", i) // 11
    
    // 2. Pre-Increment (++i)
    // Should increment to 12, then print 12
    printf("Pre-inc expression (++i): %d\n", ++i)
    printf("After pre-inc: %d\n", i) // 12
    
    // 3. Post-Decrement (i--)
    // Should print 12, then i becomes 11
    printf("Post-dec expression (i--): %d\n", i--)
    printf("After post-dec: %d\n", i) // 11
    
    // 4. Pre-Decrement (--i)
    // Should decrement to 10, then print 10
    printf("Pre-dec expression (--i): %d\n", --i)
    printf("After pre-dec: %d\n", i) // 10
    
    // 5. Mixed Expression
    // (i++) + (++i) is undefined behavior in C, but here we enforce order
    // i is 10.
    // LHS: i++ returns 10, i becomes 11.
    // RHS: ++i increments 11->12, returns 12.
    // Result: 10 + 12 = 22. i is 12.
    // Note: Depends on evaluation order of operands (Left-to-Right in Arc)
    
    // Reset i
    i = 10
    let result = (i++) + (++i)
    printf("Mixed (i++) + (++i) [10 + 12]: %d\n", result)
    printf("Final i value: %d\n", i) // 12
    
    return 0
}