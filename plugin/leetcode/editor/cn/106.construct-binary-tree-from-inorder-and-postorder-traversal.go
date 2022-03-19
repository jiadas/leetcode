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
func buildTree(inorder []int, postorder []int) *TreeNode {
	il, pl := len(inorder), len(postorder)
	if il == 0 || pl == 0 || il != pl {
		return nil
	}
	rootVal := postorder[pl-1]
	root := &TreeNode{Val: rootVal}
	var rootIndex int
	for i, val := range inorder {
		if rootVal == val {
			rootIndex = i
			break
		}
	}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:pl-1])
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
