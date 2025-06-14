func rowAndMaximumOnes(mat [][]int) []int {
	maxRow, maxSum := 0, 0
	for i, row := range mat {
		sum := 0
		for _, x := range row {
			sum += x
		}
		if maxSum < sum {
			maxSum, maxRow = sum, i
		}
	}
	return []int{maxRow, maxSum}
}
