package main

//快排
//时间复杂度O(logn)， 最坏时间复杂度O(n2)
//通过搬移数据使快排的空间复杂度是 O(1)
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置，让小于pivot的数在前面，大有pivot的数在后面
				arr[i], arr[j] = arr[j], arr[i] //
			}
			i++
		}
	}
	//此时i的前面是小于pivot的数，i以及i的后面是大于pivot的数，交换之后i前面的小于p,i后面的大于p
	arr[i], arr[end] = arr[end], arr[i]

	return i
}
