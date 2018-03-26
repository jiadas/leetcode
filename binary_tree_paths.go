package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	return dfsTree(root, "", nil)
}

func dfsTree(node *TreeNode, path string, paths []string) []string {
	if path != "" {
		path = fmt.Sprintf("%s->%d", path, node.Val)
	} else {
		path = fmt.Sprintf("%d", node.Val)
	}
	if node.Left == nil && node.Right == nil {
		return append(paths, path)
	}

	if node.Left != nil {
		paths = append(paths, dfsTree(node.Left, path, nil)...)
	}
	if node.Right != nil {
		paths = append(paths, dfsTree(node.Right, path, nil)...)
	}
	return paths
}

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}
