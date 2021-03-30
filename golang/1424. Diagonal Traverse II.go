package golang

// In a 2D matrix, elements in the same diagonal have same sum of their indices.
// https://leetcode.com/problems/diagonal-traverse-ii/discuss/597741/Clean-Simple-Easiest-Explanation-with-a-picture-and-code-with-comments
func findDiagonalOrder(nums [][]int) []int {
	var (
		ans    []int
		ansLen int
		maxSum int
	)
	diagonalsBySum := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			sum := i + j
			diagonalsBySum[sum] = append(diagonalsBySum[sum], nums[i][j])
			maxSum = maxInt(maxSum, sum)
			ansLen++
		}
	}

	if len(diagonalsBySum) > 0 {
		ans = make([]int, 0, ansLen)
	}

	for i := 0; i <= maxSum; i++ {
		diagonals := diagonalsBySum[i]
		for j := len(diagonals) - 1; j >= 0; j-- {
			ans = append(ans, diagonals[j])
		}
	}
	return ans
}
