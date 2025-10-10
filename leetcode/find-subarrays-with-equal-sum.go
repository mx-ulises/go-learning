package leetcode

func findSubarrays(nums []int) bool {
	current := nums[0] + nums[1]
	visited := map[int]bool{current: true}
	for i := 1; i < len(nums)-1; i++ {
		current += nums[i+1] - nums[i-1]
		if visited[current] {
			return true
		}
		visited[current] = true
	}
	return false

}
