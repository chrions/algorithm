package LinkListByGo

import "fmt"

/*
单链表反转
时间复杂度：O(N)
*/
func (this *LinkedList) Reverse() {
	if this.head.next == nil || this.head.next.next == nil || this.head == nil {
		return
	}
	var pre *ListNode
	cur := this.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
}

/*
判断单链表是否有环
*/
func (this *LinkedList) HasCycle() bool {
	if this.head.next == nil || this.head.next.next == nil || this.head == nil {
		return false
	}
	if this.head != nil {
		slow := this.head
		fast := this.head
		for slow != nil && fast != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

/*
*两个有序链表合并
 */
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return nil
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return nil
	}
	l := &LinkedList{head: &ListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for curl1 != nil && curl2 != nil {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}
	if curl1 != nil {
		cur.next = curl1
	} else {
		cur.next = curl2
	}
	return l
}

/*
删除倒数第N个节点
*/
func (this *LinkedList) DeleteBottomN(n int) {
	if this.head == nil || this.head.next == nil || n == 0 {
		return
	}

	fast := this.head
	for i := 1; i <= n && fast != nil; i++ {
		//到n的节点
		fmt.Println(fast.GetValue())
		fast = fast.next
	}
	if fast == nil {
		return
	}
	slow := this.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
}

/*
获取中间节点 （快的走两步，慢的走一步，相差2n，遍历完就找到中间结点）
*/
func (this *LinkedList) FindMiddleNode() *ListNode {
	if this.head == nil || this.head.next == nil {
		return nil
	}
	if this.head.next.next == nil {
		return this.head.next
	}
	slow, fast := this.head, this.head
	for fast == nil && fast.next == nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
