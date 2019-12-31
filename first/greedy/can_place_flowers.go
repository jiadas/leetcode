package main

import "fmt"

// change the flowerbed
func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 || len(flowerbed) < n {
		return false
	}
	if n == 0 || (len(flowerbed) == 1 && flowerbed[0] == 0 && n == 1) {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if i == 0 && i+1 < len(flowerbed) {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
			}
		} else if i == len(flowerbed)-2 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
			}
		} else if i > 0 && i < len(flowerbed)-2 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
				n--
				flowerbed[i+1] = 1
				//i++
			}
		}

		if n == 0 {
			return true
		}

	}

	return false
}

// 更清晰的解决方案
func canPlaceFlowers1(flowerbed []int, n int) bool {
	var cnt int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		var pre int
		if i != 0 {
			pre = flowerbed[i-1]
		}

		var next int
		if i != len(flowerbed)-1 {
			next = flowerbed[i+1]
		}

		if pre == 0 && next == 0 {
			cnt++
			flowerbed[i] = 1
		}
	}
	return cnt >= n
}

// unchange the flowerbed
func canPlaceFlowers2(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 || len(flowerbed) < n {
		return false
	}
	if n == 0 || (len(flowerbed) == 1 && flowerbed[0] == 0 && n == 1) {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if i == 0 && i+1 < len(flowerbed) {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				n--
			}
		} else if i == len(flowerbed)-2 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				n--
			}
		} else if i > 0 && i < len(flowerbed)-2 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
				n--
				i++
			}
		}

		if n == 0 {
			return true
		}

	}

	return false
}

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers1([]int{0, 0}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers2([]int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, 4))
}
