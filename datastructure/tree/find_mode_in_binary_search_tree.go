package main

var (
	count    int
	maxCount int
	preNode  *TreeNode
	modes    []int
)

func findMode(root *TreeNode) []int {
	count, maxCount = 1, 1
	preNode = nil
	modes = nil
	inOrderFind(root)
	return modes
}

func inOrderFind(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderFind(root.Left)

	// handle node
	if preNode != nil {
		if preNode.Val == root.Val {
			count++
		} else {
			count = 1
		}
	}
	if count > maxCount {
		maxCount = count
		modes = append([]int(nil), root.Val)
	} else if count == maxCount { // If a tree has more than one mode, you can return them in any order.
		modes = append(modes, root.Val)
	}
	preNode = root

	inOrderFind(root.Right)
}
