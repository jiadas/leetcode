package golang

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	// 状态：f[i] 表示 nums[0..i] 中以 nums[i] 结尾的最长严格递增子序列的长度
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	f := make([]int, n)
	f[0] = 1
	max := f[0]
	for i := 1; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = maxInt(f[i], f[j]+1)
			}
		}
		max = maxInt(max, f[i])
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
