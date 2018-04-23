package main

// 树相关的 case 构建太麻烦了，直接在 leetcode 上执行 run code 检查
func invertTree(root *TreeNode) *TreeNode {
	// 不用要求左右子树对称
	// if root == nil || root.Left == nil || root.Right == nil {
	if root == nil {
		return root
	}

	// 如果只简单递归的交换左右节点的 Val 值，只会左子树里左右交换&右子树里左右交换，结果如下：
	//      4                 4
	//    /   \             /   \
	//   2     7      ->   7     2
	//  / \   / \         / \   / \
	// 1   3 6   9       3   1 9   6
	// 正确的结果应该是：
	//      4                 4
	//    /   \             /   \
	//   2     7      ->   7     2
	//  / \   / \         / \   / \
	// 1   3 6   9       9   6 3   1
	//left := invertTree(root.Left)
	//right := invertTree(root.Right)
	//left.Val, right.Val = right.Val, left.Val

	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)
	return root
}
