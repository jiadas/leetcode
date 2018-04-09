package main

import "fmt"

// 可以转换为求两个字符串的最长公共子序列问题。
//
// 对于两个子序列 S1 和 S2，找出它们最长的公共子序列。
// 定义一个二维数组 dp 用来存储最长公共子序列的长度，其中 dp[i][j] 表示 S1 的前 i 个字符与 S2 的前 j 个字符最长公共子序列的长度。考虑 S1i 与 S2j 值是否相等，分为两种情况：
//   当 S1i==S2j 时，那么就能在 S1 的前 i-1 个字符与 S2 的前 j-1 个字符最长公共子序列的基础上再加上 S1i 这个值，最长公共子序列长度加 1 ，即 dp[i][j] = dp[i-1][j-1] + 1。
//   当 S1i != S2j 时，此时最长公共子序列为 S1 的前 i-1 个字符和 S2 的前 j 个字符最长公共子序列，与 S1 的前 i 个字符和 S2 的前 j-1 个字符最长公共子序列，它们的最大者，即 dp[i][j] = max{ dp[i-1][j], dp[i][j-1] }。
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	var dp [][]int
	for i := 0; i <= n1; i++ {
		dp = append(dp, make([]int, n2+1))
	}

	for index1, w1 := range word1 {
		for index2, w2 := range word2 {
			i, j := index1+1, index2+1
			if w1 == w2 {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return n1 + n2 - 2*dp[n1][n2]
}

func main() {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance("", "e"))
}
