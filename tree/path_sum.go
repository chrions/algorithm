package tree

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	find(root, sum, path)
	return res
}

func find(root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		// path是全局的，避免后面的路径把它污染，所以先把当前结果拷贝出来
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	find(root.Left, sum-root.Val, path)
	find(root.Right, sum-root.Val, path)
}
