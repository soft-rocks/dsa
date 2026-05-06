func minCostClimbingStairs(cost []int) int {
    if len(cost) <= 1 {
        return 0
    }
    dp := make([]int, len(cost))
    dp[0] = cost[0]
    dp[1] = cost[1]
    for i := 2; i < len(cost); i++ {
        dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
    }
    return min(dp[len(cost) - 1], dp[len(cost) - 2])
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}