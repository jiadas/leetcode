package main

import (
	"fmt"
	"sort"
)

type peopleSlice [][]int

func (ps peopleSlice) Len() int {
	return len(ps)
}

func (ps peopleSlice) Less(i, j int) bool {
	if ps[i][0] == ps[j][0] {
		return ps[i][1] < ps[j][1]
	}
	return ps[i][0] > ps[j][0]
}

func (ps peopleSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 || len(people[0]) == 0 {
		return nil
	}

	sort.Sort(peopleSlice(people))

	result := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		if result[index] == nil {
			result[index] = people[i]
		} else {
			copy(result[index+1:], result[index:])
			result[index] = people[i]
		}
	}

	return result
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}

	fmt.Println(reconstructQueue(people))
}
