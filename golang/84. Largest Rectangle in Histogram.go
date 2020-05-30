package golang

func largestRectangleArea(heights []int) int {
	var ans int
	// 在原 heights 的头尾加上高度为 0 的哨兵（sentinel）保证 stack 操作始终不为空
	newHeights := make([]int, len(heights)+2)
	for i, height := range heights {
		newHeights[i+1] = height
	}
	heights = newHeights
	stack := NewIntStack()
	stack.Push(0)
	for i := 1; i < len(heights); i++ {
		for heights[stack.Top()] > heights[i] {
			h := heights[stack.Pop()]
			w := i - stack.Top() - 1
			ans = maxInt(ans, h*w)
		}
		stack.Push(i)
	}
	return ans
}

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/zhu-zhuang-tu-zhong-zui-da-de-ju-xing-by-leetcode-/
// 栈
func largestRectangleArea2(heights []int) int {
	var ans int
	stack := NewIntStack()
	for i := 0; i < len(heights); i++ {
		for !stack.IsEmpty() && heights[stack.Top()] > heights[i] {
			h := heights[stack.Pop()]
			w := i
			if !stack.IsEmpty() {
				w = i - stack.Top() - 1
			}

			ans = maxInt(ans, h*w)
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		hi := stack.Pop()
		h := heights[hi]
		w := len(heights)
		if !stack.IsEmpty() {
			w = len(heights) - stack.Top() - 1
		}
		ans = maxInt(ans, h*w)
	}
	return ans
}

// 暴力解法
func largestRectangleArea1(heights []int) int {
	var ans int
	for i := 0; i < len(heights); i++ {
		var l, r int
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < heights[i] {
				break
			}
			r++
		}
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}
			l++
		}
		width := l + r + 1
		s := width * heights[i]
		if ans < s {
			ans = s
		}
	}
	return ans
}
