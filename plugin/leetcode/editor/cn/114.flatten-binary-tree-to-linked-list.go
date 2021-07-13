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
func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	right := root.Right
	flatten(right)

	left := root.Left
	if left != nil {
		flatten(left)
		bottomRight := left
		for bottomRight.Right != nil {
			bottomRight = bottomRight.Right
		}
		root.Left = nil
		root.Right = left
		bottomRight.Right = right
	}
}

// leetcode submit region end(Prohibit modification and deletion)
