package main

import "fmt"

//1) 数组的插入、删除、按照下标随机访问操作；
//2）数组中的数据是int类型的；
type array struct {
	data  []int
	count int
	n     int
}

func (a *array) NweArray(capacity int) {
	for i := int(0); i < capacity; i++ {
		a.data = append(a.data, i+1)
	}
	a.count = 0
	a.n = capacity
}

func main() {
	var a array
	a.NweArray(10)
	fmt.Println(a)
	a.insert(11, 3)
	fmt.Println(a)
}

//随机查找
func (a *array) find(index int) int {
	if index < 0 || index > a.n {
		return -1
	}
	return a.data[index]
}

//删除index位置的值
func (a *array) delete(index int) bool {
	if index < 0 || index > a.n {
		return false
	}
	for i := index + 1; i < a.n; i++ {
		a.data[i-1] = a.data[i]
	}
	//a.data由len = 10 变成了9，所以要去掉最后一位
	a.data = a.data[:a.n-1]
	return true
}

//给index的位置插入element
func (a *array) insert(element, index int) {
	firstSlice := a.data[0:index]
	secondSlice := make([]int, 0)
	for i, v := range a.data {
		if i > index-1 {
			secondSlice = append(secondSlice, v)
		}
	}
	firstSlice = append(firstSlice, element)
	a.data = append(firstSlice, secondSlice...)
}
//时间复杂度O(n)

//func insertArray2(element int) {
//	firstSlice := []int{}
//	secondSlice := []int{}
//	for i := 0; i < len(array); i++ {
//		if i < 3 {
//			firstSlice = append(firstSlice, array[i])
//		} else {
//			secondSlice = append(secondSlice, array[i])
//		}
//	}
//	array = append(append(firstSlice, element), secondSlice...)
//}

//纯循环写，时间复杂度O(n)
