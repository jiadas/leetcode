package golang

import (
	"math"
	"sort"
)

// 解题思路借鉴 15.3Sum
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	var (
		result  int
		minDiff = math.MaxInt64
	)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi, sum := i+1, len(nums)-1, target-nums[i]
			var td int
			for lo < hi {
				ts := nums[lo] + nums[hi]
				if ts == sum {
					return target
				} else if ts > sum {
					td = ts - sum
					hi--
				} else {
					td = sum - ts
					lo++
				}
				if td < minDiff {
					result = ts + nums[i]
					minDiff = td
				}
			}
		}
	}
	return result
}
