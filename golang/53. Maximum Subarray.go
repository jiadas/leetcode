package golang

import "math"

func maxSubArray(nums []int) int {
	var curSum int
	maxSum := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		maxSum = maxInt(maxSum, curSum)
	}
	return maxSum
}
