func minOperations(nums []int, k int) int {
	visited := [101]bool{}
	operations := 0
	for _, num := range nums {
		if num < k {
			return -1
		} else if num > k && visited[num] == false {
			operations++
			visited[num] = true
		}
	}
	return operations
}
