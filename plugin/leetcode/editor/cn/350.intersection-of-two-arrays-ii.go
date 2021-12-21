package golang

// leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	set := map[int]int{}
	for _, i := range nums1 {
		set[i] += 1
	}

	var ans []int
	for _, i := range nums2 {
		n, ok := set[i]
		if ok && n > 0 {
			ans = append(ans, i)
			set[i] -= 1
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
