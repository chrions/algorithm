package search

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

//二分查找，因为是升序，找到左边第一个=target的下标，再找右边最后一个=target的下标 log(n)
func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	left := leftBound(nums, target)
	right := rightBound(nums, target)
	ret[0] = left
	ret[1] = right
	return ret
}

func leftBound(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

func rightBound(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}
