func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	minimal := float64(nums[end])
	for start < end {
		minimal = min(minimal, float64(nums[start]+nums[end])/2.0)
		start++
		end--
	}
	return minimal
}
