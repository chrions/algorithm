package main

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//你可以假设 nums1 和 nums2 不会同时为空。
//先排序，在判断奇偶找中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	mergerNmus := GetMergeSortnums(nums1, nums2, n, m)

	i := (n + m) % 2
	if i == 0 {
		mid := len(mergerNmus) / 2
		right := mergerNmus[mid]
		left := mergerNmus[mid-1]
		return (float64(left) + float64(right)) / 2
	}
	return float64((mergerNmus[len(mergerNmus)/2]))
}

func GetMergeSortnums(nums1 []int, nums2 []int, n, m int) []int {
	i := n - 1
	j := m - 1
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	if nums1[0] >= nums2[j] {
		nums2 = append(nums2, nums1...)
		return nums2
	}
	if nums1[i] <= nums2[0] {
		nums1 = append(nums1, nums2...)
		return nums1
	}
	mergeNum := make([]int, n+m)
	a := 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			mergeNum[n+m-a] = nums1[i]
			a++
			i--
		} else {
			mergeNum[n+m-a] = nums2[j]
			a++
			j--
		}
	}
	if i == -1 {
		for n1 := 0; n1 <= j; n1++ {
			mergeNum[n1] = nums2[n1]
		}

	}
	if j == -1 {
		for n1 := 0; n1 <= i; n1++ {
			mergeNum[n1] = nums1[n1]
		}
	}
	return mergeNum
}
