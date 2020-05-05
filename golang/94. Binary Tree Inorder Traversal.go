package golang

// https://leetcode.com/problems/validate-binary-search-tree/discuss/32112/Learn-one-iterative-inorder-traversal-apply-it-to-multiple-tree-questions-(Java-Solution)
func inorderTraversal(root *TreeNode) []int {
	var inorder []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		inorder = append(inorder, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return inorder
}
