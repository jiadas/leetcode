package golang

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] text1 前 i 个和 text2 前 j 个字符最长公共子序列
	// dp[m+1][n+1]
	//   ' a d c e
	// ' 0 0 0 0 0
	// a 0 1 1 1 1
	// c 0 1 1 2 1
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
