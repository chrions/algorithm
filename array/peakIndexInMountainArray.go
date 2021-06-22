package main

//力扣852
func peakIndexInMountainArray(arr []int) int {
	low := 0
	hight := len(arr) - 1
	for low <= hight {
		mid := (low + hight) / 2
		if mid-1 >= 0 && mid+1 < len(arr) && arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			low = mid
		}
		if mid-1 >= 0 && mid+1 < len(arr) && arr[mid] > arr[mid+1] && arr[mid] < arr[mid-1] {
			hight = mid
		}
		if mid-1 >= 0 && mid+1 < len(arr) && arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
	}
	return -1
}
