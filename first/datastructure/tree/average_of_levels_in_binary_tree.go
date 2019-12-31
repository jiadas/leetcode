package main

func averageOfLevels(root *TreeNode) []float64 {
	var ret []float64
	if root == nil {
		return ret
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		cnt := len(queue)
		var sum float64
		for i := 0; i < cnt; i++ {
			node := queue[i]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, sum/float64(cnt))
		queue = queue[cnt:]
	}

	return ret
}
