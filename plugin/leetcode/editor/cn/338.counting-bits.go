package golang

// leetcode submit region begin(Prohibit modification and deletion)
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 令 y=x&(x-1)，则 y 为将 x 的最低设置位从 1 变成 0 之后的数，显然 0<=y<x，bits[x]=bits[y]+1
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
