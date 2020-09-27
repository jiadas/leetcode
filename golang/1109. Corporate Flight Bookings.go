package golang

// 差分数组
// https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/cha-fen-ji-qiao
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		if len(booking) < 3 {
			continue
		}
		i := booking[0]
		j := booking[1]
		k := booking[2]
		if i >= 0 && i <= n {
			diff[i] += k
		}
		if j >= 0 && j < n {
			diff[j+1] -= k
		}
	}
	answer := make([]int, n+1)
	for i := 1; i <= n; i++ {
		answer[i] = answer[i-1] + diff[i]
	}
	return answer[1:]
}
