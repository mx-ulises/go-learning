func initializeOutput(rows, cols int) [][]byte {
	output := make([][]byte, rows)
	for i, _ := range output {
		output[i] = make([]byte, cols)
		for j := 0; j < cols; j++ {
			output[i][j] = '.'
		}
	}
	return output
}

func rotateTheBox(boxGrid [][]byte) [][]byte {
	rows, cols := len(boxGrid), len(boxGrid[0])
	output := initializeOutput(cols, rows)
	for i := 0; i < rows; i++ {
		output_row := cols - 1
		output_col := rows - i - 1
		for j := cols - 1; 0 <= j; j-- {
			if boxGrid[i][j] == '#' {
				output[output_row][output_col] = '#'
				output_row--
			}
			if boxGrid[i][j] == '*' {
				output[j][output_col] = '*'
				output_row = j - 1
			}
		}
	}
	return output
}
