package main

// https://leetcode.com/problems/course-schedule-ii/discuss/59317/Two-AC-solution-in-Java-using-BFS-and-DFS-with-explanation
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	incomeEdgeCount := make([]int, numCourses)
	for _, p := range prerequisites {
		incomeEdgeCount[p[0]] += 1
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	return solveByBFS(graph, incomeEdgeCount)
}

// We observe that if a node has incoming edges, it has prerequisites.
// Therefore, the first few in the order must be those with no prerequisites, i.e. no incoming edges.
// Any non-empty DAG must have at least one node without incoming links.
// If we visit these few and remove all edges attached to them, we are left with a smaller DAG, which is the same problem.
// This will then give our BFS solution.
func solveByBFS(graph [][]int, incomeEdgeCount []int) []int {
	var toVisit, order []int

	for i, c := range incomeEdgeCount {
		if c == 0 {
			toVisit = append(toVisit, i)
		}
	}

	for len(toVisit) != 0 {
		v := toVisit[0]
		toVisit = toVisit[1:]

		order = append(order, v)

		for _, e := range graph[v] {
			incomeEdgeCount[e]--
			if incomeEdgeCount[e] == 0 {
				toVisit = append(toVisit, e)
			}
		}
	}

	if len(order) != len(incomeEdgeCount) {
		return []int{}
	}

	return order
}

// https://www.cs.usfca.edu/~galles/visualization/TopoSortDFS.html
// https://www.geeksforgeeks.org/topological-sorting/
// https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
func solveByDFS(graph [][]int) []int {
	n := len(graph)
	visited, onPath := make([]bool, n), make([]bool, n)

	var order []int

	for i := 0; i < n; i++ {
		if !visited[i] {
			has, o := hasOrder(graph, visited, onPath, i)
			if !has {
				return []int{}
			}
			order = append(o, order...)
		}
	}
	return order
}

func hasOrder(graph [][]int, visited, onPath []bool, v int) (bool, []int) {
	var order []int
	visited[v] = true
	onPath[v] = true
	for _, value := range graph[v] {
		if onPath[value] {
			return false, []int{}
		}
		if !visited[value] {
			has, o := hasOrder(graph, visited, onPath, value)
			if !has {
				return false, []int{}
			}
			order = append(o, order...)
		}
	}
	order = append([]int{v}, order...)
	onPath[v] = false
	return true, order
}
