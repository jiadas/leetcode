package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ret := make([]int, n)
	stack := make([]int, n)
	top := -1
	for i := 0; i < n; i++ {
		// 遇到一个比栈顶元素还大的数组元素，说明该元素比栈里暂存的元素都大，所有全部出栈，并计算下标差值
		for top > -1 && temperatures[i] > temperatures[stack[top]] {
			index := stack[top]
			top--
			ret[index] = i - index
		}

		// 每一个元素的 index 都入栈
		top++
		stack[top] = i
	}
	return ret
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
