package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSubtreeWithRoot(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSubtreeWithRoot(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSubtreeWithRoot(s.Left, t.Left) && isSubtreeWithRoot(s.Right, t.Right)
}
