package main

import "math"

//剑指 Offer 42. 连续子数组的最大和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp[i] 表示以nums[i] 结尾的连续子数组的最大和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))
		max = int(math.Max(float64(dp[i]), float64(max)))
	}
	return max
}
