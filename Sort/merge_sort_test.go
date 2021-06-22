package main

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	//arr := []int64{5, 4}
	//MergeSort(arr) //方法是值传递，但是数组是指针
	//fmt.Println(arr)

	arr := []int64{1, 6, 5, 3, 2, 1}
	MergeSort(arr)
	fmt.Println(arr)
}
