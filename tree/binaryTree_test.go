package tree

import "testing"

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(4)
	binaryTree.root.left = NewNode(2)
	binaryTree.root.right = NewNode(7)
	binaryTree.root.left.left = NewNode(1)
	binaryTree.root.left.right = NewNode(3)
	binaryTree.root.right.left = NewNode(6)
	binaryTree.root.right.right = NewNode(9)

	binaryTree.InOrderTraverse()
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(4)
	binaryTree.root.left = NewNode(2)
	binaryTree.root.right = NewNode(7)
	binaryTree.root.left.left = NewNode(1)
	binaryTree.root.left.right = NewNode(3)
	binaryTree.root.right.left = NewNode(6)
	binaryTree.root.right.right = NewNode(9)

	binaryTree.PreOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PostOrderTraverse()
}
