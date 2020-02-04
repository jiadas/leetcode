package golang

func isValid(s string) bool {
	left := map[byte]struct{}{
		'(': {},
		'{': {},
		'[': {},
	}
	right := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, ok := left[b]; ok {
			stack = append(stack, b)
		} else {
			var l byte
			if len(stack) > 0 {
				// pop
				l, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if right[b] != l {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
