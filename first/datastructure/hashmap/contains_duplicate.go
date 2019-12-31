package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, value := range nums {
		_, ok := m[value]
		if ok {
			return true
		}
		m[value] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 1}))
}
