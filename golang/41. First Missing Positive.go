package golang

// https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode/
// 检查 1 是否存在于数组中。如果没有，则已经完成，1 即为答案。
// 如果 nums = [1]，答案即为 2 。
// 将负数，零，和大于 n 的数替换为 1 。
// 遍历数组。当读到数字 a 时，替换第 a 个元素的符号。
// 注意重复元素：只能改变一次符号。由于没有下标 n ，使用下标 0 的元素保存是否存在数字 n。
// 再次遍历数组。返回第一个正数元素的下标。
// 如果 nums[0] > 0，则返回 n 。
// 如果之前的步骤中没有发现 nums 中有正数元素，则返回 n + 1。
func firstMissingPositive(nums []int) int {
	if !in(nums, 1) {
		return 1
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}
	for _, num := range nums {
		num = absInt(num)
		if num == n {
			nums[0] = -absInt(nums[0])
		} else {
			nums[num] = -absInt(nums[num])
		}
	}
	// 先从 1 开始找
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	// 如果从 1 开始找，找遍了都没有，在看看是不是 n
	if nums[0] > 0 {
		return n
	}
	return n + 1
}

func in(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
