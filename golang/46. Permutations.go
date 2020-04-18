package golang

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for i, num := range nums {
		left := append(append(nums[:0:0], nums[:i]...), nums[i+1:]...)
		tmp := permute(left)
		for _, t := range tmp {
			r := append([]int{num}, t...)
			result = append(result, r)
		}
	}
	return result
}
