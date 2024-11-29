func duplicateNumbersXOR(nums []int) int {
	visited := map[int]bool{}
	dupes := 0
	for _, num := range nums {
		dupes ^= num
		if visited[num] == false {
			visited[num] = true
			dupes ^= num
		}
	}
	return dupes
}
