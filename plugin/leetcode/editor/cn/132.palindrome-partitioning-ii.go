package golang

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	n := len(s)
	// g(i, j) 表示 s[i..j] 是否为回文串
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
		for j := range g[i] {
			g[i][j] = true
		}
	}
	// i 和 j 的初始值方向跟 g[i+1][j-1] 相关
	// 因为 g[i+1][j-1] 是 i+1 和 j-1，表示从两侧向中间收缩
	// 所以 i 和 j 要从两头开始，提前为 i+1 和 j-1 算好结果
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
		}
	}
	// 状态：f[i] 表示字符串的前缀 s[0..i] 的最少分割次数
	f := make([]int, n)
	for i := 0; i < n; i++ {
		// s[0..i] 本身就是一个回文串，此时其不需要进行任何分割
		if g[0][i] {
			continue
		}
		f[i] = math.MaxInt
		for j := 0; j < i; j++ {
			if g[j+1][i] {
				f[i] = minInt(f[i], f[j]+1)
			}
		}
	}
	return f[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
