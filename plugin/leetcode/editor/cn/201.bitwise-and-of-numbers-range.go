package golang

// leetcode submit region begin(Prohibit modification and deletion)
func rangeBitwiseAnd(left int, right int) int {
	var count int
	for left != right {
		left, right = left>>1, right>>1
		count++
	}
	return left << count
}

// leetcode submit region end(Prohibit modification and deletion)
