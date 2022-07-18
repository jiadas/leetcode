package golang

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	var ans int
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
