package Link

//单链表反转
func Reverse(node *Node) *Node {
	var headNode *Node
	var previousNode *Node
	var currentNode = node
	for currentNode != nil {
		nextNode := currentNode.Next
		if nextNode == nil {
			headNode = currentNode
		}
		//交换位置
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return headNode
}

//检测环
func CheckCircle(node *Node) bool {
	if node == nil {
		return false
	}
	fast := node.Next
	slow := node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

//合并两个有序链表
func MergeSortedLists(la *Node, lb *Node) *Node {
	if la == nil || lb == nil {
		return nil
	}

	p := la.Next //0,3,6,9,12
	q := lb.Next //0,2,4,6,8
	head := &Node{}

	for p != nil && q != nil {
		//循环取小
		if p.Data > q.Data {
			head.Next = q
			q = q.Next
		} else {
			head.Next = p
			p = p.Next
		}
		head = head.Next
	}

	if p != nil {
		head.Next = p
	} else {
		head.Next = q
	}

	return head
}

//删除倒数第K个节点
func deleteLastKth(node *Node, k int64) *Node {
	if node.Next == nil && k == 0 && node == nil {
		return nil
	}
	fast := node
	for i := int64(1); i <= k && fast != nil; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return nil
	}
	slow := node
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return slow
}

//获取中间节点
func findMiddleNode(node *Node) *Node {
	if node == nil || node.Next == nil {
		return nil
	}
	if node.Next.Next == nil {
		return node.Next
	}
	slow := node
	fast := node
	for slow != nil && fast !=nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}