package golang

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/discuss/14699/Clean-iterative-solution-with-two-binary-searches-(with-explanation)
func searchRange(nums []int, target int) []int {
	return binarySearchPart(nums, 0, len(nums)-1, target)
}

func binarySearchPart(nums []int, l, r, target int) []int {
	result := []int{-1, -1}
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			if mid-1 >= l {
				if nums[mid-1] == target {
					firsts := binarySearchPart(nums, l, mid-1, target)
					result[0] = firsts[0]
				} else {
					result[0] = mid
				}
			} else {
				result[0] = mid
			}
			if mid+1 <= r {
				if nums[mid+1] == target {
					lasts := binarySearchPart(nums, mid+1, r, target)
					result[1] = lasts[1]
				} else {
					result[1] = mid
				}
			} else {
				result[1] = mid
			}
			return result
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return result
}
