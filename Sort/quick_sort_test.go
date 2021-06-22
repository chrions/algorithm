package main

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{2, 2, 2, 2, 1, 1, 1}
	QuickSort(arr)
	fmt.Println(arr)
}
