func removeDuplicates(nums []int) int {
	twice := false
	tail := 0
	for i := 1; i < len(nums); i++ {
		saveNumber := false
		if nums[i] == nums[tail] && twice == false {
			saveNumber = true
			twice = true
		} else if nums[i] != nums[tail] {
			saveNumber = true
			twice = false
		}
		if saveNumber {
			tail += 1
			nums[tail] = nums[i]
		}
	}
	return tail + 1
}
