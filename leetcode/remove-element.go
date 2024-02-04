func removeElement(nums []int, val int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		if nums[start] == val {
			nums[start] = nums[end]
			nums[end] = val
			end--
		} else {
			start++
		}
	}
	return start
}
