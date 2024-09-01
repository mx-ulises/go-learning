func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	output := make([][]int, m)
	i, j := 0, 0
	for _, x := range original {
		if i == 0 {
			output[j] = make([]int, n)
		}
		output[j][i] = x
		i++
		if i == n {
			i, j = 0, j+1
		}
	}
	return output
}
