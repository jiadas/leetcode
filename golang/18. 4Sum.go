package golang

import (
	"sort"
)

// https://leetcode.com/problems/4sum/discuss/8545/Python-140ms-beats-100-and-works-for-N-sum-(Ngreater2)/9494
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			threeResult := threeSumWithTarget(nums[i+1:], target-nums[i])
			for _, ints := range threeResult {
				result = append(result, append(ints, nums[i]))
			}
		}
	}
	return result
}

func threeSumWithTarget(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi, sum := i+1, len(nums)-1, target-nums[i]
			for lo < hi {
				ts := nums[lo] + nums[hi]
				if ts == sum {
					result = append(result, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if ts > sum {
					hi--
				} else {
					lo++
				}
			}
		}
	}
	return result
}
