package golang

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var ans []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if len(ans) == 0 || ans[len(ans)-1] < x {
				ans = append(ans, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func intersection2(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	var ans []int
	for i := range set1 {
		if _, ok := set2[i]; ok {
			ans = append(ans, i)
		}
	}
	return ans
}
