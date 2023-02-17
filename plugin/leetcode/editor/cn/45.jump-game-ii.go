package golang

// leetcode submit region begin(Prohibit modification and deletion)
// 动态规划 https://greyireland.gitbook.io/algorithm-pattern/ji-chu-suan-fa-pian/dp#jump-game-ii
func jump(nums []int) int {
	n := len(nums)
	// 状态：f[i]表示从起点到 i 点的最小跳跃次数
	f := make([]int, n)
	f[0] = 0

	for i := 1; i < n; i++ {
		// f[i]的最大值就是从起点一步一步跳到 i
		f[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = minInt(f[j]+1, f[i])
			}
		}
	}

	return f[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
