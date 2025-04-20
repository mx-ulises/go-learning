func numRabbits(answers []int) int {
	digitCount := map[int]int{}
	total := 0
	for _, digit := range answers {
		if digitCount[digit] == 0 {
			total += digit + 1
		}
		digitCount[digit]++
		if (digit + 1) == digitCount[digit] {
			digitCount[digit] = 0
		}
	}
	return total
}
