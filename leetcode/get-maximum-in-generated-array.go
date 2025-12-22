package leetcode

func getMaximumGenerated(n int) int {
	nums := [101]int{}
	nums[1] = 1
	maximal := min(n, 1)
	for i := 0; i < n-1; i++ {
		j, k := i*2, i*2+1
		if j <= n {
			nums[j] = nums[i]
			maximal = max(maximal, nums[j])
		}
		if k <= n {
			nums[k] = nums[i] + nums[i+1]
			maximal = max(maximal, nums[k])
		}
	}
	return maximal
}
