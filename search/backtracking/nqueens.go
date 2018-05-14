package main

import "fmt"

func solveNQueens(n int) [][]string {
	var queens [][]byte
	for i := 0; i < n; i++ {
		var dot []byte
		for j := 0; j < n; j++ {
			dot = append(dot, '.')
		}
		queens = append(queens, dot)
	}
	return doSolveNQ(queens, 0, n, nil)
}

func doSolveNQ(queens [][]byte, row int, n int, result [][]string) [][]string {
	if row == n {
		var r []string
		for _, value := range queens {
			r = append(r, string(value))
		}
		return append(result, r)
	}

	for col := 0; col < n; col++ {
		if isValidNQ(queens, row, col, n) {
			queens[row][col] = 'Q'
			result = append(result, doSolveNQ(queens, row+1, n, nil)...)
			queens[row][col] = '.'
		}
	}

	return result
}

func isValidNQ(quees [][]byte, row, col int, n int) bool {
	// 检查同一列上有无Q，只检查已经放置过Q的行，下面检查45°和135°也一样
	for i := 0; i < row; i++ {
		if quees[i][col] == 'Q' {
			return false
		}
	}
	// 向上检查45°对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if quees[i][j] == 'Q' {
			return false
		}
	}
	//  向上检查135°对角巷
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if quees[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func main() {
	result := solveNQueens(4)
	fmt.Println(result)
}
