package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	var dp [][]int // dp 表示 m 个0和 n 个1 最多可以组成 strs 前 i 个 string 的最大个数
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for _, str := range strs {
		var zero, one int
		for _, s := range str {
			switch s {
			case '0':
				zero += 1
			case '1':
				one += 1
			}
		}
		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
				if i >= zero && j >= one {
					if dp[i-zero][j-one]+1 > dp[i][j] {
						dp[i][j] = dp[i-zero][j-one] + 1
					}
				}
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
