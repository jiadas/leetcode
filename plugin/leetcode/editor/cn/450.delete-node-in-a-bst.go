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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNodeInRight := getMinNodeInRight(root.Right)
		root.Right = deleteNode(root.Right, minNodeInRight.Val)
		minNodeInRight.Left = root.Left
		minNodeInRight.Right = root.Right
		root = minNodeInRight
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMinNodeInRight(rightRoot *TreeNode) *TreeNode {
	for rightRoot.Left != nil {
		rightRoot = rightRoot.Left
	}
	return rightRoot
}

// leetcode submit region end(Prohibit modification and deletion)
