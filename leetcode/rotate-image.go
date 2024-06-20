func rotate(matrix [][]int) {
	start, end := 0, len(matrix)-1
	for start < end {
		n := end - start
		for i := 0; i < n; i++ {
			start_next := start + i
			end_next := end - i
			a := matrix[start][start_next]
			b := matrix[start_next][end]
			c := matrix[end][end_next]
			d := matrix[end_next][start]
			matrix[start][start_next] = d
			matrix[start_next][end] = a
			matrix[end][end_next] = b
			matrix[end_next][start] = c
		}
		start++
		end--
	}
}
