package golang

// https://leetcode-cn.com/problems/subarray-sum-equals-k/solution/he-wei-kde-zi-shu-zu-by-leetcode-solution/
// 方法一：枚举
func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 方法二：前缀和 + 哈希表优化
// 定义 pre[i] 为 [0..i] 里所有数的和
// 考虑以 i 结尾的和为 k 的连续子数组个数时，只要统计有多少个前缀和为 pre[i] − k 的 pre[j] 即可
func subarraySum2(nums []int, k int) int {
	var count, pre int
	mp := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre]++
	}
	return count
}