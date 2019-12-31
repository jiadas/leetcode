package main

import "fmt"

// https://leetcode.com/problems/unique-paths/discuss/22954/0ms-5-lines-DP-Solution-in-C++-with-Explanations
func uniquePathsDP(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	// 一次走一排，因为规定只能往下或者往右，所以在走下一排的时候，路径数就是左边一点+上边一点(上边一点刚好是上一排存储的值)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

// 也可以直接用数学公式求解，这是一个组合问题。
// 机器人总共移动的次数 S=m+n-2，向下移动的次数 D=m-1，那么问题可以看成从 S 从取出 D 个位置的组合数量，这个问题的解为 C(S, D)。
func uniquePaths(m int, n int) int {
	s, d := m+n-2, m-1
	result := 1
	// C(S,D) = S!/D!(S-D)! = (S-D)!(S-D+1)(S-D+2)...(S-D+D)/D!(S-D)! = (S-D+1)(S-D+2)...(S-D+D)/D!
	for i := 1; i <= d; i++ {
		result = result * (s - d + i) / i
	}

	return result
}

func main() {
	fmt.Println(uniquePathsDP(3, 7))
	fmt.Println(uniquePaths(3, 7))
}
