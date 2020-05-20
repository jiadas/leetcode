package golang

// https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/solution/mei-ge-yuan-yin-bao-han-ou-shu-ci-de-zui-chang-z-2/
func findTheLongestSubstring(s string) int {
	ans, status := 0, 0
	pos := make([]int, 1<<5)
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	pos[0] = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if pos[status] >= 0 {
			ans = maxInt(ans, i+1-pos[status])
		} else {
			pos[status] = i + 1
		}
	}
	return ans
}

// 暴力枚举，超时啦
//var vowels = []byte{'a', 'e', 'i', 'o', 'u'}
//
//func findTheLongestSubstring(s string) int {
//	vm := make(map[byte][]int, len(vowels))
//	for _, vowel := range vowels {
//		vm[vowel] = nil
//	}
//
//	for i := 0; i < len(s); i++ {
//		if _, ok := vm[s[i]]; ok {
//			vm[s[i]] = append(vm[s[i]], i)
//		}
//	}
//
//	flag := true
//	for _, indexes := range vm {
//		if len(indexes)%2 != 0 {
//			flag = false
//			break
//		}
//	}
//	if flag {
//		return len(s)
//	}
//
//	var result int
//	for _, indexes := range vm {
//		if len(indexes)%2 != 0 {
//			var tmpMax int
//			for _, index := range indexes {
//				var l int
//				if len(s[:index]) > tmpMax {
//					l = findTheLongestSubstring(s[:index])
//				}
//				var r int
//				if index+1 < len(s) && len(s[index+1:]) > tmpMax {
//					r = findTheLongestSubstring(s[index+1:])
//				}
//				tmpMax = maxInt(tmpMax, maxInt(l, r))
//			}
//			result = maxInt(result, tmpMax)
//		}
//	}
//	return result
//}
