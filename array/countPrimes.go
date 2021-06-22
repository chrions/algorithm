package main

//统计所有小于非负整数 n 的质数的数量。
//
//示例 1：
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//
//来源：力扣（LeetCode）204。
//思路：厄拉多塞筛法
func countPrimes(n int) int {
	count := 0
	flag := make([]bool, n)
	for i := 2; i < n; i++ {
		if flag[i] {
			continue
		}
		count++
		for j := i * 2; j < n; j += i {
			flag[j] = true
		}
	}
	return count
}
