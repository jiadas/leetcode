package main

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// 以 root 为起点和为 sum 的路径数 + 左子树里和为 sum 的路径数 + 右子树里和为 sum 的路径数
	return pathSumStartWithRoot(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// pathSumStartWithRoot 计算以 root 为起点，和为 sum 的路径数
func pathSumStartWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var ret int
	if root.Val == sum {
		ret++
	}
	ret += pathSumStartWithRoot(root.Left, sum-root.Val) + pathSumStartWithRoot(root.Right, sum-root.Val)
	return ret
}
