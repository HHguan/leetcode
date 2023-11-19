package dongtaiguihua

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a1, a2, a3 := 1, 2, 3
	for i := 2; i < n; i++ {
		a3 = a1 + a2
		a1 = a2
		a2 = a3
	}
	return a3
}
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a1, a2, a3 := 0, 1, 1
	for i := 2; i < n; i++ {
		a3 = a1 + a2
		a1 = a2
		a2 = a3
	}
	return a3
}
