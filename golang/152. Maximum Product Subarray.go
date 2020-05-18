package golang

// https://leetcode-cn.com/problems/maximum-product-subarray/solution/cheng-ji-zui-da-zi-shu-zu-by-leetcode-solution/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxF, minF, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = maxInt(mx*nums[i], maxInt(mn*nums[i], nums[i]))
		minF = minInt(mx*nums[i], minInt(mn*nums[i], nums[i]))
		result = maxInt(maxF, result)
	}
	return result
}
