package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var dp [][]int
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j > 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case j == 0 && i > 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
			case i > 0 && j > 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
				if dp[i][j] > dp[i-1][j]+grid[i][j] {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
