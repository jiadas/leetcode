package golang

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	val   int
	index int
}

func countSmaller(nums []int) []int {
	l := len(nums)
	count := make([]int, l)
	pairs := make([]pair, l)
	for i, num := range nums {
		pairs[i] = pair{
			val:   num,
			index: i,
		}
	}

	doCountSmaller(pairs, &count)

	return count
}

func doCountSmaller(pairs []pair, count *[]int) []pair {
	pairsLen := len(pairs)
	if pairsLen <= 1 {
		return pairs
	}

	mid := pairsLen >> 1
	sortedLeft := doCountSmaller(pairs[:mid], count)
	sortedRight := doCountSmaller(pairs[mid:], count)

	var (
		sortedIndex = pairsLen - 1
		sortedPairs = make([]pair, pairsLen)
	)

	l, r := len(sortedLeft)-1, len(sortedRight)-1
	for l >= 0 && r >= 0 {
		if sortedLeft[l].val > sortedRight[r].val {
			(*count)[sortedLeft[l].index] += r + 1
			sortedPairs[sortedIndex] = sortedLeft[l]
			l--
		} else {
			sortedPairs[sortedIndex] = sortedRight[r]
			r--
		}
		sortedIndex--
	}

	for ; l >= 0; l-- {
		sortedPairs[sortedIndex] = sortedLeft[l]
		sortedIndex--
	}

	for ; r >= 0; r-- {
		sortedPairs[sortedIndex] = sortedRight[r]
		sortedIndex--
	}

	return sortedPairs
}

// leetcode submit region end(Prohibit modification and deletion)

// 超时
func countSmallerTO(nums []int) []int {
	l := len(nums)
	count := make([]int, l)
	pairs := make([]pair, l)
	for i, num := range nums {
		pairs[i] = pair{
			val:   num,
			index: i,
		}
	}

	_, ans := doCountSmallerTO(pairs, count)

	return ans
}

func doCountSmallerTO(pairs []pair, count []int) ([]pair, []int) {
	// 超时大概是因为重复分配 int slice 造成的
	result := make([]int, len(count))

	pairsLen := len(pairs)
	if pairsLen <= 1 {
		return pairs, count
	}

	mid := pairsLen >> 1
	sortedLeft, lc := doCountSmallerTO(pairs[:mid], count)
	sortedRight, rc := doCountSmallerTO(pairs[mid:], count)

	var (
		sortedIndex = pairsLen - 1
		sortedPairs = make([]pair, pairsLen)
	)

	l, r := len(sortedLeft)-1, len(sortedRight)-1
	for l >= 0 && r >= 0 {
		if sortedLeft[l].val > sortedRight[r].val {
			result[sortedLeft[l].index] += r + 1
			sortedPairs[sortedIndex] = sortedLeft[l]
			l--
		} else {
			sortedPairs[sortedIndex] = sortedRight[r]
			r--
		}
		sortedIndex--
	}

	for ; l >= 0; l-- {
		sortedPairs[sortedIndex] = sortedLeft[l]
		sortedIndex--
	}

	for ; r >= 0; r-- {
		sortedPairs[sortedIndex] = sortedRight[r]
		sortedIndex--
	}

	for i := range count {
		result[i] = result[i] + lc[i] + rc[i]
	}

	return sortedPairs, result
}
