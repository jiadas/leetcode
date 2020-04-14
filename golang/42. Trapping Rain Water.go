package golang

// 讨论区里高票的解法是双指针，解释很少而且直接看代码也不理解，
// 官方题解里双指针做了图解，单也只是帮助理解代码是如何执行的。
// 官方题解下的有一个评论对双指针解释的很好：
// https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/327718
// 对于位置 left 而言，它左边最大值一定是 left_max，右边最大值 “大于等于” right_max，这时候，
// 如果 left_max<right_max 成立，那么它就知道自己能存多少水了。无论右边将来会不会出现更大的 right_max，都不影响这个结果。
func trap(height []int) int {
	var ans int
	left, right := 0, len(height)-1
	var maxLeft, maxRight int
	for left <= right {
		if maxLeft < maxRight {
			ans += maxInt(0, maxLeft-height[left])
			maxLeft = maxInt(maxLeft, height[left])
			left++
		} else {
			ans += maxInt(0, maxRight-height[right])
			maxRight = maxInt(maxRight, height[right])
			right--
		}
	}
	return ans
}
