package LinkListByGo

import "fmt"

/**
单链表操作
*/
//链表在内存中并不是连续存储，所以对CPU 缓存不友好
type ListNode struct {
	next  *ListNode
	value interface{}
}

//head是哨兵节点，head.next是第一个节点
type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfterNode(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNode := p.next
	p.next = newNode
	newNode.next = oldNode
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBeforeNode(p *ListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) Inserthead(v interface{}) bool {
	return this.InsertAfterNode(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfterNode(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}

	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		cur = cur.next
		pre = pre.next
	}
	pre.next = p.next
	this.length--
	p = nil
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
