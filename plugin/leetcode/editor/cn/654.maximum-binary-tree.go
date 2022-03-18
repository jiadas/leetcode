package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var maxIndex int
	for i, num := range nums {
		if num > nums[maxIndex] {
			maxIndex = i
		}
	}
	root := &TreeNode{
		Val: nums[maxIndex],
	}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

// leetcode submit region end(Prohibit modification and deletion)
