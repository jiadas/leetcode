package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 会重复计算同一层节点的高度，可以优化
	d := high(root.Left) - high(root.Right)
	if d < 0 {
		d = -d
	}
	if d > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 可以在计算高度的同时判断是否平衡
func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := high(root.Left)
	r := high(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
