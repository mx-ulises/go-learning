func waysToSplitArray(nums []int) int {
	right := -nums[0]
	for _, num := range nums {
		right += num
	}
	left := nums[0]
	ways := 0
	if right <= left {
		ways++
	}
	for i := 1; i < len(nums)-1; i++ {
		left += nums[i]
		right -= nums[i]
		if right <= left {
			ways++
		}
	}
	return ways
}
