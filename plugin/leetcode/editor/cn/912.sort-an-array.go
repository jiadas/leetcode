package golang

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	n1 := sortArray(append([]int(nil), nums[:n/2]...))
	n2 := sortArray(append([]int(nil), nums[n/2:]...))

	var p1, p2 int
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}

	return nums
}

// leetcode submit region end(Prohibit modification and deletion)
