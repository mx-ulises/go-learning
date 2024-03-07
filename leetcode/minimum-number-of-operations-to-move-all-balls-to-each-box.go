func minOperations(boxes string) []int {
	boxMap := map[byte]int{'0': 0, '1': 1}
	n := len(boxes)
	ballLeft, movesLeft := 0, 0
	ballRight, movesRight := 0, 0
	output := make([]int, n)
	for i := 1; i < n; i++ {
		ballLeft += boxMap[boxes[i-1]]
		movesLeft += ballLeft
		output[i] += movesLeft
		j := n - i
		ballRight += boxMap[boxes[j]]
		movesRight += ballRight
		output[j-1] += movesRight
	}
	return output
}
