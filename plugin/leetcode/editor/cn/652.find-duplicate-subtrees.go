package golang

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode

	treeCountByTreeSerial := make(map[string]int)
	var lookup func(root *TreeNode) string
	lookup = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		serial := fmt.Sprintf("%d,%s,%s", root.Val, lookup(root.Left), lookup(root.Right))
		treeCountByTreeSerial[serial] += 1
		if treeCountByTreeSerial[serial] == 2 {
			ans = append(ans, root)
		}
		return serial
	}

	lookup(root)

	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
