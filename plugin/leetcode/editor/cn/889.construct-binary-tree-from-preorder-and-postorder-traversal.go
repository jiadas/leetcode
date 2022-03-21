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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	preL, postL := len(preorder), len(postorder)
	if preL == 0 || postL == 0 || preL != postL {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	rootVal := preorder[1]
	var leftRootIndex int
	for i, val := range postorder {
		if rootVal == val {
			leftRootIndex = i
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:leftRootIndex+2], postorder[:leftRootIndex+1])
	root.Right = constructFromPrePost(preorder[leftRootIndex+2:], postorder[leftRootIndex+1:postL-1])
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
