package LinkListByGo

import "container/list"

//原生双向链表实现LRU
type LRUCache struct {
	m        map[int]*node
	l        *list.List
	capacity int
}

type node struct {
	v int
	e *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        map[int]*node{},
		l:        list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.m[key]; ok {
		this.l.MoveToFront(value.e)
		return value.v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if nodeV, ok := this.m[key]; ok {
		this.l.MoveToFront(nodeV.e)
		this.m[key].v = value
	} else {
		this.l.PushFront(key)
		this.m[key] = &node{
			v: value,
			e: this.l.Front(),
		}
	}
	if len(this.m) > this.capacity { //容量到达最大
		delete(this.m, this.l.Remove(this.l.Back()).(int))
	}
}
