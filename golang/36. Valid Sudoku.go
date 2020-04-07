package golang

// https://leetcode-cn.com/problems/valid-sudoku/solution/you-xiao-de-shu-du-by-leetcode/
// https://leetcode.com/problems/valid-sudoku/discuss/15472/Short%2BSimple-Java-using-Strings
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]int, 9)
	cols := make([]map[byte]int, 9)
	boxes := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]int)
		cols[i] = make(map[byte]int)
		boxes[i] = make(map[byte]int)
	}

	var num byte
	var boxIndex int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			num = board[i][j]
			if num != '.' {
				boxIndex = (i/3)*3 + j/3

				rows[i][num] += 1
				cols[j][num] += 1
				boxes[boxIndex][num] += 1
				if rows[i][num] > 1 || cols[j][num] > 1 || boxes[boxIndex][num] > 1 {
					return false
				}
			}
		}
	}
	return true
}
