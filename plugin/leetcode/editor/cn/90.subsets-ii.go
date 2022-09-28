package golang

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	var (
		result [][]int
		list   []int
	)
	// 先排序
	sort.Ints(nums)
	subsetsWithDupBacktrack(nums, 0, list, &result)
	return result
}

func subsetsWithDupBacktrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i := pos; i < len(nums); i++ {
		// 排序之后，如果再遇到重复元素，则不选择此元素
		if i != pos && nums[i] == nums[i-1] {
			continue
		}

		// pos 这个元素还是要选择的，如果判断 i > 0 则会跳过 pos 这个元素
		// if i > 0 && nums[i] == nums[i-1] {
		// 	continue
		// }

		list = append(list, nums[i])
		subsetsWithDupBacktrack(nums, i+1, list, result)
		list = list[:len(list)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
