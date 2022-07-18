package golang

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) []int {
	var diff int
	for _, num := range nums {
		diff ^= num
	}
	ans := make([]int, 2)
	diff = (diff & (diff - 1)) ^ diff
	for _, num := range nums {
		// 正确的条件应该是 num&diff > 0
		// if num&diff == 1 { // 自己写的bug，以为 num 这一位为 1，num&diff == 1
		// 如果 num 这一位为 0，num&diff 应该是 0
		if num&diff == 0 {
			ans[0] ^= num
		} else {
			ans[1] ^= num
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
