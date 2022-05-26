package golang

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// func isValidBST(root *TreeNode) bool {
// 	return helper(root, math.MinInt64, math.MaxInt64)
// }
//
// func helper(root *TreeNode, low, up int) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if root.Val <= low || root.Val >= up {
// 		return false
// 	}
// 	// 那么根据二叉搜索树的性质，在递归调用左子树时，我们需要把上界 upper 改为 root.val，即调用 helper(root.left, lower, root.val)，
// 	// 因为左子树里所有节点的值均小于它的根节点的值。
// 	// 同理递归调用右子树时，我们需要把下界 lower 改为 root.val，即调用 helper(root.right, root.val, upper)。
// 	// 链接：https://leetcode.cn/problems/validate-binary-search-tree/solution/yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution/
// 	// 来源：力扣（LeetCode）
// 	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, up)
// }

// leetcode submit region end(Prohibit modification and deletion)
