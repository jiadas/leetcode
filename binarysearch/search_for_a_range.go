package main

func searchRange(nums []int, target int) []int {
	start := firstGreaterEqual(nums, target)
	if start == len(nums) || nums[start] > target {
		return []int{-1, -1}
	}
	return []int{start, firstGreaterEqual(nums, target+1) - 1}
}

// find the first number that is greater than or equal to target.
// **could return A.length if target is greater than A[A.length-1].**
// actually this is the same as lower_bound in C++ STL.
func firstGreaterEqual(nums []int, target int) int {
	l, h := 0, len(nums) // 注意 h 的初始值
	for l < h {
		m := l + (h-l)>>1
		// 这里移动 l 和 h 也有讲究
		// 因为要找到第一个大于或者等于 target 的位置(记为 first)，
		// 所以当 nums[m] < target 时，说明 first 应该在 m 的右边，所以 l = m + 1
		if nums[m] < target {
			l = m + 1
		} else {
			// 当 nums[m] >= target 时，说明 first 应该在 m 的左边，所以 h = m
			// 尤其是 nums[m] == target 时，因为要找 first，所以 m 左边可能还有等于 target 的元素
			h = m
		}
	}
	return l
}
