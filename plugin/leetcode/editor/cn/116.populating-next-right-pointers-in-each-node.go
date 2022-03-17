package golang

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node116) *Node116 {
	if root == nil {
		return root
	}
	connectTwoNodes(root.Left, root.Right)
	return root
}

func connectTwoNodes(node1, node2 *Node116) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNodes(node1.Left, node1.Right)
	connectTwoNodes(node2.Left, node2.Right)
	connectTwoNodes(node1.Right, node2.Left)
}

// leetcode submit region end(Prohibit modification and deletion)
