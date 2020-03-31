package golang

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		pre, cur := nums[j-1], nums[j]
		if pre != cur {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
