func climbStairs(n int, costs []int) int {
    dp := make([]int, n+1)
    dp[0] = 0 
    for j := 1; j <= n; j++ {
        dp[j] = math.MaxInt // infinite
        for k := 1; k <= 3 && j-k >= 0; k++ {
            i := j-k // i = previous step (j-k), skip if invalid (i < 0)
            if i >= 0 {
                dp[j] = min(dp[j], dp[i] + costs[j-1] + k*k)
            }
        }
    }
    return dp[n]
}