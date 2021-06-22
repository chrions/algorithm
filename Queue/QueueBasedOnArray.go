package Queue

import "fmt"

type ArrayQueue struct {
	p        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (this *ArrayQueue) Enqueue(v interface{}) bool {
	//队列已经满了
	if this.tail == this.capacity {
		return false
	}
	this.p[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) Dequeue() interface{} {
	//队列是空的
	if this.head == this.tail {
		return nil
	}
	v := this.p[this.head]
	this.head++
	return v
}

func (this *ArrayQueue) String() string {


	if this.head == this.tail {
		return "empty queue"
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", this.p[i])
	}
	result += "<-tail"
	return result
}
