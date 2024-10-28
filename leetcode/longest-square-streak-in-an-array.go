func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	maximal := 0
	visited := map[int]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[len(nums)-i-1]
		visited[num] = visited[num*num] + 1
		maximal = max(maximal, visited[num])
	}
	if maximal < 2 {
		return -1
	}
	return maximal
}
