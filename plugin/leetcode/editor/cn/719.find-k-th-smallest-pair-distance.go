package golang

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(i int) bool {
		var left, count int // count 表示比 i 小的距离的数量
		for right := range nums {
			for nums[right]-nums[left] > i {
				left++
			}
			count += right - left
		}
		return count >= k
	})
}

// leetcode submit region end(Prohibit modification and deletion)
