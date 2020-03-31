package golang

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					count++
					break
				}
			}
		} else {
			count++
		}
	}
	return count
}
