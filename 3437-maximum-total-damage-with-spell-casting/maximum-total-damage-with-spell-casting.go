func maximumTotalDamage(power []int) int64 {
    freq := map[int]int{}
    for _, p := range power {
        freq[p] += p
    }
    keys := make([]int, 0)
    for key := range freq {
        keys = append(keys, key)
    }
    sort.Ints(keys)
    dp := make([]int, len(keys))
    dp[0] = freq[keys[0]]
    for i:= 1; i < len(dp); i++ {
        curr := keys[i]
        dp[i] = max(dp[i-1], freq[curr])
        for j := i-1; j >= 0 && j >= i - 3; j-- {
            num := keys[j]
            if curr - num > 2 {
                dp[i] = max(dp[i], dp[j] + freq[curr])
            }
        }
    }
    return int64(dp[len(dp)-1])
    // freq := map[int]int{}
    // for _, p := range power {
    //     freq[p]++
    // }

    // nums := []int{}
    // for k, _ := range freq {
    //     nums = append(nums, k)
    // }

    // sort.Ints(nums)

    // n := len(nums)
    // dp := make([]int, n)
    // dp[0] = nums[0] * freq[nums[0]]

    // for i := 1; i < n; i++ {
    //     num := nums[i]
    //     dp[i] = max(dp[i-1], num * freq[num])
        
    //     for j := i-1; j >= 0 && j >= i-3; j-- {
    //         num2 := nums[j]
    //         if num - num2 > 2 {
    //             dp[i] = max(dp[i], num * freq[num] + dp[j])
    //         }
    //     }
    // }

    // return int64(dp[n-1])
}