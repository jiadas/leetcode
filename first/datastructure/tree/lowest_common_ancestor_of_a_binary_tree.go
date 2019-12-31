package main

func lca(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// If only one of them is in that subtree, then the result is that one of them.
	// 当 p 和 q 只有一个在以 root 为根的子树中是，返回它们自己作为结果
	if root == nil || root == p || root == q {
		return root
	}
	left := lca(root.Left, p, q)
	right := lca(root.Right, p, q)

	// p 和 q 不在左子树中
	if left == nil {
		return right
	}
	// p 和 q 不在右子树中
	if right == nil {
		return left
	}
	// p 和 q 分别在左右子树中，lca 为 root
	return root
}
