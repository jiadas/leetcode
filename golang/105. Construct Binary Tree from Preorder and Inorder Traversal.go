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
			// https://golang.org/ref/spec#Slice_expressions
			// a[low : high]
			// For convenience, any of the indices may be omitted.
			// A missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
			// If both indices are constant, they must satisfy low <= high.
			// For arrays or strings, the indices are in range if 0 <= low <= high <= len(a), otherwise they are out of range.
			// For slices, the upper index bound is the slice capacity cap(a) rather than the length.
			// a := make([]int, 3, 20)
			// s := a[3:]    // valid, same as s := a[3:3]
			// s := a[4:]    // invalid, same as s := a[4:3], because default high is len(a), so not satisfy with low <= high
			// s := a[15:20] // valid, for slices, 0 <= low <= high <= cap(a), and satisfy with low <= high
			// so can use inorder[i+1:] directly without check i+1 < len(inorder)
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return root
}
