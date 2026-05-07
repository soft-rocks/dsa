func countTexts(pressedKeys string) int {
    mod := 1000000007
    n := len(pressedKeys)
    dp := make([]int, n + 1)
    dp[0] = 1
    for i := 1; i < n + 1; i++ {
        currChar := pressedKeys[i-1]
        steps := 3
        if currChar == '7' || currChar == '9' {
            steps = 4
        }
        for step := 1; step < steps + 1 && i - step >= 0; step++ {
            if pressedKeys[i - step] == currChar  {
                dp[i] = (dp[i] + dp[i-step]) % mod
            } else {
                break
            }   
        }
    }
    return dp[len(dp) - 1]
}