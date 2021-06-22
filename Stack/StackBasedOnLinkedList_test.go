package Stack

import "testing"

func TestLinkedListStack_Push(t *testing.T) {
	s := NewLinkedListStaack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
}

func TestLinkedListStack_Pop(t *testing.T) {
	s := NewLinkedListStaack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}

func TestLinkedListStack_Top(t *testing.T) {
	s := NewLinkedListStaack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Top())
}
