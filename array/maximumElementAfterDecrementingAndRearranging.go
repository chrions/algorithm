package main

import (
	"sort"
)

// 力扣 1846. 减小和重新排列数组后的最大元素
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1], arr[i]) > 1 {
			arr[i+1] = arr[i] + 1
		}
	}
	max := arr[len(arr)-1]
	return max
}

func abs(i, j int) int {
	if i-j >= 0 {
		return i - j
	} else {
		return -(i - j)
	}
}
