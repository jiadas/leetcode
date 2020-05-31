package golang

func isSymmetric(root *TreeNode) bool {
	// if root == nil {
	// 	return true
	// }
	// return checkSame(root.Left, root.Right)
	// 	可以简化成
	return checkSame(root, root)
}

func checkSame(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	// if node1 == nil && node2 != nil || node1 != nil && node2 == nil {
	// 	return false
	// }
	// 可以简化成
	if node1 == nil || node2 == nil {
		return false
	}
	// if node1.Val != node2.Val {
	// 	return false
	// }
	// 可以放在 return 里
	return node1.Val == node2.Val && checkSame(node1.Left, node2.Right) && checkSame(node1.Right, node2.Left)
}
