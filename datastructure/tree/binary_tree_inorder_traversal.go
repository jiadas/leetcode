package main

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		cur = node.Right
	}

	return ret
}
