package DynamicProgramming

//一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。

//dp[i] = dp[i-1] + dp[i-2] +...+dp[1]
func jumpFloor(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			dp[i] += dp[j]
		}
	}
	return dp[n]
}
