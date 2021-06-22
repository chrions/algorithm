package link

import "testing"

func TestCreateListNode(t *testing.T) {
	a := []int{1, 2, 3, 6, 4, 5, 6}
	ret := createListNode(a)
	t.Log(printListNode(ret))

}

func TestRemoveElements(t *testing.T) {
	a := []int{1, 1, 3, 6, 1, 5, 6}
	l := createListNode(a)

	ret := removeElements(l, 1)
	t.Log(printListNode(ret))
}

func TestReverseList(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	l := createListNode(a)

	ret := reverseList(l)
	t.Log(printListNode(ret))
}

func TestDetectCycle(t *testing.T) {
	a := []int{1, 2}
	l := createListNode(a)
	head := l
	for l.Next != nil {
		l = l.Next
		if l.Next == nil {
			l.Next = head
			break
		}

	}
	ret := detectCycle(l)
	t.Log(printListNode(ret))
}
