package DynamicProgramming

//力扣746. 使用最小花费爬楼梯

//dp[i] 表示到i所花费的全部体力
//dp[i] = min(dp[i-1],dp[i-2]) + cost[i]
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
