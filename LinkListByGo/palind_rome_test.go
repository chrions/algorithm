package LinkListByGo

import "testing"

func TestIsPalindrome(t *testing.T) {
	l := NewListNode(1)
	l.next = NewListNode(2)
	l.next.next = NewListNode(2)
	l.next.next.next = NewListNode(1)
	ret := IsPalindrome(l)
	t.Log(ret)
}
