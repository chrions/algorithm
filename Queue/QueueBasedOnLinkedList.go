package Queue

import "fmt"

type ListNode struct {
	value interface{}
	next  *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (this *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if this.tail == nil { //空的链表队列
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.value
	this.head = this.head.next
	this.length--
	return v
}

func (this *LinkedListQueue) String() string {
	if this.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.value)
	}
	result += "<-tail"
	return result
}
