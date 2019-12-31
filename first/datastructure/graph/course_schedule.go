package main

// This problem is equivalent to finding if a cycle exists in a directed graph.
// If a cycle exists, no topological ordering exists and therefore it will be impossible to take all courses.
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	visited, onpath := make([]bool, numCourses), make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if !visited[i] && dfsCycle(graph, visited, onpath, i) {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/course-schedule/discuss/58509/18-22-lines-C++-BFSDFS-Solutions
// For DFS, it will first visit a node, then one neighbor of it, then one neighbor of this neighbor... and so on.
// If it meets a node which was visited in the current process of DFS visit, a cycle is detected and we will return false.
// Otherwise it will start from another unvisited node and repeat this process till all the nodes have been visited.
// Note that you should make two records:
// 	one is to record all the visited nodes and the other is to record the visited nodes in the current DFS visit.
//
// We use a []bool visited to record all the visited nodes and another []bool onpath to record the visited nodes of the current DFS visit.
// Once the current visit is finished, we reset the onpath value of the starting node to false.
func dfsCycle(graph [][]int, visited, onpath []bool, node int) bool {
	if visited[node] {
		return false
	}
	visited[node] = true
	onpath[node] = true
	for _, v := range graph[node] {
		if onpath[v] || dfsCycle(graph, visited, onpath, v) {
			return true
		}
	}
	onpath[node] = false
	return false
}
