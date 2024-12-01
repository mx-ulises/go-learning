func checkIfExist(arr []int) bool {
	visited := map[int]bool{}
	for _, num := range arr {
		if visited[num<<1] {
			return true
		}
		if num&1 == 0 && visited[num>>1] {
			return true
		}
		visited[num] = true
	}
	return false
}
