func arithmeticTriplets(nums []int, diff int) int {
	visited := map[int]bool{}
	tripletCount := 0
	for _, num := range nums {
		second := num - diff
		first := second - diff
		if visited[first] && visited[second] {
			tripletCount++
		}
		visited[num] = true
	}
	return tripletCount
}
