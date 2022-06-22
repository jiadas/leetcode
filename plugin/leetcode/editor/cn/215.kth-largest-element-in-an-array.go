package golang

import (
	"math/rand"
	"time"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	lo, hi := 0, len(nums)-1
	k = len(nums) - k
	for lo <= hi {
		p := randomizedPartition(nums, lo, hi)
		switch {
		case p < k:
			lo = p + 1
		case p > k:
			hi = p - 1
		default:
			return nums[p]
		}
	}
	return -1
}

func randomizedPartition(nums []int, lo, hi int) int {
	p := rand.Intn(hi-lo+1) + lo
	nums[hi], nums[p] = nums[p], nums[hi]
	pivot := nums[hi]
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

// leetcode submit region end(Prohibit modification and deletion)
