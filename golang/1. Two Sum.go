package golang

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		j, ok := m[target-num]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
