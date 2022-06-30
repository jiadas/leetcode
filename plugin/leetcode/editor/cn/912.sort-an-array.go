package golang

import (
	"math/rand"
	"time"
)

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	return heapSort(nums)
}

// 合并排序
func mergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	n1 := sortArray(append([]int(nil), nums[:n/2]...))
	n2 := sortArray(append([]int(nil), nums[n/2:]...))

	var p1, p2 int
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}

	return nums
}

// 快速排序
func doRandomizedQuickSort(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	randomizedQuickSort(nums, 0, len(nums)-1)
	return nums
}

func randomizedQuickSort(nums []int, l, r int) {
	if l < r {
		p := rand.Intn(r-l+1) + l
		nums[p], nums[r] = nums[r], nums[p]
		pivot := nums[r]
		i := l - 1
		for j := l; j <= r-1; j++ {
			if nums[j] < pivot {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		i++
		nums[i], nums[r] = nums[r], nums[i]
		randomizedQuickSort(nums, l, i-1)
		randomizedQuickSort(nums, i+1, r)
	}
}

// 堆排序
func heapSort(nums []int) []int {
	initHeap(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		down(nums, 0, i)
	}

	return nums
}

func initHeap(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
}

func down(nums []int, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && nums[j1] < nums[j2] {
			j = j2 // right child
		}
		if nums[j] <= nums[i] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

// leetcode submit region end(Prohibit modification and deletion)
