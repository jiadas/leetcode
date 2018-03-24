package main

import "fmt"

func numIslands(grid [][]byte) int {
	var result int
	if len(grid) == 0 {
		return result
	}

	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfsNumIslands(grid, i, j)
				result++
			}
		}
	}

	return result
}

func dfsNumIslands(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	steps := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, s := range steps {
		dfsNumIslands(grid, i+s[0], j+s[1])
	}
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}
