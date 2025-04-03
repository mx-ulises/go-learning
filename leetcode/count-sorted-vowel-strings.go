func countVowelStrings(n int) int {
	memory := []int{1, 0, 0, 0, 0}
	total := 1
	for i := 0; i < n; i++ {
		for j := 1; j < 5; j++ {
			memory[j] += memory[j-1]
			total += memory[j-1]
		}
	}
	return total
}
