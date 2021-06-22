package Stack

import "fmt"

/*
基于链表实现的栈
*/
type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	//栈顶节点
	topNode *node
}

func NewLinkedListStaack() *LinkedListStack {
	return &LinkedListStack{topNode: nil}
}

func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	}
	return false
}

func (this *LinkedListStack) Push(v interface{}) {
	//从栈顶往下加
	this.topNode = &node{next: this.topNode, val: v}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.topNode
	this.topNode = this.topNode.next
	return v
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.val
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
