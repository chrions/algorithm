package main

import (
	"math"
)

//给定一个二进制数组， 计算其中最大连续1的个数。
//
//示例 1:
//
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
//
//来源：力扣（LeetCode）485
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		}
		if nums[i] != 1 {
			count = 0

		}
	}
	return maxCount
}

//来源：力扣（LeetCode）495
func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	}
	ret := duration
	for i := 0; i < n-1; i++ {
		repeat := timeSeries[i] + duration - timeSeries[i+1]
		if repeat > 0 {
			ret += duration - repeat
		} else {
			ret += duration
		}
	}
	return ret
}

//来源：力扣（LeetCode）414
func thirdMax(nums []int) int {
	max, sec, third := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == max || v == sec || v == third {
			continue
		}
		if max < v {
			third = sec
			sec = max
			max = v

		} else if sec < v {
			third = sec
			sec = v
		} else if third < v {
			third = v
		}
	}
	if third > math.MinInt64 {
		return third
	}
	return max
}

//来源：力扣（LeetCode）659
func isPossible(nums []int) bool {
	countMap := make(map[int]int) //记录每个数字的剩余个数
	for _, v := range nums {
		countMap[v]++
	}
	endCnt := make(map[int]int) //记录以某个数字结尾的子序列的个数
	for _, v := range nums {
		if countMap[v] == 0 {
			continue
		}
		if endCnt[v-1] > 0 { //如果存在以v-1结尾的子序列，将v加在其后面
			countMap[v]--
			endCnt[v-1]--
			endCnt[v]++
		} else if countMap[v+1] > 0 && countMap[v+2] > 0 { // 否者，生成一个长度为3的连续子序列
			endCnt[v+2]++
			countMap[v]--
			countMap[v+1]--
			countMap[v+2]--
		} else {
			return false
		}
	}
	return true
}
