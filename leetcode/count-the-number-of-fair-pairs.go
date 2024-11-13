func countFairPairs(nums []int, lower int, upper int) int64 {
	pairs := 0
	n := len(nums)
	i, left, right := 0, n-1, n-1
	sort.Ints(nums)
	for i < right {
		// get first position we can meet lower using nums[i]
		for i < left && lower <= (nums[i]+nums[left]) {
			left--
		}
		left++
		// get last position we can meet upper using nums[i]
		for left <= right && upper < (nums[i]+nums[right]) {
			right--
		}
		// validate the range
		if left <= right {
			// if the inclusive range is valid, add the pairs
			pairs += right - left + 1
		} else {
			// if it is not valid, then reset left to right, so we can use it in the next iteration
			left = right
		}
		i++
	}
	return int64(pairs)
}
