package golang

// leetcode submit region begin(Prohibit modification and deletion)
func splitArray(nums []int, m int) int {
	var left, right int
	for _, num := range nums {
		right += num
		if left < num {
			left = num
		}
	}
	return search(left, right, func(i int) bool {
		sum, count := 0, 1
		for _, num := range nums {
			if sum+num > i {
				count++
				sum = num
			} else {
				sum += num
			}
		}
		return count <= m
	})
}

func search(i, j int, f func(int2 int) bool) int {
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

// leetcode submit region end(Prohibit modification and deletion)
