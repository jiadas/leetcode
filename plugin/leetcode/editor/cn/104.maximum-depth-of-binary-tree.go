package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return maxInt(l, r) + 1
}

// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// leetcode submit region end(Prohibit modification and deletion)
