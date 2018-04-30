package main

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil || root.Right == nil {
		return -1
	}

	//             1
	//         /         \
	//        1           3
	//      /    \       / \
	//    1       1     3   4
	//   / \     / \   / \ / \
	//  3   1   1   1  3 8 4 8
	// / \ / \ / \
	// 3 3 1 6 2 1
	// 在遇到上述 case 时, 下面的 if 就会导致结果错误
	//if root.Val == root.Left.Val && root.Val == root.Right.Val {
	//	return -1
	//}

	lsv, rsv := root.Left.Val, root.Right.Val
	// 根值和左子树根值相等，第二小的值有可能在左子树里
	// 所以需要找出左子树的第二小值，和右子树的根值作比较
	if root.Val == root.Left.Val {
		lsv = findSecondMinimumValue(root.Left)
	}
	// 根值和右子树根值相等
	if root.Val == root.Right.Val {
		rsv = findSecondMinimumValue(root.Right)
	}
	if lsv != -1 && rsv != -1 {
		if lsv < rsv {
			return lsv
		}
		return rsv
	}
	// 如果代码能到这一行，说明 lsv 和 rsv 两者有一个为-1或者两者都为-1
	if lsv != -1 {
		return lsv
	}
	return rsv
}
