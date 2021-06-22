package link

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//删除链表中等于给定值 val 的所有节点。力扣203
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val != val {
			break
		}
		head = head.Next
	}
	cur := head
	pre := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return head
}

//206. 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur.Next != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	cur.Next = pre
	return cur
}

//力扣：142. 环形链表 II
//思路：我们遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func createListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}
	// temp对res复制，二者没有关系了
	temp := res
	for i := 1; i < len(nums); i++ {
		// 对temp的子属性赋值
		temp.Next = &ListNode{Val: nums[i]}
		// 再将temp指向当前子属性 作为下一轮循环的当前值
		temp = temp.Next
	}

	return res
}

func printListNode(head *ListNode) string {
	s := ""
	for head != nil {
		s += fmt.Sprintf("%d->", head.Val)
		head = head.Next
	}
	return s
}
