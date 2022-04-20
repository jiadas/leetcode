package golang

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	sums []int
}

func Constructor2(nums []int) NumArray {
	// 将前缀和数组 sums 的长度设为 n+1 的目的是为了方便计算 sumRange(i,j)，不需要对 i=0 的情况特殊处理
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		// 假设数组 nums 的长度为 n，创建长度为 n+1 的前缀和数组 sums，
		// 对于 0≤i<n 都有 sums[i+1]=sums[i]+nums[i]，则当 0<i≤n 时，sums[i] 表示数组 nums 从下标 0 到下标 i−1 的前缀和
		sums[i+1] = sums[i] + num
	}
	return NumArray{sums: sums}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.sums[right+1] - na.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// leetcode submit region end(Prohibit modification and deletion)
