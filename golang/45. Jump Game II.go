package golang

// 贪心算法，每次找局部最优，最后达到全局最优
// https://leetcode.wang/leetCode-45-Jump-Game-II.html
func jump(nums []int) int {
	var steps, maxPosition, end int
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = maxInt(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
