package main

import (
	"fmt"
	"sort"
)

type pairSlice [][]int

func (ps pairSlice) Len() int {
	return len(ps)
}

func (ps pairSlice) Less(i, j int) bool {
	return ps[i][0] < ps[j][0]
}

func (ps pairSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	if n == 0 {
		return 0
	}
	dp := make([]int, n) // dp[i] 表示以第 i 对数字为结尾的 chain 的最大长度
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	sort.Sort(pairSlice(pairs))
	fmt.Println(pairs)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				tmp := dp[j] + 1
				if tmp > dp[i] {
					dp[i] = tmp
				}
			}
		}
	}
	var result int
	for _, value := range dp {
		if value > result {
			result = value
		}
	}
	return result
}

func main() {
	fmt.Println(findLongestChain([][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}}))
	fmt.Println(findLongestChain([][]int{{9, 10}, {9, 10}, {4, 5}, {-9, -3}, {-9, 1}, {0, 3}, {6, 10}, {-5, -4}, {-7, -6}}))
}
