package golang

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
// 官方题解的动态规划
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, i+1)
	}

	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = minInt(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}

	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = minInt(ans, f[n-1][i])
	}
	return ans
}

// 官方题解的动态规划+空间优化的解释不好理解
// https://leetcode.cn/problems/triangle/solution/di-gui-ji-yi-hua-dp-bi-xu-miao-dong-by-sweetiee/
// 👆这个题解更好理解
func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = minInt(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

//leetcode submit region end(Prohibit modification and deletion)
