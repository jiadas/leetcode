package main

import "fmt"

var (
	steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n  int
)

// 以矩阵最外层的元素为起点，深度遍历为'O'的区域，用'T'临时标记
// 这样就去除了最外侧没被'X'包围的'O'区域，剩下的都是里侧被'X'包围的'O'区域了
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n = len(board), len(board[0])

	for i := 0; i < m; i++ {
		dfsSolve(board, i, 0)
		dfsSolve(board, i, n-1)
	}

	for j := 0; j < n; j++ {
		dfsSolve(board, 0, j)
		dfsSolve(board, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

// 去除一些重复dfs的点，但是提交到 LeetCode 之后 Runtime 没有提高，都是52ms（59个case，可能 case 太少没体现出来）
func solve1(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n = len(board), len(board[0])

	for i := 0; i < m; i++ {
		dfsSolve(board, i, 0)
		if n > 1 {
			dfsSolve(board, i, n-1)
		}
	}

	if n > 2 {
		for j := 1; j < n-1; j++ {
			dfsSolve(board, 0, j)
			if m > 1 {
				dfsSolve(board, m-1, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsSolve(board [][]byte, x, y int) {
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'T'
	for i := 0; i < len(steps); i++ {
		dfsSolve(board, x+steps[i][0], y+steps[i][1])
	}
}

func main() {
	//X X X X
	//X O O X
	//X X O X
	//X O X X
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
}
