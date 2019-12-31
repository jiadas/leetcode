package main

import "fmt"

func findCircleNum(M [][]int) int {
	var result int
	n := len(M)
	mark := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		if mark[i] == 0 {
			stack = append(stack, i)
			var node int
			for len(stack) != 0 {
				node, stack = stack[len(stack)-1], stack[:len(stack)-1]
				mark[node] = 1 // 标记已经访问过
				for i := 0; i < n; i++ {
					// 寻找跟 node 相连的下个节点
					if mark[i] == 0 && M[node][i] == 1 {
						stack = append(stack, i)
					}
				}
			}
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))

	fmt.Println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	}))
	fmt.Println(findCircleNum([][]int{
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}))
}
