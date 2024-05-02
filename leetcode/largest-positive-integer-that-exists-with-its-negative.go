func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findMaxK(nums []int) int {
	visited := map[int]bool{}
	maximal := -1
	for _, num := range nums {
		if visited[-num] {
			maximal = max(maximal, abs(num))
		}
		visited[num] = true
	}
	return maximal
}
