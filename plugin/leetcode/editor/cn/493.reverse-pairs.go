package golang

// leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	mid := n >> 1
	n1 := append([]int(nil), nums[:mid]...)
	n2 := append([]int(nil), nums[mid:]...)
	cnt := reversePairs(n1) + reversePairs(n2)

	var j int
	for _, v := range n1 {
		// 在 for 循环外层定义 j，为的是不用每次都重新遍历 n2
		// 如果每次都重新遍历 n2，就会超时
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		cnt += j
	}

	var p1, p2 int
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}

	return cnt
}

// leetcode submit region end(Prohibit modification and deletion)
