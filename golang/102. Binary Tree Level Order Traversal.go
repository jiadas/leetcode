package golang

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	currentLevel := []*TreeNode{root}
	var nextLevel []*TreeNode
	var tmp []int
	for i := 0; i < len(currentLevel); {
		node := currentLevel[i]
		tmp = append(tmp, node.Val)
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
		if i == len(currentLevel)-1 {
			currentLevel = nextLevel
			nextLevel = nil
			result = append(result, tmp)
			tmp = nil
			i = 0
		} else {
			i++
		}
	}
	return result
}
