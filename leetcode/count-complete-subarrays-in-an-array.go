func getUniqueCount(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	return len(numSet)
}

func countCompleteSubarrays(nums []int) int {
	uniqueCount := getUniqueCount(nums)
	numCount := map[int]int{}
	count, num := 0, -1
	left, right, n := 0, 0, len(nums)
	for right < n {
		num = nums[right]
		numCount[num]++
		if numCount[num] == 1 {
			uniqueCount--
		}
		for left <= right && uniqueCount == 0 {
			num = nums[left]
			count += n - right
			numCount[num]--
			if numCount[num] == 0 {
				uniqueCount++
			}
			left++
		}
		right++
	}
	return count
}
