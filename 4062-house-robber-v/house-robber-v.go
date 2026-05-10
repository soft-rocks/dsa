func rob(nums []int, colors []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return int64(nums[0])
	}

	// rob[i] 表示搶第 i 間房的最大收益
	// skip[i] 表示不搶第 i 間房的最大收益
	rob := make([]int64, n)
	skip := make([]int64, n)

	// 初始狀態：第 0 間房
	rob[0] = int64(nums[0])
	skip[0] = 0

	for i := 1; i < n; i++ {
		// 情況 1：如果不搶第 i 間房
		// 最大收益就是前一間房「搶」或「不搶」中的較大值
		skip[i] = max(rob[i-1], skip[i-1])

		// 情況 2：如果搶第 i 間房
		if colors[i] != colors[i-1] {
			// 如果顏色不同，前一間搶不搶都沒關係
			rob[i] = max(rob[i-1], skip[i-1]) + int64(nums[i])
		} else {
			// 如果顏色相同，前一間一定不能搶，只能接在 skip[i-1] 後面
			rob[i] = skip[i-1] + int64(nums[i])
		}
	}

	// 最終結果是最後一間房「搶」或「不搶」的最大值
	return max(rob[n-1], skip[n-1])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}