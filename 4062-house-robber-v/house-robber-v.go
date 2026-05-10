func rob(nums []int, colors []int) int64 {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return int64(nums[0])
    }
    robs := make([]int, n)
    skip := make([]int, n)
    robs[0] = nums[0]
    skip[0] = 0
    for i := 1; i < n; i++ {
        skip[i] = max(robs[i-1], skip[i-1])
        if colors[i] != colors[i-1] {
            robs[i] = nums[i] + max(robs[i-1], skip[i-1])
        } else {
            robs[i] = skip[i-1] + nums[i]
        }
    }
    return int64(max(robs[n-1], skip[n-1]))
}