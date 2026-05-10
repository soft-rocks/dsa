func rob(nums []int, colors []int) int64 {
    if len(nums) == 1 { return int64(nums[0]) }
    var a, b, c int // ~ dp[i-2], dp[i-1], dp[i]
    
    // init first two maximums to rob:
    a = nums[0] // rob house 0
    if colors[0] == colors[1] {  // if colors are the same:
        b = max(a, nums[1])      // rob either house 0 or house 1
    } else {
        b = a + nums[1]          // if colors are different, rob both houses
    }
    if len(nums) == 2 { return int64(b) } // no more houses to rob
    for i := 2; i < len(nums); i++ {  // iterate from 2 index (third house; previous max's are initialised: a, b)
        if colors[i-1] != colors[i] { // if colors are different
            c = max(a + nums[i], b + nums[i]) // colors are different; we need to rob house 100%, maximizing previous robs: a, b
        } else {
            c = max(b, a + nums[i]) // if colors are the same: either do not rob it, or rob (adding previous previous state)
        }
        a, b = b, c // fixup previous states
    }
    return int64(c)
}