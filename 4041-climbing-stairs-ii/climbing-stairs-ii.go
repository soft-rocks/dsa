func climbStairs(n int, costs []int) int {
	a, b, c := 0, 0, 0
	for _, cost := range costs {
		a, b, c = b, c, cost+min(a+9, min(b+4, c+1))
	}
	return c
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}