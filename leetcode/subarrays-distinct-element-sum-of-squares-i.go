func sumCounts(nums []int) int {
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		set := map[int]bool{}
		for j := i; j < n; j++ {
			set[nums[j]] = true
			count += len(set) * len(set)
		}
	}
	return count
}
