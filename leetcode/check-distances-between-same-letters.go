package leetcode

func checkDistances(s string, distance []int) bool {
	lastIndex := map[int]int{}
	for current := 0; current < len(s); current++ {
		c := int(s[current] - 'a')
		previous, ok := lastIndex[c]
		if ok && (current-previous-1) != distance[c] {
			return false
		}
		lastIndex[c] = current
	}
	return true
}
