package golang

// https://leetcode.com/problems/integer-to-roman/discuss/6274/Simple-Solution
func intToRomanFD(num int) string {
	M := []string{"", "M", "MM", "MMM"}                                       // 1000
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} // 100
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} // 10
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"} // 1
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

// 自己写的穷举，没有 Discuss 里的简洁。。。
func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}

	var result string
	for i := 1000; i > 0; i /= 10 {
		x := num / i
		if x > 0 {
			tmp := x * i
			result += allCases(tmp)
		}
		num %= i
	}
	return result
}

func allCases(num int) string {
	var r string
	switch {
	case num < 4:
		for i := 0; i < num; i++ {
			r += "I"
		}
	case num == 4:
		return "IV"
	case num == 5:
		return "V"
	case num > 5 && num < 9:
		r += "V"
		for i := 5; i < num; i++ {
			r += "I"
		}
	case num == 9:
		return "IX"
	case num == 10:
		return "X"
	case num > 10 && num < 40:
		for i := 0; i < num; i += 10 {
			r += "X"
		}
	case num == 40:
		return "XL"
	case num == 50:
		return "L"
	case num > 50 && num < 90:
		r += "L"
		for i := 50; i < num; i += 10 {
			r += "X"
		}
	case num == 90:
		return "XC"
	case num == 100:
		return "C"
	case num > 100 && num < 400:
		for i := 0; i < num; i += 100 {
			r += "C"
		}
	case num == 400:
		return "CD"
	case num == 500:
		return "D"
	case num > 500 && num < 900:
		r += "D"
		for i := 500; i < num; i += 100 {
			r += "C"
		}
	case num == 900:
		return "CM"
	case num == 1000:
		return "M"
	case num > 1000 && num < 3999:
		for i := 0; i < num; i += 1000 {
			r += "M"
		}
	}
	return r
}
