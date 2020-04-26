package w186

// 代码应该没问题，思路就是模拟对角线路线遍历，但是超时了
func findDiagonalOrder(nums [][]int) []int {
	var result []int
	for i := range nums {
		result = append(result, find(nums, i, 0)...)
	}
	x := len(nums) - 1
	y := 1
	for x >= 0 && y >= len(nums[x]) {
		x, y = x-1, y+1
	}
	for ; x >= 0 && y < len(nums[x]); y++ {
		result = append(result, find(nums, x, y)...)

		if y == len(nums[x])-1 {
			nextRow := x - 1
			for nextRow >= 0 {
				if y+2 < len(nums[nextRow]) {
					x = nextRow
					y = y + 1
					break
				} else {
					nextRow--
					y = y + 1
				}
			}
		}
	}
	return result
}

func find(nums [][]int, x, y int) []int {
	var result []int
	for x >= 0 {
		if y < len(nums[x]) {
			result = append(result, nums[x][y])
		}
		x--
		y++
	}

	return result
}
