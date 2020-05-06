package golang

// https://leetcode-cn.com/problems/minimum-cost-for-tickets/solution/zui-di-piao-jie-by-leetcode-solution/
func mincostTickets(days []int, costs []int) int {
	daySet := make(map[int]struct{}, len(days))
	for _, day := range days {
		daySet[day] = struct{}{}
	}
	tp := TravelPlanner{
		costs:  costs,
		daySet: daySet,
		memo:   make([]int, 366),
	}
	return tp.dpCost(1)
}

type TravelPlanner struct {
	costs  []int
	daySet map[int]struct{}
	memo   []int
}

func (tp *TravelPlanner) dpCost(i int) int {
	if i > 365 {
		return 0
	}
	if tp.memo[i] != 0 {
		return tp.memo[i]
	}
	_, ok := tp.daySet[i]
	if ok {
		tp.memo[i] = minInt(tp.dpCost(i+1)+tp.costs[0], minInt(tp.dpCost(i+7)+tp.costs[1], tp.dpCost(i+30)+tp.costs[2]))
	} else {
		tp.memo[i] = tp.dpCost(i + 1)
	}
	return tp.memo[i]
}
