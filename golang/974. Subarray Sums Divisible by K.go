package golang

// https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/solution/he-ke-bei-k-zheng-chu-de-zi-shu-zu-by-leetcode-sol/
func subarraysDivByK(A []int, K int) int {
	records := map[int]int{0: 1}
	var sum, ans int
	for _, num := range A {
		sum += num
		// 注意 Go 取模的特殊性，当被除数为负数时取模结果为负数，需要纠正
		mod := (sum%K + K) % K
		ans += records[mod]
		records[mod]++
	}
	return ans
}
