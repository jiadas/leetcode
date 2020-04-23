package golang

func guessNumber(n int) int {
	left, right := 0, n
	var mid int
	for left <= right {
		mid = left + (right-left)>>1
		switch guess(mid) {
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		case 0:
			return mid
		}
	}
	return -1
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return 0
}
