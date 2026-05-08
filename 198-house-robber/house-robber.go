func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    f1 := nums[0]
    f2 := max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        next := max(nums[i] + f1, f2)
        f1 = f2
        f2 = next
    }
    return f2
}