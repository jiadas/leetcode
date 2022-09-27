package golang

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	var (
		result [][]int
		list   []int
	)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[:len(list)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
