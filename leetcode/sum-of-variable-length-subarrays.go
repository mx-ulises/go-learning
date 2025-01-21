func subarraySum(nums []int) int {
	n := len(nums)
	sumatory := make([]int, n+1)
	total := 0
	for i := 0; i < n; i++ {
		sumatory[i+1] = sumatory[i] + nums[i]
		start := max(0, i-nums[i])
		total += sumatory[i+1] - sumatory[start]
	}
	return total
}
