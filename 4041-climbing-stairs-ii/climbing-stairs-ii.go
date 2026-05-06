func climbStairs(n int, costs []int) int {
    a,b,c:= 0,0,0
    for _, cost := range costs {
        next := cost + min(a + 9, b + 4, c + 1)
        a = b
        b = c
        c = next
    }
    return c
}
