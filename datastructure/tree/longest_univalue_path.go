package main

var path int

func longestUnivaluePath(root *TreeNode) int {
	path = 0
	dns(root)
	return path
}

// 深度遍历二叉树，返回以 root 为起点的相同节点值的最大路径长度
func dns(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftPath, rightPath int
	if root.Left != nil {
		left := dns(root.Left)
		if root.Val == root.Left.Val {
			leftPath = left + 1
		}
	}
	if root.Right != nil {
		right := dns(root.Right)
		if root.Val == root.Right.Val {
			rightPath = right + 1
		}
	}
	if path < leftPath+rightPath {
		path = leftPath + rightPath
	}
	if leftPath > rightPath {
		return leftPath
	}
	return rightPath
}
