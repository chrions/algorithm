package search

//搜索旋转排序数组
//数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2]
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

//力扣  35
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	temp := 0
	for left <= right {
		mid := (left + right) / 2
		temp = mid
		if target > nums[mid] {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
		if target == nums[mid] {
			return mid
		}
	}
	if nums[temp] > target {
		return temp
	}

	return temp + 1

}
