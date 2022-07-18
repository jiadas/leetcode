package golang

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	// 注意将所有中间变量声明成 int32
	var ans int32
	for i := 0; i < 32; i++ {
		var total int32
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	// 最后返回的时候再转换成 int
	return int(ans)
}

// leetcode submit region end(Prohibit modification and deletion)
