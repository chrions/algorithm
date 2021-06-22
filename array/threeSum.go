package main

import (
	"sort"
)

//数组中三数之和等于0
//思路：先将数组排序，然后从头节点开始遍历，left为外层遍历的i+1，right为数组尾
//因为三数之和为0，所以只需要遍历数组中小于0的数
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		right := len(nums) - 1
		if nums[i] > 0 {
			break //三数之和为0，i和left必定是负数。大于0就不必遍历了
		}
		if i > 0 && nums[i] == nums[i-1] { //对i去重，如果前一个i对应的数存在三数之和=0，则已经存在，如果没有则当前i也不会有
			continue
		}
		left := i + 1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] { // 对同一个i来说，如果有两个相同的left则重复
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] { // 对同一个i来说，如果有两个相同的right则重复
				right--
				continue
			}

			if nums[left]+nums[right]+nums[i] > 0 {
				right--
				continue
			}
			if nums[left]+nums[right]+nums[i] < 0 {
				left++
				continue
			}
			if nums[left]+nums[right]+nums[i] == 0 {
				temp := make([]int, 0)
				temp = append(temp, nums[left], nums[right], nums[i])
				res = append(res, temp)
				left++
				right--
			}
		}
	}
	return res
}

//复习的一次

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot := nums[0]
	left, right := 0, len(nums)-1
	for i := 1; i < len(nums); {
		if nums[i] < pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		}
		if nums[i] > pivot {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}
	}
}
