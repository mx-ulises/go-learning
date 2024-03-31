func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	left, right := 0, len(nums)-1
	for left < right {
		if target <= nums[left]+nums[right] {
			right--
		} else {
			count += right - left
			left++
		}
	}
	return count
}
