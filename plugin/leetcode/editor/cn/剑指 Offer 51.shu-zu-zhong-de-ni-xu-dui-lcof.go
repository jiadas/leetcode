package golang

// leetcode submit region begin(Prohibit modification and deletion)
func reversePairsOffer51(nums []int) int {
	_, ans := reversePairsCore(nums)
	return ans
}

func reversePairsCore(nums []int) ([]int, int) {
	if len(nums) <= 1 {
		return nums, 0
	}

	mid := len(nums) >> 1
	sortedLeft, lc := reversePairsCore(nums[:mid])
	sortedRight, rc := reversePairsCore(nums[mid:])

	var (
		sortedLen   = len(sortedLeft) + len(sortedRight)
		sortedIndex = sortedLen - 1
		sorted      = make([]int, sortedLen)
		count       int
	)

	// bug1: l 和 r 错误的从 len 开始了，而不是 len-1
	l, r := len(sortedLeft)-1, len(sortedRight)-1
	for l >= 0 && r >= 0 {
		if sortedLeft[l] > sortedRight[r] {
			count += r + 1
			sorted[sortedIndex] = sortedLeft[l]
			l--
		} else {
			sorted[sortedIndex] = sortedRight[r]
			r--
		}
		sortedIndex--
	}

	// bug2: 直接用了 for i:=0; i<=l; i++{} 来将剩余的元素用 sorted[sortedIndex] 加到 sorted 里，
	// 而忽略了 sortedIndex 是从后往前遍历的，i 则是从前往后遍历的
	for ; l >= 0; l-- {
		sorted[sortedIndex] = sortedLeft[l]
		sortedIndex--
	}

	for ; r >= 0; r-- {
		sorted[sortedIndex] = sortedRight[r]
		sortedIndex--
	}

	return sorted, lc + rc + count
}

// leetcode submit region end(Prohibit modification and deletion)
