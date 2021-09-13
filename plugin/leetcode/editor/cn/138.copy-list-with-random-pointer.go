package golang

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	newNodeByOldNode := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		t := &Node{
			Val: cur.Val,
		}
		newNodeByOldNode[cur] = t
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		newNode := newNodeByOldNode[cur]
		if cur.Next != nil {
			newNode.Next = newNodeByOldNode[cur.Next]
		}
		if cur.Random != nil {
			newNode.Random = newNodeByOldNode[cur.Random]
		}
		cur = cur.Next
	}
	return newNodeByOldNode[head]
}

// leetcode submit region end(Prohibit modification and deletion)
