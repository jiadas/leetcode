package golang

// https://leetcode-cn.com/problems/find-the-duplicate-number/solution/er-fen-fa-si-lu-ji-dai-ma-python-by-liweiwei1419/
func findDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		var count int
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
