package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}

	if isLeaf(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func isLeaf(node *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}
