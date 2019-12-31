package main

import "fmt"

// This problem is essentially let us to find whether there are several numbers in a set which are able to sum to a specific value (in this problem, the value is sum/2).
//
// Actually, this is a 0/1 knapsack problem, for each number, we can pick it or not.
// Let us assume dp[i][j] means whether the specific sum j can be gotten from the first i numbers.
// If we can pick such a series of numbers from 0-i whose sum is j, dp[i][j] is true, otherwise it is false.
//
//Base case: dp[0][0] is true; (zero number consists of sum 0 is true)
//
// Transition function:
// For each number,
// if we don’t pick it, dp[i][j] = dp[i-1][j], which means if the first i-1 elements has made it to j, dp[i][j] would also make it to j (we can just ignore nums[i]).
// If we pick nums[i]. dp[i][j] = dp[i-1][j-nums[i]], which represents that j is composed of the current value nums[i] and the remaining composed of other previous numbers.
// Thus, the transition function is dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
func canPartition(nums []int) bool {
	var sum int
	for _, value := range nums {
		sum += value
	}
	if sum%2 != 0 {
		return false
	}

	w := sum / 2
	dp := make([]bool, w+1) // dp[i][j] 表示前前 i 个数的和是否等于 j，用一维的 dp 数组来优化空间
	dp[0] = true
	for _, num := range nums {
		for j := w; j >= 0; j-- {
			if j >= num {
				dp[j] = dp[j] || dp[j-num]
			}
		}
	}
	return dp[w]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
