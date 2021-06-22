package LinkListByGo

//输入：(2 -> 4 -> 4) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 9
//原因：342 + 465 = 807

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func AddTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	a := 0
	L3 := new(ListNode2)
	head := L3
	for l1 != nil || l2 != nil || a > 0 {
		L3.Next = new(ListNode2)
		L3 = L3.Next
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		L3.Val = (val1+val2)%10 + a
		a = (val1 + val2) / 10
	}
	return head.Next
}
