package golang

// leetcode submit region begin(Prohibit modification and deletion)
// 吊打官方题解：
// 1、如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点
// 2、可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新
// 3、如果可以一直跳到最后，就成功了
func canJump(nums []int) bool {
	var k int
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = maxInt(k, i+nums[i])
	}
	return true
}

func canJumpDP(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	// 状态：f[i] 表示是否能从 0 跳到 i
	f := make([]bool, n)
	f[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 推导：f[i] = AND(f[j],j<i&&j能跳到i) 判断之前所有的点最后一跳是否能跳到当前点
			if f[j] && j+nums[j] >= i {
				f[i] = true
			}
		}
	}

	return f[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
