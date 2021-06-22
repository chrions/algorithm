package Link

import "fmt"

type Element int

type Node struct {
	Data Element
	Next *Node
}

//添加 头结点，数据
func Add(head *Node, element Element) {
	point := head
	//遍历到最后Node
	for point.Next != nil {
		point = point.Next
	}
	node := &Node{} //创建一个新的节点
	point.Next = node
	node.Data = element
	//if GetLength(head) > 0 {
	//	Traverse(head)
	//}
}

//删除index上的节点
func Delete(head *Node, index int) Element {
	if index < 0 || index > GetLength(head) {
		return 0
	}
	point := head
	for i := 0; i < index-1; i++ {
		point = point.Next //移到下一个节点
	}
	point.Next = point.Next.Next //跳过index,直接到index的下一个Node
	data := point.Next.Data
	return data
}

//插入index位置节点
func Insert(head *Node, index int, element Element) {
	if index < 0 || index > GetLength(head) {
		return
	}
	point := head
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	node := &Node{}
	node.Next = point.Next.Next
	point.Next = node
	node.Data = element
}

func GetLength(head *Node) int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

func GetDataByIndex(head *Node, index int) Element {
	if index < 0 || index > GetLength(head) {
		return 0
	}
	point := head
	for i := 0; i < index; i++ {
		point = point.Next
	}
	return point.Data
}

func GetIndexByData(head *Node, data Element) *Node {
	point := head
	var index int
	for point.Next != nil {
		if point.Data == data {
			return point
		}
		index++
		point = point.Next
		if index > GetLength(head)-1 {
			return nil
		}
	}
	return nil
}

//遍历头节点
func Traverse(head *Node){
	point := head.Next
	for point.Next != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("Traverse ok")
}
