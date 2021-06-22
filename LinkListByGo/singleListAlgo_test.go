package LinkListByGo

import "testing"

var l *LinkedList

func init() {
	l = NewLinkedList()
	for i := 0; i< 10; i++ {
		l.InsertTail(i + 1)
	}
}

func TestReverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}