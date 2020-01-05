package golang

// https://leetcode.com/problems/zigzag-conversion/discuss/3403/Easy-to-understand-Java-solution
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([][]byte, numRows)
	for i := 0; i < len(s); {
		// vertically down
		for j := 0; j < numRows && i < len(s); j++ {
			rows[j] = append(rows[j], s[i])
			i++
		}
		// obliquely up
		for j := numRows - 2; j >= 1 && i < len(s); j-- {
			rows[j] = append(rows[j], s[i])
			i++
		}
	}
	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}
	return string(result)
}
