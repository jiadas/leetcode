package golang

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// leetcode submit region end(Prohibit modification and deletion)
