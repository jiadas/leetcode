package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	l, h := 0, len(nums)-1
	for l < h {
		pos := partitionNums(nums, l, h)
		if pos == k-1 {
			return nums[pos]
		} else if pos < k-1 {
			l = pos + 1
		} else {
			h = pos - 1
		}
	}
	return nums[k-1]
}

// 降序排列 nums
func partitionNums(nums []int, l, h int) int {
	p := nums[l]
	for l < h {
		// 从右往左，找到第一个大于 p 的位置
		for l < h && nums[h] <= p {
			h--
		}
		if l < h {
			nums[l] = nums[h]
			l++
		}

		// 从左往右，找到第一个小于或等于 p 的位置
		for l < h && nums[l] > p {
			l++
		}
		if l < h {
			nums[h] = nums[l]
			h--
		}
	}
	nums[l] = p
	return l
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
