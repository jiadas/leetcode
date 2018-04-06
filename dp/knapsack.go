package main

import "fmt"

// 有一个容量为 N 的背包，要用这个背包装下物品的价值最大，这些物品有两个属性：体积 w 和价值 v。
//
// 定义一个二维数组 dp 存储最大价值，其中 dp[i][j] 表示体积不超过 j 的情况下，前 i 件物品能达到的最大价值。设第 i 件物品体积为 w，价值为 v，根据第 i 件物品是否添加到背包中，可以分两种情况讨论：
//
// 第 i 件物品没添加到背包，总体积不超过 j 的前 i 件物品的最大价值就是总体积不超过 j 的前 i-1 件物品的最大价值，dp[i][j] = dp[i-1][j]。
// 第 i 件物品添加到背包中，dp[i][j] = dp[i-1][j-w] + v。
// 第 i 件物品可添加也可以不添加，取决于哪种情况下最大价值更大。
//
// 综上，0-1 背包的状态转移方程为：
//        dp[i][j] = max(dp[i-1][j],dp[i-1][j-w]+v)
func knapsack2(weight, n int, weights, values []int) int {
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, weight+1))
	}
	for i := 1; i <= n; i++ {
		w, v := weights[i-1], values[i-1]
		for j := 1; j <= weight; j++ {
			if j >= w {
				if dp[i-1][j] < dp[i-1][j-w]+v {
					dp[i][j] = dp[i-1][j-w] + v
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][weight]
}

// 空间优化
// 在程序实现时可以对 0-1 背包做优化。
// 观察状态转移方程可以知道，前 i 件物品的状态仅由前 i-1 件物品的状态有关，因此可以将 dp 定义为一维数组，其中 dp[j] 既可以表示 dp[i-1][j] 也可以表示 dp[i][j]。此时，
// dp[j] = max(dp[j],dp[j-w]+v)
func knapsack1(weight, n int, weights, values []int) int {
	dp := make([]int, weight+1)
	for i := 0; i < n; i++ {
		w, v := weights[i], values[i]
		// 因为dp中下标大的要依赖小的来计算，如果小的提前有值了，会导致大的计算错误
		// 所以j要从大往小遍历
		for j := weight; j >= 1; j-- {
			//for j := 1; j <= weight; j++ {
			if j >= w {
				if dp[j] < dp[j-w]+v {
					dp[j] = dp[j-w] + v
				}
			}
		}
	}

	return dp[weight]
}

func main() {
	fmt.Println(knapsack2(5, 3, []int{1, 2, 3}, []int{6, 10, 12}))

	fmt.Println(knapsack1(5, 3, []int{1, 2, 3}, []int{6, 10, 12}))
}
