package golang

// leetcode submit region begin(Prohibit modification and deletion)
// 最终决定用官方 solution 里的 golang example，而不是剑指 offer 51 里的归并模板
// 原因 1: 不用显示传递 lower 和 upper，函数声明更简洁
// 原因 2: 一次 for 循环实现有序数字的合并
func countRangeSum(nums []int, lower int, upper int) int {
	// 先声明函数类型，为了能在 mergeCount 这个匿名函数里实现递归调用自己，即 mergeCount
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		// 问题 1: 为什么 n<=1 要 return 0
		// 解答 1: 问题转化为在前缀和数组中求满足条件的下标对(i,j)，其中 i!=j，sums[3]-sums[3] 永远等于 0
		//
		// 问题 2: [0,0],[2,2]这种区间和怎么在代码中体现的?
		// 解答 2: [0,0] => sums[1]-sums[0]；[2,2] => sums[3]-sums[2]
		if n <= 1 {
			return 0
		}

		mid := n >> 1
		n1 := append([]int(nil), arr[:mid]...)
		n2 := append([]int(nil), arr[mid:]...)
		// 递归完毕后，n1 和 n2 均为有序，因为下面合并的时候，是用下标操作的 arr
		cnt := mergeCount(n1) + mergeCount(n2)

		// 统计下标对的数量
		var l, r int
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		// n1 和 n2 归并填入 arr
		var p1, p2 int
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}

		return cnt
	}

	// 构建前缀和
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	return mergeCount(sums)
}

// leetcode submit region end(Prohibit modification and deletion)
