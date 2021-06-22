package main

import "fmt"

//有序度是数组中具有有序关系的元素对的个数。有序元素对用数学表示：有序元素对：a[i] <= a[j], 如果 i < j。
//对于一个倒序排列的数组，比如 6，5，4，3，2，1，有序度是 0；对于一个完全有序的数组，比如 1，2，3，4，5，6，
//有序度就是n*(n-1)/2，也就是 15。我们把这种完全有序的数组的有序度叫作满有序度。
//逆序度 = 满有序度 - 有序度。我们排序的过程就是一种增加有序度，减少逆序度的过程，最后达到满有序度，就说明排序完成了.

//原地排序算法：特指空间复杂度是O(1)的排序算法
/*
冒泡排序、插入排序、选择排序
*/

//冒泡排序
//稳定性：冒泡排序是稳定的排序算法。
//空间复杂度：冒泡排序是原地排序算法。
func BubbleSort(a []int64) {
	if len(a) < 1 {
		return
	}
	n := len(a)
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

//插入排序
//空间复杂度：插入排序是原地排序算法。
//稳定性：插入排序是稳定的排序算法。
func InsertionSort(a []int64) {
	if len(a) < 1 {
		return
	}
	n := len(a)
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if value < a[j] {
				a[j+1] = a[j] //交换位置
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

//选择排序
//空间复杂度：选择排序是原地排序算法。
//稳定性：选择排序不是稳定的排序算法。
func SelectionSort(a []int64) {
	if len(a) <= 1 {
		return
	}
	n := len(a)
	for i := 0; i < n; i++ {
		//查找最小值
		min := i
		for j := i + 1; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		//交换
		a[i], a[min] = a[min], a[i]
	}
}

func main() {
	a := []int64{100, 102, 8, 78, 1, 20}
	InsertionSort(a)
	fmt.Println(a)
}
