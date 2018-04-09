package main

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n <= 2 {
		return 0
	}

	var count int
	dp := make([]int, n) // dp[i] 表示以 A[i] 结尾的等差数列子集的个数
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
		count += dp[i]
	}
	return count
}

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}
