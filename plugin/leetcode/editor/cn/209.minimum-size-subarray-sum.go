package golang

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt64
	var sum int
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			ans = minInt(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
