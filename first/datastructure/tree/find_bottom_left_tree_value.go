package main

func findBottomLeftValue(root *TreeNode) int {
	var ret int

	if root == nil {
		return ret
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			node := queue[i]
			if i == 0 {
				ret = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[cnt:]
	}

	return ret
}
