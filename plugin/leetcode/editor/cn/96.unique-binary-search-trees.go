package golang

// leetcode submit region begin(Prohibit modification and deletion)
var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return countWithRange(1, n)
}

// https://mp.weixin.qq.com/s/kcwz2lyRxxOsC3n11qdVSw
// 常规动态规划，利用二维数组消除重叠子问题
// 更容易理解
func countWithRange(lo, hi int) int {
	if lo > hi {
		return 1
	}

	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}

	var count int
	for i := lo; i <= hi; i++ {
		left := countWithRange(lo, i-1)
		right := countWithRange(i+1, hi)
		count += left * right
	}

	memo[lo][hi] = count

	return count
}

// https://leetcode.cn/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
// 也是动态规划，但是有一个复杂的推到过程，代码量虽然更少，但是理解难度要大于上面
func doNumTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

// leetcode submit region end(Prohibit modification and deletion)
