package main

func findTarget(root *TreeNode, k int) bool {
	num := inOrder(root, nil)
	l, r := 0, len(num)-1
	for l < r {
		sum := num[l] + num[r]
		switch {
		case sum == k:
			return true
		case sum > k:
			r -= 1
		case sum < k:
			l += 1
		}
	}
	return false
}

func inOrder(root *TreeNode, num []int) []int {
	if root == nil {
		return num
	}
	num = inOrder(root.Left, num)
	num = append(num, root.Val)
	num = inOrder(root.Right, num)
	return num
}
