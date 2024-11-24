func maxMatrixSum(matrix [][]int) int64 {
	negativeCount := 0
	matrixSum := 0
	toBeRemove := 1000000
	for _, row := range matrix {
		for _, num := range row {
			if num < 0 {
				negativeCount ^= 1
				num = -num
			}
			matrixSum += num
			toBeRemove = min(toBeRemove, num)
		}
	}
	if negativeCount == 1 {
		matrixSum -= 2 * toBeRemove
	}
	return int64(matrixSum)
}
