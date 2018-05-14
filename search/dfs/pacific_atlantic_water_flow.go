package main

import "fmt"

var (
	stepsPA = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m1, n1  int
)

// 题目要求：Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.
// 让太平洋和大西洋的水能相互流动的点，所以从最外围的点开始dfs，从一边（太平洋）遍历到另一边（大西洋）就说明谁能流过去，反之亦然
func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m1, n1 = len(matrix), len(matrix[0])

	var canReachP, canReachA [][]bool
	for i := 0; i < m1; i++ {
		canReachP = append(canReachP, make([]bool, n1))
		canReachA = append(canReachA, make([]bool, n1))
	}

	for i := 0; i < m1; i++ {
		dfsPA(matrix, canReachP, i, 0)
		dfsPA(matrix, canReachA, i, n1-1)
	}

	for j := 0; j < n1; j++ {
		dfsPA(matrix, canReachP, 0, j)
		dfsPA(matrix, canReachA, m1-1, j)
	}

	var result [][]int
	for i := 0; i < m1; i++ {
		for j := 0; j < n1; j++ {
			if canReachA[i][j] && canReachP[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func dfsPA(matrix [][]int, canReach [][]bool, x, y int) {
	if canReach[x][y] {
		return
	}
	// 默认能运行到这儿的都是可以流动到的，即流动条件在调用递归前已经判断
	canReach[x][y] = true
	for i := 0; i < len(stepsPA); i++ {
		nx := x + stepsPA[i][0]
		ny := y + stepsPA[i][1]
		if nx < 0 || nx >= m1 || ny < 0 || ny >= n1 || matrix[x][y] > matrix[nx][ny] {
			continue
		}

		dfsPA(matrix, canReach, nx, ny)
	}
}

func main() {
	fmt.Println(pacificAtlantic([][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}))
}
