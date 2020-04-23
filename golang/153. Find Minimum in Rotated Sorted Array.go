package golang

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
