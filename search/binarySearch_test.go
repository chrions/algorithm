package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	t.Log(BinarySearch(a, 5))
}

func TestBinarySearchRecursive(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	t.Log(BinarySearchRecursive(a, 8))
}

func TestBinarySearchFirst(t *testing.T) {
	a := []int{1, 2, 3, 3, 6, 8}
	t.Log(BinarySearchFirst(a, 3))
}

func TestBinarySearchLast(t *testing.T) {
	a := []int{1, 2, 3, 3, 3, 8}
	t.Log(BinarySearchLast(a, 3))
}

func TestBinarySearchFirstGT(t *testing.T) {
	a := []int{1, 2, 3, 3, 5, 8}
	t.Log(BinarySearchFirstGT(a, 3))
}

func TestBinarySearchLastLT(t *testing.T) {
	a := []int{1, 2, 3, 3, 5, 8}
	t.Log(BinarySearchLastLT(a, 5))
}

func TestSearchRange(t *testing.T) {
	nums := []int{1}
	target := 1
	t.Log(searchRange(nums, target))
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 0
	t.Log(searchInsert(nums, target))
}
