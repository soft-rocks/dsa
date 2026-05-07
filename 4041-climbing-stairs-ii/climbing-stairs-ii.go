func climbStairs(n int, costs []int) int {
    a, b, c := 0, 0, 0
    for _, cost := range costs {
        next := min(a + 9, b + 4, c + 1) + cost
        a = b
        b = c
        c = next
    }
    return c
}
