package golang

import "sort"

// https://leetcode.com/problems/3sum/discuss/7380/Concise-O(N2)-Java-solution
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi, sum := i+1, len(nums)-1, 0-nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] == sum {
					result = append(result, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] > sum {
					// improve: skip duplicates
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					hi--
				} else {
					// improve: skip duplicates
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					lo++
				}
			}
		}
	}
	return result
}

// Wrong Answer
// 思路：先排序，a, b, c 中至少有一正一负，所以从两端开始遍历，找剩下的那个数
// 错误的点在于从两端开始找，处理不好游标的移动
func threeSumWA(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var result [][]int
	for i, j := 0, len(nums)-1; i < j; {
		a, b := nums[i], nums[j]
		c := a + b
		if c < 0 && c+b < 0 {
			i++
			continue
		}
		if c > 0 && c+a > 0 {
			j--
			continue
		}
		for k := i + 1; k < j; k++ {
			if c+nums[k] == 0 {
				result = append(result, []int{a, b, -c})
				break
			}
		}
		if c > 0 {
			j--
			for j > 0 && nums[j] == nums[j+1] {
				j--
			}
		} else {
			i++
			for i < len(nums)-1 && nums[i] == nums[i-1] {
				i++
			}
		}
	}
	return result
}
