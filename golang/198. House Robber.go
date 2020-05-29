package golang

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	pre1, pre2 := nums[0], maxInt(nums[0], nums[1])
	ans := maxInt(pre1, pre2)
	for i := 2; i < len(nums); i++ {
		ans = maxInt(pre1+nums[i], pre2)
		pre1, pre2 = pre2, ans
	}
	return ans
}
