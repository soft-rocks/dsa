const mod = 1000000007

func countHousePlacements(n int) int {
    side := fib(n)
    return side * side % mod
}

func fib(n int) int {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return 2
    }

    a := 1
    b := 2
    result := b
    for i:= 2; i<= n; i++ {
        result = (a + b) % mod
        a = b
        b = result
    }

    return result % mod
}