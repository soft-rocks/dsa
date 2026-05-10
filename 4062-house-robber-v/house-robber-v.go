func rob(nums []int, colors []int) int64 {
    var res int64
    var lastEvenMoney, lastOddMoney, evenMoney, oddMoney int64
    curColor := 0
    isEven := true
    for i, color := range colors {
        money := nums[i]
        if color != curColor {
            res += max(evenMoney, oddMoney)
            curColor = color
            evenMoney = int64(money)
            oddMoney, lastEvenMoney, lastOddMoney = 0,0,0
            isEven = true
        } else {
            isEven = !isEven
            if isEven {
                lastEvenMoney, evenMoney = evenMoney, max(evenMoney + int64(money), lastOddMoney + int64(money))
            } else {
                lastOddMoney, oddMoney = oddMoney, max(oddMoney + int64(money), lastEvenMoney + int64(money))
            }
        }
    }
    res += max(evenMoney, oddMoney)
    return res
}
func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}