package main

import (
	"fmt"
	"sort"
)

type pointSlice [][]int

func (ps pointSlice) Len() int {
	return len(ps)
}

func (ps pointSlice) Less(i, j int) bool {
	return ps[i][1] < ps[j][1]
}

func (ps pointSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(pointSlice(points))
	arrowPt := points[0][1]
	arrowCount := 1
	for i := 0; i < len(points); i++ {
		if arrowPt >= points[i][0] {
			continue
		}
		arrowCount++
		arrowPt = points[i][1]
	}
	return arrowCount
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(points)
	fmt.Println(findMinArrowShots(points))
}
