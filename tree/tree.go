package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//力扣 105. 从前序与中序遍历序列构造二叉树
//前序遍历：根结点  左子节点  右子节点
//中序遍历：左子节点  根结点  右子节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { //递归终止条件：前序遍历的数组长度为0
		return nil
	}
	tree := &TreeNode{preorder[0], nil, nil}
	pIndex := 0
	for i, v := range inorder {
		if v == preorder[0] {
			pIndex = i //找到中序遍历中的根结点，根结点前面的左子节点，后面的是右子节点
			break
		}
	}
	tree.Left = buildTree(preorder[1:pIndex+1], inorder[0:pIndex])                           //左子节点
	tree.Right = buildTree(preorder[pIndex+1:len(preorder)], inorder[pIndex+1:len(inorder)]) //右子节点
	return tree
}

//力扣 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	// 总队列
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		// 临时队列，保存每层的结点
		tmpQueue := []*TreeNode{}
		tmpRes := []int{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}

			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}

		// 处理完一层，处理下一层
		queue = tmpQueue
		res = append(res, tmpRes)
	}

	return res
}

//力扣 617. 合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}

	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

//力扣 144. 二叉树的前序遍历 (递归)
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	preOrderBs(root, &ret)
	return ret
}

func preOrderBs(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	preOrderBs(root.Left, ret)
	preOrderBs(root.Right, ret)

}

//力扣 144. 二叉树的前序遍历 (迭代)
func preorderTraversal1(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		temp := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = temp.Right
	}
	return ret
}

//力扣 145. 二叉树的后序遍历 (递归)
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	ret = postOrder(root, ret)
	return ret
}

func postOrder(root *TreeNode, ret []int) []int {
	if root == nil {
		return ret
	}
	ret = append(postOrder(root.Left, ret), postOrder(root.Right, ret)...)
	ret = append(ret, root.Val)
	return ret
}

//力扣 145. 二叉树的后序遍历 (迭代)
func postorderTraversal1(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	//翻转
	if len(ret) > 0 {
		for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
			ret[i], ret[j] = ret[j], ret[i]
		}
	}
	return ret
}

//力扣 94. 二叉树的中序遍历（递归）
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	ret = inOrder(root, ret)
	return ret
}

func inOrder(root *TreeNode, ret []int) []int {

	if root != nil {
		ret = inOrder(root.Left, ret)
		ret = append(ret, root.Val)
		ret = inOrder(root.Right, ret)
	}

	return ret
}

//力扣 94. 二叉树的中序遍历（迭代）
func inorderTraversal1(root *TreeNode) []int {
	stack := []*TreeNode{}
	ret := make([]int, 0)

	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, temp.Val)
			root = temp.Right
		}
	}
	return ret
}

//力扣 226
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
