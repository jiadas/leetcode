package main

import "fmt"

// https://leetcode.com/problems/is-graph-bipartite/discuss/115487/Java-Clean-DFS-solution-with-Explanation
//
// Our goal is trying to use two colors to color the graph and see if there are any adjacent nodes having the same color.
// Initialize a color[] array for each node. Here are three states for colors[] array:
//   -1: Haven't been colored.
//   0: Blue.
//   1: Red.
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		colors[i] = -1
	}

	for i := 0; i < n; i++ {
		if colors[i] == -1 && !validColor(graph, colors, i, 0) {
			return false
		}
	}
	return true
}

// For each node,
//   If it hasn't been colored, use a color to color it. Then use the other color to color all its adjacent nodes (DFS).
//   If it has been colored, check if the current color is the same as the color that is going to be used to color it.
func validColor(graph [][]int, colors []int, node, color int) bool {
	if colors[node] != -1 {
		return colors[node] == color
	}

	colors[node] = color
	for _, v := range graph[node] {
		if !validColor(graph, colors, v, 1-color) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}
