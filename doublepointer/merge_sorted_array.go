package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	index := m + n - 1
	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[index] = nums2[j]
			j--
		} else if j < 0 {
			nums1[index] = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 11, 0, 0, 0, 0, 0}
	merge(nums1, 4, []int{4, 5, 6, 7, 8}, 5)
	fmt.Println(nums1)
}
