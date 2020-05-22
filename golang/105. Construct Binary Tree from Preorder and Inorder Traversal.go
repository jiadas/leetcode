package golang

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9/
func buildTree(preorder []int, inorder []int) *TreeNode {
	pl, il := len(preorder), len(inorder)
	if pl == 0 || il == 0 || pl != il {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if pl == 1 {
		return root
	}
	for i, in := range inorder {
		if preorder[0] == in {
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			if i+1 <= len(inorder)-1 {
				root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			}
			break
		}
	}
	return root
}
