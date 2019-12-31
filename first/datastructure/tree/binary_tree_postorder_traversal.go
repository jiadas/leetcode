package main

// 前序遍历为 root -> left -> right，后序遍历为 left -> right -> root，
// 可以修改前序遍历成为 root -> right -> left，那么这个顺序就和后序遍历正好相反。
func postorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}
		ret = append(ret, node.Val)
		stack = append(stack, node.Left) // 这里跟前序遍历的非递归实现刚好相反
		stack = append(stack, node.Right)
	}
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}
	return ret
}
