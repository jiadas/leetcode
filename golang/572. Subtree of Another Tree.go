package golang

// func isSubtree(s *TreeNode, t *TreeNode) bool {
// 	if isSame(s, t) {
// 		return true
// 	}
// 	if s == nil && t != nil {
// 		return false
// 	}
// 	if s != nil && t == nil {
// 		return true
// 	}
// 	if s != nil {
// 		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
// 	}
// 	return false
// }
//
// func isSame(s, t *TreeNode) bool {
// 	if s == nil && t == nil {
// 		return true
// 	}
// 	if (s == nil && t != nil) || (s != nil && t == nil) {
// 		return false
// 	}
// 	if s.Val != t.Val {
// 		return false
// 	}
// 	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
// }

// 上述注释部分是自己写的，一些 if 判断显得啰嗦，根据官方题解进行了精简
// https://leetcode-cn.com/problems/subtree-of-another-tree/solution/ling-yi-ge-shu-de-zi-shu-by-leetcode-solution/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s == nil && t != nil) || (s != nil && t == nil) || (s.Val != t.Val) {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
