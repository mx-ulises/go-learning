func firstUniqChar(s string) int {
	visited := map[rune]int{}
	for _, c := range s {
		visited[c]++
	}
	for i, c := range s {
		if visited[c] == 1 {
			return i
		}
	}
	return -1
}
