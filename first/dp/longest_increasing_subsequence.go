package main

import "fmt"

func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums)) // lis[n] 表示以 Sn 结尾的序列的最长递增子序列长度
	for i := 0; i < len(nums); i++ {
		max := 1 // 因为在求 lis[n] 时可能无法找到一个满足条件的递增子序列，此时 {Sn} 就构成了递增子序列，令 dp[n] 最小为 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lis[j]+1 > max {
					max = lis[j] + 1
				}
			}
		}
		lis[i] = max
	}

	var result int
	for _, value := range lis {
		if value > result {
			result = value
		}
	}
	return result
}

// tails is an array storing the {{smallest}} tail of all increasing subsequences with length i+1 in tails[i].
// For example, say we have nums = [4,5,6,3], then all the available increasing subsequences are:
//
// len = 1   :      [4], [5], [6], [3]   => tails[0] = 3
// len = 2   :      [4, 5], [5, 6]       => tails[1] = 5
// len = 3   :      [4, 5, 6]            => tails[2] = 6
// We can easily prove that tails is a increasing array. Therefore it is possible to do a binary search in tails array to find the one needs update.
//
// Each time we only do one of the two:
//
// (1) if x is larger than all tails, append it, increase the size by 1
// (2) if tails[i-1] < x <= tails[i], update tails[i]
// Doing so will maintain the tails invariant. The the final answer is just the size.
func lengthOfLISBinary(nums []int) int {
	tails := make([]int, len(nums))
	var size int
	for _, value := range nums {
		l, h := 0, size // 令 h = size 是为了当 value 没有在 tails 里时，正确返回 value 应该正确插入的位置
		for l < h {
			m := l + (h-l)/2
			if tails[m] >= value {
				h = m
			} else {
				l = m + 1
			}
		}
		tails[l] = value
		if l == size {
			size++
		}
	}
	return size
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISBinary([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISBinary([]int{2, 2}))
}
