package main

import "fmt"

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	var ret int
	for key, c1 := range m {
		// 一定要判断 key+1 是否在 map 中
		c2, ok := m[key+1]
		if ok {
			if c1+c2 > ret {
				ret = c1 + c2
			}
		}
	}
	return ret
}

func main() {
	//fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
}
