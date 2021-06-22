package LinkListByGo

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode2{
		Val: 2,
		Next: &ListNode2{
			Val: 4,
			Next: &ListNode2{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode2{
		Val: 5,
		Next: &ListNode2{
			Val: 6,
			Next: &ListNode2{
				Val:  4,
				Next: nil,
			},
		},
	}
	ret := AddTwoNumbers(l1, l2)
	t.Log(ret)
}
