func GetMax(nums []int) int {
	maximal := 0
	for _, num := range nums {
		if maximal < num {
			maximal = num
		}
	}
	return maximal
}

func countSubarrays(nums []int, k int) int64 {
	var count int64
	max, maxCount := GetMax(nums), 0
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] == max {
			maxCount++
		}
		right++
		candidates := int64(1 + n - right)
		for maxCount == k && left < right {
			count += candidates
			if nums[left] == max {
				maxCount--
			}
			left++
		}
	}
	return count
}
