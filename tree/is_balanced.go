package tree

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：
//
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//力扣：110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if abs(getHeight(root)) > 1 {
		return false
	}
	return true
}

func abs(height int) int {
	if height > 0 {
		return height
	}
	return -height
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
