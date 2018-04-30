package main

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	// 全局变量在每次处理新 case 时要重新初始化
	diameter = 0
	depth(root)
	return diameter
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left+right > diameter {
		diameter = left + right
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// 效率不及用全局变量的，因为存在重复计算 depth
func diameterOfBinaryTreeWithoutGlobalVar(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 经过根节点的直径
	leftDepth := depthT(root.Left)
	rightDepth := depthT(root.Right)
	dt := leftDepth + rightDepth

	// 左子树的直径，即最大直径只出现在左子树的情况
	dl := diameterOfBinaryTreeWithoutGlobalVar(root.Left)

	// 右子树的直径，即最大直径只出现在右子树的情况
	dr := diameterOfBinaryTreeWithoutGlobalVar(root.Right)

	d := dt
	if d < dl {
		d = dl
	}
	if d < dr {
		d = dr
	}
	return d
}

func depthT(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depthT(root.Left)
	right := depthT(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
