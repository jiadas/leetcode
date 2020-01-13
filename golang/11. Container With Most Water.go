package golang

// https://leetcode.com/problems/container-with-most-water/discuss/6100/Simple-and-clear-proofexplanation
// ↑是原先的题解，但是没有对解题思路做很好的证明
// https://leetcode.com/problems/container-with-most-water/discuss/6099/Yet-another-way-to-see-what-happens-in-the-O(n)-algorithm
// ↑是对解题的补充证明，更好理解
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var water int
	for i < j {
		left, right := height[i], height[j]
		water = maxInt(water, (j-i)*minInt(left, right))
		if left < right {
			i += 1
		} else {
			j -= 1
		}
	}
	return water
}
