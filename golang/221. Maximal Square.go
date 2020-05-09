package golang

// https://leetcode-cn.com/problems/maximal-square/solution/zui-da-zheng-fang-xing-by-leetcode-solution/
func maximalSquare1(matrix [][]byte) int {
	var maxSide int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	rows, columns := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				maxSide = maxInt(maxSide, 1)
				curMaxSide := minInt(rows-i, columns-j)
			find:
				// m 从 1 开始指的是从边长为 2 的正方形开始遍历，因为边长是 1 的在上一个 if 已经检查过了
				for m := 1; m < curMaxSide; m++ {
					// 先检查对角线
					if matrix[i+m][j+m] == '0' {
						// 因为是从小边长往大边长找，小的都不满足条件，更别提大的了，所以直接 break
						break find
					}
					for n := 0; n < m; n++ {
						// 再检查新增的一行和一列
						if matrix[i+m][j+n] == '0' || matrix[i+n][j+m] == '0' {
							break find
						}
					}
					maxSide = maxInt(maxSide, m+1)
				}
			}
		}
	}
	return maxSide * maxSide
}
