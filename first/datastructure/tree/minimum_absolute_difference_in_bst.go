package main

import "math"

var (
	minDiff int
	preVal  int // 用全局变量，实现O(1)的空间复杂度。也可以用数组记录遍历结果，再遍历数组，但是空间复杂度为O(n)
)

func getMinimumDifference(root *TreeNode) int {
	minDiff = math.MaxInt32
	preVal = -1 // Given a binary search tree with non-negative values, 所以用-1表示没有
	processInOrder(root)
	return minDiff
}

func processInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	processInOrder(root.Left)
	if preVal != -1 {
		// 因为 bst 的中序遍历是有序的，所以直接用后值减去前值，就是绝对值
		diff := root.Val - preVal
		if diff < minDiff {
			minDiff = diff
		}
	}
	preVal = root.Val
	processInOrder(root.Right)
}
