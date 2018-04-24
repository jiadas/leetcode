package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return doCheck(root.Left, root.Right)
}

func doCheck(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return doCheck(t1.Left, t2.Right) && doCheck(t1.Right, t2.Left)
}
