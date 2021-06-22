package LinkListByGo

import "testing"

/**
单链表操作
*/

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i< 10; i++ {
		l.Inserthead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(3))
	t.Log(l.FindByIndex(6))
	t.Log(l.FindByIndex(9))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.FindByIndex(6)))
	l.Print()
}