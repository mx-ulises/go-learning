func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findDuplicate(nums []int) int {
	repeated := -1
	for i := 0; i < len(nums); i++ {
		num := abs(nums[i]) - 1
		if nums[num] < 0 {
			repeated = num + 1
		} else {
			nums[num] = -nums[num]
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = abs(nums[i])
	}
	return repeated
}
