package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1) // dp[i] 表示和为 i 的正整数的最大乘积
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			// j*dp[i-j] 表示 j 与保存(和为 i-j 的正整数组合)乘积相乘
			// i=10,j=6 -> i-j=4 和为4的组合有1+1+1+1，1+1+2，1+3，2+2，
			// 当j=6，和为10的组合就有6+(1+1+1+1),6+(1+1+2),6+(1+3),6+(2+2)
			// j*dp[i-j]就是上述这几种组合中的最大乘积
			if j*dp[i-j] > dp[i] {
				dp[i] = j * dp[i-j]
			}
			// 上述少一种组合，就是6+4
			if j*(i-j) > dp[i] {
				dp[i] = j * (i - j)
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(10))
}
