func minimumPushes(word string) int {
	charCount := map[rune]int{}
	for _, c := range word {
		charCount[c]++
	}
	counts := []int{}
	for _, count := range charCount {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	n, i, pushes := len(counts), 0, 0
	for 0 < n {
		n--
		pushes += (1 + i/8) * counts[n]
		i++
	}
	return pushes
}
