func GetMinimalsAndMaximals(matrix *[][]int, n, m int) (*[]int, *[]int) {
	rowMinimals := make([]int, n)
	colMaximals := make([]int, m)
	for i := 0; i < n; i++ {
		rowMinimals[i] = 100001
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rowMinimals[i] = min(rowMinimals[i], (*matrix)[i][j])
			colMaximals[j] = max(colMaximals[j], (*matrix)[i][j])
		}
	}
	return &rowMinimals, &colMaximals
}

func GetLuckyNumbers(matrix *[][]int, rowMinimals, colMaximals *[]int, n, m int) *[]int {
	lucky := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			candidate := (*matrix)[i][j]
			if candidate == (*rowMinimals)[i] && candidate == (*colMaximals)[j] {
				lucky = append(lucky, candidate)
			}
		}
	}
	return &lucky
}

func luckyNumbers(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	rowMinimals, colMaximals := GetMinimalsAndMaximals(&matrix, n, m)
	return *(GetLuckyNumbers(&matrix, rowMinimals, colMaximals, n, m))
}
