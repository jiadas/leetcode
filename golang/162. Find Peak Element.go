package golang

import "math"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)
	var mid int
	var front, cur, back int
	for left < right {
		mid = left + (right-left)>>1

		cur = nums[mid]
		front, back = math.MinInt64, math.MinInt64
		if mid-1 >= 0 {
			front = nums[mid-1]
		}
		if mid+1 < len(nums) {
			back = nums[mid+1]
		}

		if front < cur && cur > back {
			return mid
		} else if cur < back {
			left = mid + 1
		} else if front > cur {
			right = mid
		}
	}

	return left
}
