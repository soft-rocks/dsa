const mod = 1000000007

func fib(n int) int {
	a, b := 0, 1
    for i := 0; i < n; i++ {
        a, b = b, (a+b)%mod
    }
    return a % mod
}

func countHousePlacements(n int) int {
	onSide := fib(n + 2)
	return (onSide * onSide) % mod
}