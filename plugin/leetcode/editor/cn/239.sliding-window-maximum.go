package golang

import "container/heap"

// leetcode submit region begin(Prohibit modification and deletion)

type Item struct {
	val   int
	index int
}

type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	h := &ItemHeap{}
	heap.Init(h)

	n := len(nums)
	ans := make([]int, n-k+1)
	var ansIndex int
	for right := 0; right < n; right++ {
		// remove numbers out of range k
		for h.Len() != 0 && (*h)[0].index < right-k+1 {
			heap.Pop(h)
		}

		// remove smaller numbers in k range as they are useless
		for h.Len() != 0 && nums[right] > (*h)[h.Len()-1].val {
			heap.Remove(h, h.Len()-1)
		}

		heap.Push(h, Item{val: nums[right], index: right})

		if right >= k-1 {
			ans[ansIndex] = (*h)[0].val
			ansIndex++
		}
	}

	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 在最后一个用例失败了
func doMaxSlidingWindowButFail(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)

	n := len(nums)
	ans := make([]int, n-k+1)
	var right, ansIndex int
	for right < n {
		// remove smaller numbers in k range as they are useless
		for h.Len() != 0 && nums[right] > (*h)[h.Len()-1] {
			heap.Remove(h, h.Len()-1)
		}
		heap.Push(h, nums[right])
		right++
		if right >= k {
			ans[ansIndex] = (*h)[0]
			ansIndex++
			if (*h)[0] == nums[right-k] {
				heap.Pop(h)
			}
		}
	}

	return ans
}
