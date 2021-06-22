package LinkListByGo

import "fmt"

//判断是否是回文

func IsPalindrome(head *ListNode) bool {
	if head.next == nil {
		return false
	}
	//找到中结点
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	fmt.Println(slow.GetValue())
	//反转中结点之后的链表
	pre := &ListNode{}
	for slow != nil {
		tmp := slow.next
		slow.next = pre
		pre = slow
		slow = tmp
	}
	//对比,奇数链表会多一个，不过不影响
	print(pre)
	print(head)
	for pre.next != nil && head.next != nil {
		if pre.GetValue().(int) != head.GetValue().(int) {
			return false
		}
		pre = pre.next
		head = head.next
	}
	return true
}

func print(head *ListNode) {
	format := ""
	for head != nil {
		format += fmt.Sprintf("%v -->", head.GetValue())
		head = head.next
	}
	fmt.Println(format)
}
