func oddCells(m int, n int, indices [][]int) int {
	rows, cols := make([]int, m), make([]int, n)
	for _, index := range indices {
		rows[index[0]] ^= 1
		cols[index[1]] ^= 1
	}
	count := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			count += rows[r] ^ cols[c]
		}
	}
	return count
}
