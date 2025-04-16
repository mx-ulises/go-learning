func countGood(nums []int, k int) int64 {
	count, pairs := 0, 0
	left, right, n := 0, 0, len(nums)
	prevNums := map[int]int{}
	for right < n {
		num := nums[right]
		pairs += prevNums[num]
		prevNums[num]++
		for k <= pairs {
			count += n - right
			aux := nums[left]
			left++
			prevNums[aux]--
			pairs -= prevNums[aux]
		}
		right++
	}
	return int64(count)
}
