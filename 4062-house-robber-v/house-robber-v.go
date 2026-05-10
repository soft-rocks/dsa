func rob(nums []int, colors []int) int64 {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return int64(nums[0])
    }
    robs := make([]int, len(nums))
    skip := make([]int, len(nums))
    robs[0] = nums[0]
    skip[0] = 0
    for i := 1; i < len(nums); i++ {
        skip[i] = max(robs[i-1], skip[i-1])
        if colors[i] != colors[i-1] {
            robs[i] = max(robs[i-1], skip[i-1]) + nums[i]
        } else {
            robs[i] = skip[i-1] + nums[i]
        }
    }
    return int64(max(robs[len(nums)-1], skip[len(nums)-1]))
}