package main

import "fmt"

// https://leetcode.com/problems/sudoku-solver/discuss/15752/Straight-Forward-Java-Solution-Using-Backtracking
func solveSudoku(board [][]byte) {
	solveSK(board)
}

func solveSK(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				for k := byte('1'); k <= byte('9'); k++ { //trial. Try 1 through 9
					if isValid(board, i, j, k) {
						// Put c for this cell
						board[i][j] = k

						if solveSK(board) {
							// If it's the solution return true
							return true
						} else {
							// Otherwise go back
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, b byte) bool {
	for i := 0; i < 9; i++ {
		// 检查行
		if board[row][i] != '.' && board[row][i] == b {
			return false
		}
		// 检查列
		if board[i][col] != '.' && board[i][col] == b {
			return false
		}
		// 检查3*3的区块，整个board有9个3*3区块
		// 先用 row 和 col 确认当前位于哪个3*3区块
		// 然后再在该区块里遍历，利用i
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != '.' && board[3*(row/3)+i/3][3*(col/3)+i%3] == b {
			return false
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	for _, b := range board {
		for _, v := range b {
			fmt.Print(string(v), " ")
		}
		fmt.Println()
	}
}
