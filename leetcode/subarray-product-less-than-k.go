func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	n := len(nums)
	left, right := 0, 0
	current := 1
	for left < n {
		// build largest subset that multiplies less than k
		for right < n && current*nums[right] < k {
			current *= nums[right]
			right++
		}
		// We find an element that doesn't fit (or end of list), remove elements
		for left < right && right < n && k <= current*nums[right] {
			current /= nums[left]
			count += right - left
			left++
		}
		if right < n && current*nums[right] < k {
			// we can fit new number in current multiplication
			current *= nums[right]
		} else if left == right {
			// Current number is too big, we need to skip it
			left++
		} else {
			// End of the list, we need to count all subsets
			for left < n {
				count += right - left
				left++
			}
		}
		right++
	}
	return count
}
