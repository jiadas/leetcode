package golang

// leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	sums [][]int
}

func Constructor3(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	sums := make([][]int, m+1)
	n := len(matrix[0])
	sums[0] = make([]int, n+1)
	for i, row := range matrix {
		sums[i+1] = make([]int, len(row)+1)
		for j, value := range row {
			sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + value
		}
	}
	return NumMatrix{sums: sums}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.sums[row2+1][col2+1] - nm.sums[row1][col2+1] - nm.sums[row2+1][col1] + nm.sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// leetcode submit region end(Prohibit modification and deletion)
