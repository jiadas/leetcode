package golang

type matchRecord map[string]map[string]bool

// 官方题解的 dp 解法太难懂，还是 10. 正则表达式匹配 的 dp 更清晰，稍微改一下就能 ac
func isMatch44(s string, p string) bool {
	p = removeDuplicateStars(p)

	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 匹配模式开头的 '*' 可以匹配空字符串
	if p != "" && p[0] == '*' {
		dp[0][1] = true
	}
	dp[0][0] = true

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			curS, curP := s[i-1], p[j-1]
			if curP == curS || curP == '?' {
				dp[i][j] = dp[i-1][j-1]
			}
			if curP == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

// https://leetcode-cn.com/problems/wildcard-matching/solution/tong-pei-fu-pi-pei-by-leetcode/
// 方法一：带记忆的递归会超时
func isMatchF1(s string, p string) bool {
	p = removeDuplicateStars(p)
	dp := make(matchRecord)
	return matchHelper(dp, s, p)
}

func removeDuplicateStars(p string) string {
	if p == "" {
		return p
	}
	result := make([]byte, 0, len(p))
	result = append(result, p[0])
	var lastIndex int
	for i := 1; i < len(p); i++ {
		if result[lastIndex] != '*' || p[i] != '*' {
			result = append(result, p[i])
			lastIndex++
		}
	}
	return string(result)
}

func matchHelper(dp matchRecord, s, p string) bool {
	if dp[s][p] {
		return true
	}
	if _, ok := dp[s][p]; !ok {
		dp[s] = make(map[string]bool)
	}
	switch {
	case p == s || p == "*":
		dp[s][p] = true
	case p == "" || s == "":
		dp[s][p] = false
	case s[0] == p[0] || p[0] == '?':
		dp[s][p] = matchHelper(dp, s[1:], p[1:])
	case p[0] == '*':
		dp[s][p] = matchHelper(dp, s[1:], p) || matchHelper(dp, s, p[1:])
	default:
		dp[s][p] = false
	}
	return dp[s][p]
}
