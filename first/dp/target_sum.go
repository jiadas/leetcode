package main

import "fmt"

// 该问题可以转换为 Subset Sum 问题，从而使用 0-1 背包的方法来求解。
//
// 可以将这组数看成两部分，P 和 N，其中 P 使用正号，N 使用负号，有以下推导：
//
//                   sum(P) - sum(N) = target
// sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
//                        2 * sum(P) = target + sum(nums)
// 因此只要找到一个子集，令它们都取正号，并且和等于 (target + sum(nums))/2，就证明存在解。
func findTargetSumWays(nums []int, S int) int {
	var sum int
	for _, value := range nums {
		sum += value
	}
	if S > sum || (sum+S)&1 == 1 {
		return 0
	}

	w := (sum + S) >> 1
	dp := make([]int, w+1) // dp[i][j] 表示前 i 个数里和为 j 的子集个数
	dp[0] = 1
	for _, num := range nums {
		for i := w; i >= 0; i-- {
			if i >= num {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}
	return dp[w]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
