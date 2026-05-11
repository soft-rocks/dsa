func deleteAndEarn(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    maxVal := 0
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
    }

    sum := make([]int, maxVal+1)

    for _, num := range nums {
        sum[num] += num
    }
    var prev, curr int
    for i := 1; i < len(sum); i++ {
        next := max(curr, prev+ sum[i])
        prev = curr
        curr = next
    }
    return curr
}