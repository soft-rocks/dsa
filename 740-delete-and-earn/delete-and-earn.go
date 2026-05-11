func deleteAndEarn(nums []int) int {
	sums := make([]int, 10001)
	for _, num := range nums {
		sums[num] += num
	}
	var prev, cur int
	for _, sum := range sums {
		prev, cur = cur, max(cur, prev+sum)
	}
	return cur
}