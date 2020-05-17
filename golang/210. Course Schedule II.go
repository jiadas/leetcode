package golang

// https://leetcode-cn.com/problems/course-schedule-ii/solution/ke-cheng-biao-ii-by-leetcode-solution/
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  = make([]int, 0, numCourses)
		invalid bool
		dfs     func(int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if invalid {
					return
				}
			} else if visited[v] == 1 {
				invalid = true
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
	}

	for i := 0; i < numCourses && !invalid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if invalid {
		return nil
	}

	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}
	return result
}
