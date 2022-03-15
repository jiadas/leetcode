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
func diameterOfBinaryTree(root *TreeNode) int {
	_, ans := maxDepthAndDiameter(root)
	return ans
}

func maxDepthAndDiameter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l, ld := maxDepthAndDiameter(root.Left)
	r, rd := maxDepthAndDiameter(root.Right)
	return maxInt(l, r) + 1, maxInt(l+r, maxInt(ld, rd))
}

// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// leetcode submit region end(Prohibit modification and deletion)
