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

func (dsu *DSU) find(v int) int {
	if dsu.parent[v] != v {
		dsu.parent[v] = dsu.find(dsu.parent[v])
	}
	return dsu.parent[v]
}

func (dsu *DSU) union(x, y int) {
	xp := dsu.find(x)
	yp := dsu.find(y)
	if xp != yp {
		switch {
		case dsu.rank[xp] > dsu.rank[yp]:
			dsu.parent[yp] = xp
		case dsu.rank[xp] < dsu.rank[yp]:
			dsu.parent[xp] = yp
		case dsu.rank[xp] == dsu.rank[yp]:
			dsu.parent[xp] = yp
			dsu.rank[yp] += 1
		}
	}
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}
