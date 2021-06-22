package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHead(6)
	data := []int{7, 5, 6, 4, 2, 1}
	for _, v := range data {
		h.InsertData(v)
	}
	fmt.Println(h)
	h.RemoveMax()
	fmt.Println(h)
}

func TestHeapSort(t *testing.T) {
	data := []int{1, 9, 6, 4, 8, 10}
	h := NewHead(7)
	for _, v := range data {
		h.InsertData(v)
	}
	ret := []int{}
	ret = append(ret, data[1])
}
