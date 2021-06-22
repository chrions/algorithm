package main

import (
	"strings"
)

//给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。
//现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
//
//求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
//
//说明: 请尽可能地优化你算法的时间和空间复杂度。
//
//示例 1:
//
//输入:
//nums1 = [3, 4, 6, 5]
//nums2 = [9, 1, 2, 5, 8, 3]
//k = 5
//输出:
//[9, 8, 6, 5, 3]
//
//来源：力扣（LeetCode） 321
//思路：要在两个数组中取K个数，可理解为从nums1中取i个数，nums2中取k-i个数，通过这个条件取出两个数组中所有符合的情况，作比对取最大的数
func maxNumber(nums1 []int, nums2 []int, k int) []int {

	m, n := len(nums1), len(nums2)

	res := ""
	strNum1, strNum2 := toString(nums1), toString(nums2)
	for i := max(0, k-n); i <= min(k, m); i++ {
		k1, k2 := m-i, n-k+i
		stack1 := removeKdigits(strNum1, k1)
		stack2 := removeKdigits(strNum2, k2)

		temp := merge(stack1, stack2)
		if temp > res {
			res = temp
		}
	}

	return toNums(res)
}

func toNums(num string) []int {
	res := make([]int, 0, len(num))
	for i := range num {
		res = append(res, int(num[i]-'0'))
	}
	return res
}

func toString(nums []int) string {
	var res strings.Builder
	for i := range nums {
		res.WriteByte(byte(nums[i] + '0'))
	}
	return res.String()
}

func removeKdigits(num string, k int) string {

	res := []byte{}
	for i := range num {
		for k > 0 && len(res) > 0 && res[len(res)-1] < num[i] {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, num[i])
	}

	if k > 0 {
		res = res[:len(res)-k]
	}

	return string(res)
}

func merge(nums1, nums2 string) string {
	i, j := 0, 0
	res := make([]byte, 0, len(nums1)+len(nums2))
	for i < len(nums1) || j < len(nums2) {
		var temp byte
		if i < len(nums1) && j < len(nums2) {
			//精髓：转化成string来比较最高位的大小。string比较，从最高位开始比，最高位相同，从次高位开始比较
			//比如： 10013 和 10014，两个string相比，nums2更大
			if nums1[i:] > nums2[j:] {
				temp = nums1[i]
				i++
			} else {
				temp = nums2[j]
				j++
			}
		} else if i < len(nums1) { //当nums2走完时，把nums1的拼在后面
			temp = nums1[i]
			i++
		} else if j < len(nums2) { //如上，同理
			temp = nums2[j]
			j++
		}
		res = append(res, temp)
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
