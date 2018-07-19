package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	dsu := newDSU(len(edges))
	for _, e := range edges {
		if dsu.find(e[0]) == dsu.find(e[1]) {
			return e
		}
		dsu.union(e[0], e[1])
	}
	return []int{-1, -1}
}

type DSU struct {
	parent []int
	rank   []int
}

// https://www.geeksforgeeks.org/union-find-algorithm-set-2-union-by-rank/
func newDSU(n int) DSU {
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	return DSU{
		parent: parent,
		rank:   make([]int, n+1),
	}
}

// A utility function to find set of an element i
// (uses path compression technique)
func (dsu *DSU) find(v int) int {
	// find root and make root as parent of v (path compression)
	if dsu.parent[v] != v {
		dsu.parent[v] = dsu.find(dsu.parent[v])
	}
	return dsu.parent[v]
}

// A function that does union of two sets of x and y
// (uses union by rank)
func (dsu *DSU) union(x, y int) {
	xp := dsu.find(x)
	yp := dsu.find(y)
	if xp != yp {
		switch {
		// Attach smaller rank tree under root of high rank tree
		// (Union by Rank)
		case dsu.rank[xp] > dsu.rank[yp]:
			dsu.parent[yp] = xp
		case dsu.rank[xp] < dsu.rank[yp]:
			dsu.parent[xp] = yp
		// If ranks are same, then make one as root and increment
		// its rank by one
		case dsu.rank[xp] == dsu.rank[yp]:
			dsu.parent[xp] = yp
			dsu.rank[yp] += 1
		}
	}
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}
