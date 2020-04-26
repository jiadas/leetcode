package w186

func maxScore2(cardPoints []int, k int) int {
	var max int
	var l, r int
	for i := 0; i < k; i++ {
		l += cardPoints[i]
		r += cardPoints[len(cardPoints)-i-1]
	}

	for k != 0 {
		if l > r || (l == r && cardPoints[0] > cardPoints[len(cardPoints)-1]) {
			max += cardPoints[0]
			l = l - cardPoints[0]
			r = r - cardPoints[len(cardPoints)-k]
			cardPoints = cardPoints[1:]
		} else {
			max += cardPoints[len(cardPoints)-1]
			l = l - cardPoints[k-1]
			r = r - cardPoints[len(cardPoints)-1]
			cardPoints = cardPoints[:len(cardPoints)-1]
		}
		k--
	}
	return max
}
