package golang

func findMin154(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
