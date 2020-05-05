package golang

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/solution/yan-zheng-er-cha-sou-suo-shu-by-leetcode-solution/
// 官方题解方法一加入了上下边界的判断来解决
// 想到了用中序遍历来校验是否为二叉搜索树，但是代码写不出来，卧槽
func isValidBST(root *TreeNode) bool {
	return judgeWithRange(root, math.MinInt64, math.MaxInt64)
}

func judgeWithRange(root *TreeNode, low, up int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= up {
		return false
	}
	return judgeWithRange(root.Left, low, root.Val) && judgeWithRange(root.Right, root.Val, up)
}

// 错误点：以左子树为例，如果该二叉树的左子树不为空，没有判断左子树上所有节点的值均小于它的根节点的值
func isValidBSTFailed(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Val >= root.Right.Val) {
		return false
	}
	return isValidBSTFailed(root.Left) && isValidBSTFailed(root.Right)
}
