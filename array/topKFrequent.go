package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	ret := make([]int, 0)
	m := map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}
	temp := make([]int, 0)
	for _, v := range m {
		temp = append(temp, v)
	}
	sort.Ints(temp)
	temp1 := temp[len(temp)-k : len(temp)]
	for k, v := range m {
		if v >= temp1[0] && v <= temp1[len(temp1)-1] {
			ret = append(ret, k)
		}
	}
	return ret
}
