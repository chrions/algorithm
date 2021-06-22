package tree

import (
	"DataStructureAndAlgorithm/Stack"
	"fmt"
)

//前序遍历的递推公式：
//preOrder(r) = print r->preOrder(r->left)->preOrder(r->right)
//
//中序遍历的递推公式：
//inOrder(r) = inOrder(r->left)->print r->inOrder(r->right)
//
//后序遍历的递推公式：
//postOrder(r) = postOrder(r->left)->postOrder(r->right)->print r

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(rootData interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootData)}
}

//中序遍历：左中右
func (this *BinaryTree) InOrderTraverse() {
	p := this.root
	s := Stack.NewArrayStack()
	for !s.IsEmpty() || p != nil {
		if p != nil {
			s.Push(p)
			p = p.left
		} else {
			temp := s.Pop().(*Node)
			fmt.Printf("%+v ", temp.data)
			p = temp.right
		}
	}
}

//前序遍历：中左右
func (this *BinaryTree) PreOrderTraverse() {
	p := this.root
	s := Stack.NewArrayStack()
	for !s.IsEmpty() || p != nil {
		if p != nil {
			fmt.Printf("%+v ", p.data)
			s.Push(p)
			p = p.left
		} else {
			p = s.Pop().(*Node).right
		}
	}
}

//后序遍历：左右中
func (this *BinaryTree) PostOrderTraverse() {
	s1 := Stack.NewArrayStack()
	s2 := Stack.NewArrayStack()
	s1.Push(this.root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*Node)
		s2.Push(p)
		if p.left != nil {
			s1.Push(p.left)
		}
		if p.right != nil {
			s1.Push(p.right)
		}
	}
	for !s2.IsEmpty() {
		fmt.Printf("%+v ", s2.Pop().(*Node).data)
	}
}

func (this *BinaryTree) PostOrderTraverse2() {
	r := this.root
	s := Stack.NewArrayStack()

	//point to last visit node
	var pre *Node

	s.Push(r)

	for !s.IsEmpty() {
		r = s.Top().(*Node)
		if (r.left == nil && r.right == nil) || (pre != nil && (pre == r.left || pre == r.right)) {

			fmt.Printf("%+v ", r.data)
			s.Pop()
			pre = r
		} else {
			if r.right != nil {
				s.Push(r.right)
			}

			if r.left != nil {
				s.Push(r.left)
			}

		}

	}
}
