package main

import "math"

func CountingSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	var max int = math.MinInt32
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	c := make([]int, max+1)
	//计算每个元素的个数，放入 c 中
	for i := range arr {
		c[arr[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}
	r := make([]int, n)
	for i := range arr {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}
	copy(arr, r)
}
