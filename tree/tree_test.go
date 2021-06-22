package tree

import "testing"

func TestBuildTree(t *testing.T) {
	a := []int{3, 9, 20, 15, 7}
	b := []int{9, 3, 15, 20, 7}
	ret := buildTree(a, b)
	t.Log(ret)
}

func TestLayer(t *testing.T) {
	a := []int{3, 9, 20, 15, 7}
	b := []int{9, 3, 15, 20, 7}
	tree := buildTree(a, b)
	ret := levelOrder(tree)
	t.Log(ret)
}

func TestMergeTrees(t *testing.T) {
	a := []int{1, 3, 5, 2}
	b := []int{5, 3, 1, 2}
	tree := buildTree(a, b)

	a1 := []int{2, 1, 4, 3, 7}
	b1 := []int{1, 4, 2, 3, 7}
	tree1 := buildTree(a1, b1)
	ret := mergeTrees(tree, tree1)
	t.Log(ret)
}

func TestPreorderTraversal(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2}
	tree := buildTree(a, b)
	ret := preorderTraversal(tree)
	t.Log(ret)
}

func TestPreorderTraversal1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2}
	tree := buildTree(a, b)
	ret := preorderTraversal1(tree)
	t.Log(ret)
}

func TestInorderTraversal(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2}
	tree := buildTree(a, b)
	ret := inorderTraversal(tree)
	t.Log(ret)
}

func TestInorderTraversal1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2}
	tree := buildTree(a, b)
	ret := inorderTraversal1(tree)
	t.Log(ret)
}

func TestInvertTree(t *testing.T) {
	a := []int{4, 2, 1, 3, 7, 6, 9}
	b := []int{1, 2, 3, 4, 6, 7, 9}
	tree := buildTree(a, b)
	ret := invertTree(tree)
	t.Log(ret)
}
