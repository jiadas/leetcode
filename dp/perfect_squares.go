package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1) // dp[i] 表示能分割 i 的平方数的最少个数
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		j, s := 1, 1
		for s <= i {
			if dp[i-s]+1 < dp[i] {
				dp[i] = dp[i-s] + 1
			}
			j++
			s = j * j
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
}
