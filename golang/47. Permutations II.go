package golang

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return doPermuteUnique(nums)
}

func doPermuteUnique(nums []int) [][]int {
	var result [][]int
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var num int
	for i := 0; i < len(nums); i++ {
		num = nums[i]
		left := append(append(nums[:0:0], nums[:i]...), nums[i+1:]...)
		tmp := doPermuteUnique(left)
		for _, t := range tmp {
			r := append([]int{num}, t...)
			result = append(result, r)
		}
		for i+1 < len(nums) && nums[i+1] == num {
			i++
		}
	}
	return result
}
