package golang

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func uniquePaths1(m int, n int) int {
	cur := make([]int, n)
	cur[0] = 1

	pre := make([]int, n)
	for i := 0; i < n; i++ {
		// 双重循环是从第二行开始循环，所以 pre 初始是第一行
		pre[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = pre[j] + cur[j-1]
		}
		cur, pre = pre, cur
	}

	// 最后一次循环的时候，把 cur 换成 pre 了，所以最终的答案在 pre 里
	return pre[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
