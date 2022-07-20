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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesHelper(1, n)
}

func generateTreesHelper(lo, hi int) []*TreeNode {
	if lo > hi {
		return []*TreeNode{nil}
	}

	var ans []*TreeNode
	for i := lo; i <= hi; i++ {
		left := generateTreesHelper(lo, i-1)
		right := generateTreesHelper(i+1, hi)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				ans = append(ans, root)
			}
		}
	}

	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
