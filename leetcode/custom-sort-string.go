func customSortString(order string, s string) string {
	charMap := map[rune]int{}
	for _, c := range s {
		charMap[c]++
	}
	sorted := []rune{}
	for _, c := range order {
		for 0 < charMap[c] {
			sorted = append(sorted, c)
			charMap[c]--
		}
	}
	for c, _ := range charMap {
		for 0 < charMap[c] {
			sorted = append(sorted, c)
			charMap[c]--
		}
	}
	return string(sorted)
}
