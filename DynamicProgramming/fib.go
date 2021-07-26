package DynamicProgramming

//f(n) = f(n-1) + f(n-2)
//f(0) = 0, f(1) = 1, f(2) = 1, f(3) = f(2)+ f(1)

//力扣 509. 斐波那契数
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	ret := 0
	for i := 2; i <= n; i++ {
		ret = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = ret
	}
	return ret
}
