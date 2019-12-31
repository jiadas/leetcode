package main

import "fmt"

// 题目要求：The word can be constructed from letters of sequentially adjacent cell,
// where "adjacent" cells are those horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	var visited [][]bool
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfsExist(board, m, n, i, j, word, 0, visited) {
				return true
			}
		}
	}
	return false
}

func dfsExist(board [][]byte, m, n int, x, y int, word string, offset int, visited [][]bool) bool {
	if offset == len(word) {
		return true
	}
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[offset] || visited[x][y] {
		return false
	}
	visited[x][y] = true
	for _, step := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		nx := x + step[0]
		ny := y + step[1]

		if dfsExist(board, m, n, nx, ny, word, offset+1, visited) {
			return true
		}
	}
	visited[x][y] = false
	return false
}

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))

	fmt.Println(exist([][]byte{{'a'}}, "ab"))
}
