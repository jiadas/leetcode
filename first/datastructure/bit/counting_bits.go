package main

// 对于数字 6(110)，它可以看成是 4(100) 再加一个 2(10)，因此 dp[i] = dp[i&(i-1)] + 1
func countBits(num int) []int {
	ret := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// i&(i-1) 去除 i 的位级表示中最低的那一位
		ret[i] = ret[i&(i-1)] + 1
	}
	return ret
}
