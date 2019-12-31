package main

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	// 这样声明的 head 默认值是 nil，后续操作 head 里的成员会 panic
	// var head *TreeNode
	//head := &TreeNode{}
	head := new(TreeNode)
	var left1, left2 *TreeNode
	var right1, right2 *TreeNode
	switch {
	case t1 != nil && t2 != nil:
		head.Val = t1.Val + t2.Val
		left1, left2 = t1.Left, t2.Left
		right1, right2 = t1.Right, t2.Right
	case t1 != nil && t2 == nil:
		head.Val = t1.Val
		left1 = t1.Left
		right1 = t1.Right
	case t1 == nil && t2 != nil:
		head.Val = t2.Val
		left2 = t2.Left
		right2 = t2.Right
	}

	head.Left = mergeTrees(left1, left2)
	head.Right = mergeTrees(right1, right2)
	return head
}
