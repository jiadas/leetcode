package main

var (
	sc  int
	kth int
)

func kthSmallest(root *TreeNode, k int) int {
	sc, kth = 0, 0
	doSearch(root, k)
	return kth
}

func doSearch(root *TreeNode, k int) {
	if root == nil {
		return
	}
	doSearch(root.Left, k)
	sc++
	if sc == k {
		kth = root.Val
		return
	}
	doSearch(root.Right, k)
}
