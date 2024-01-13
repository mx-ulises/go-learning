func combine(n int, k int) [][]int {
	output := [][]int{{}}
	for i := 0; i < k; i++ {
		currentCombinations := len(output)
		for j := 0; j < currentCombinations; j++ {
			combination := output[0]
			output = output[1:]
			start := 0
			combinationLen := len(combination)
			if 0 < combinationLen {
				start = combination[combinationLen-1]
			}
			for start < n {
				start += 1
				newCombination := make([]int, combinationLen+1)
				copy(newCombination, combination)
				newCombination[combinationLen] = start
				output = append(output, newCombination)
			}
		}
	}
	return output
}
