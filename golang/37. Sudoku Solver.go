package golang

var (
	boxSize   = 3
	boardSize = 9
)

var (
	rows  = make([][]int, boardSize)
	cols  = make([][]int, boardSize)
	boxes = make([][]int, boardSize)
)

// 官方题解思路说得很清晰，但是代码太啰嗦：
// https://leetcode-cn.com/problems/sudoku-solver/solution/jie-shu-du-by-leetcode/
// 讨论区的代码简介直观，但是它在检查能否填某个数的时候用的是遍历，可以像官方题解一样用空间换时间，但是回溯的时候需要额外维护行、列和 box：
// https://leetcode.com/problems/sudoku-solver/discuss/15752/Straight-Forward-Java-Solution-Using-Backtracking
func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}

	for i := 0; i < boardSize; i++ {
		rows[i] = make([]int, boardSize+1)
		cols[i] = make([]int, boardSize+1)
		boxes[i] = make([]int, boardSize+1)
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] != '.' {
				d := int(board[i][j] - '0')
				// 将现有数组初始化到 rows, cols, boxes
				placeNumber(board, d, i, j)
			}
		}
	}

	solve(board)
}

func placeNumber(board [][]byte, d, row, col int) {
	rows[row][d]++
	cols[col][d]++
	boxes[boxIndex(row, col)][d]++
	board[row][col] = byte(d + '0')
}

func removeNumber(board [][]byte, d, row, col int) {
	rows[row][d]--
	cols[col][d]--
	boxes[boxIndex(row, col)][d]--
	board[row][col] = '.'
}

func couldPlace(d, row, col int) bool {
	return rows[row][d] == 0 && cols[col][d] == 0 && boxes[boxIndex(row, col)][d] == 0
}

func solve(board [][]byte) bool {
	var num byte
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			num = board[i][j]
			if num == '.' {
				for d := 1; d <= 9; d++ {
					if couldPlace(d, i, j) {
						placeNumber(board, d, i, j)
						if solve(board) { // backtracking
							return true
						} else {
							removeNumber(board, d, i, j)
						}
					}
				}
				// 试完 1-9 还没有解出数独，说明该数独没有解
				return false
			}
		}
	}
	return true
}

func boxIndex(row, col int) int {
	return (row/boxSize)*boxSize + col/boxSize
}
