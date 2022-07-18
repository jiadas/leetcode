package golang

// leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(num uint32) int {
	var ans int
	for num != 0 {
		num = num & (num - 1)
		ans++
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
