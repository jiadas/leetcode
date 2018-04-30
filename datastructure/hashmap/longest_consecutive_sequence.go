package main

import "fmt"

// We will use HashMap. The key thing is to keep track of the sequence length and store that in the boundary points of the sequence. For example, as a result, for sequence {1, 2, 3, 4, 5}, map.get(1) and map.get(5) should both return 5.
//
// Whenever a new element n is inserted into the map, do two things:
//
// See if n - 1 and n + 1 exist in the map, and if so, it means there is an existing sequence next to n. Variables left and right will be the length of those two sequences, while 0 means there is no sequence and n will be the boundary point later. Store (left + right + 1) as the associated value to key n into the map.
// Use left and right to locate the other end of the sequences to the left and right of n respectively, and replace the value with the new length.
// Everything inside the for loop is O(1) so the total time is O(n). Please comment if you see something wrong. Thanks.
func longestConsecutive(nums []int) int {
	m := make(map[int]int, len(nums))
	var ret int
	for _, value := range nums {
		if _, ok := m[value]; !ok {
			left := m[value-1]
			right := m[value+1]

			sum := left + right + 1
			m[value] = sum

			if sum > ret {
				ret = sum
			}

			m[value-left] = sum
			m[value+right] = sum
		}
	}
	return ret
}

func main() {
	fmt.Println(longestConsecutive([]int{1, 2, 3, 4, 5, 7, 8, 11, 12, 13}))
}
