package golang

// https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2481/Share-my-O(log(min(mn)))-solution-with-explanation
// https://zhuanlan.zhihu.com/p/70654378
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	m, n := len(nums1), len(nums2)
	if m > n {
		A, B, m, n = B, A, n, m
	}
	if n == 0 {
		return 0
	}
	imin, imax, halfLen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < m && B[j-1] > A[i] {
			// i 的值太小，增加它
			imin = i + 1
		} else if i > 0 && A[i-1] > B[j] {
			// 	i 的值过大，减小它
			imax = i - 1
		} else {
			// 	i is perfect
			var maxLeft int
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = maxInt(A[i-1], B[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == m {
				minRight = B[j]
			} else if j == n {
				minRight = A[i]
			} else {
				minRight = minInt(A[i], B[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
