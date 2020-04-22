package golang

func mySqrt(x int) int {
	left, right := 0, x
	var mid, square int
	for left <= right {
		mid = left + (right-left)>>1
		square = mid * mid
		if square > x {
			right = mid - 1
		} else if square < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left - 1
}
