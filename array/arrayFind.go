package main

import (
	"math"
	"sort"
)

//来源：力扣（LeetCode）645
func findErrorNums(nums []int) []int {
	ret := []int{0, 0}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for i := 1; i <= len(nums); i++ {
		if v, ok := m[i]; ok && v > 1 {
			ret[0] = i
		}
		if _, ok := m[i]; !ok {
			ret[1] = i
		}
	}
	return ret
}

//来源：力扣（LeetCode）697
func findShortestSubArray(nums []int) int {
	mCount := make(map[int]int)
	maxD := 0
	for _, v := range nums {
		mCount[v]++
		if mCount[v] > maxD {
			maxD = mCount[v]
		}
	}
	if maxD == 0 {
		return 0
	} else if maxD == 1 {
		return 1
	}
	maxDslice := make([]int, 0)
	for k, v := range mCount {
		if v == maxD {
			maxDslice = append(maxDslice, k)
		}
	}

	minLength := 0
	for _, value := range maxDslice {
		onTime := 0
		left := -1
		right := -1
		for i, v := range nums {
			if value == v && left == -1 {
				left = i
				onTime += 1
			} else if value == v && left != -1 && onTime+1 == maxD {
				right = i
				onTime += 1
				break
			} else if value == v && left != -1 && right == -1 {
				onTime += 1
			}
		}
		length := right - left + 1
		if minLength == 0 {
			minLength = length
		} else if minLength > length {
			minLength = length
		}
	}

	return minLength
}

//来源：力扣（LeetCode）448
func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			ret = append(ret, i)
		}
	}
	return ret
}

//来源：力扣（LeetCode）442
//要点，先用快排排序，时间复杂度o(logn）,空间复杂度o(1),之后循环数组o(n)找到重复的值,所以总的时间复杂度o(n)
func findDuplicates(nums []int) []int {
	ret := make([]int, 0)
	separateSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			ret = append(ret, nums[i])
		}
	}
	return ret
}

func separateSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	i := partition(nums, start, end)
	separateSort(nums, start, i-1)
	separateSort(nums, i+1, end)
}

func partition(nums []int, start int, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] > p {
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

//来源：力扣（LeetCode）41
/*
思路:如果数组中包含 x∈[1,N]，那么恢复后，数组的第 x−1 个元素为 x。
在恢复后，数组应当有 [1, 2, ..., N] 的形式，但其中有若干个位置上的数是错误的，每一个错误的位置就代表了一个缺失的正数。
以题目中的示例 [3, 4, -1, 1] 为例，恢复后的数组应当为 [1, -1, 3, 4]，1在i=0的位置，3在i=2的位置，4在i=3的位置，-1不在范围内，所以我们就可以知道缺失的数为 -1所在的位置1 + 1 = 2。

*/
func firstMissingPositive(nums []int) int {

	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] { //重新赋值的 nums[i] 也可能是符合条件需要交换的，所以此处继续遍历直到不满足条件为止,
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1] //如果 nums[i] == nums[nums[i]-1] 则会无限循环下去
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}

	}
	return n + 1 //数组中所有的数都不在（0，n]的范围内，最小正整数位n+1
}

//来源：力扣（LeetCode）247
//先按照升序排序
func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		temp := len(citations) - i
		if citations[i] >= temp {
			return temp
		}
	}
	return 0
}

//来源：力扣（LeetCode）209
/*
定义两个指针left 和right 分别表示子数组的开始位置和结束位置，维护变量 temp 存储子数组中的元素和
（即从 nums[left] 到 nums[right] 的元素和）。

初始状态下，nums[left] 和 nums[right] 都指向下标 0，temp的值为 0。

每一轮迭代，将 nums[right] 加到 temp，如果 temp ≥ s，则更新子数组的最小长度（此时子数组的长度是 right-left+1），
然后将 nums[left] 从 temp 中减去并将 left 右移，直到 temp < s，
在此过程中同样更新子数组的最小长度。在每一轮迭代的最后，将 nums[right] 右移。
*/
func minSubArrayLen(s int, nums []int) int {
	l := math.MaxInt64
	left, right, n := 0, 0, len(nums)
	temp := 0
	for right < n {
		temp += nums[right]
		for temp >= s {
			if l > right-left+1 {
				l = right - left + 1
			}
			temp -= nums[left]
			left++
		}

		right++
	}
	if l < math.MaxInt64 {
		return l
	}
	return 0
}

//来源：力扣（LeetCode）59
func generateMatrix(n int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	rows := n
	columns := n
	top, bottom, right, left := 0, rows-1, columns-1, 0
	index := 1

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ { //top行，从左到右遍历
			result[top][column] = index
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			result[row][right] = index
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column >= left; column-- {
				result[bottom][column] = index
				index++
			}
			for row := bottom - 1; row > top; row-- {
				result[row][left] = index
				index++
			}
		}
		top++
		left++
		right--
		bottom--
	}
	return result
}

//来源：力扣（LeetCode）54
/*
官方解析：按层模拟
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	top, bottom, right, left := 0, rows-1, columns-1, 0 //找好四个边的范围
	index := 0
	ret := make([]int, rows*columns)
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ { //top行，从左到右遍历
			ret[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ { //right行，从上到下遍历
			ret[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- { //bottom行，从右到左遍历
				ret[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- { //left行，从下到上遍历
				ret[index] = matrix[row][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return ret
}
