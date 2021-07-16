package main

//剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	low := 0
	hight := len(nums) - 1
	up := 0
	for low <= hight {
		mid := (hight + low) / 2
		if nums[mid] < target {
			low = mid + 1
		}
		if nums[mid] > target {
			hight = mid - 1
		}
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				up = mid
				break
			} else {
				low = mid + 1
			}
		}
	}
	if nums[up] != target {
		return 0
	}
	max := up
	down := 0
	for down <= up {
		mid := (down + up) / 2
		if nums[mid] > target {
			up = mid - 1
		}
		if nums[mid] < target {
			down = mid + 1
		}
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				down = mid
				break
			} else {
				up = mid - 1
			}
		}
	}

	return max - down + 1
}
