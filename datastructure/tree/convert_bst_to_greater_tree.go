package main

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	doConvert(root)
	return root
}

func doConvert(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Right != nil {
		doConvert(root.Right)
	}
	sum += root.Val
	root.Val = sum
	if root.Left != nil {
		doConvert(root.Left)
	}
}
