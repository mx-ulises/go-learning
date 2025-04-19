func valid(r, c, rows, cols int) bool {
	return 0 <= r && r < rows && 0 <= c && c < cols
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	total, x := rows*cols, 1
	rMoveCount, cMoveCount := 1, 1
	r, c := rStart, cStart
	rToChange, cToChange := r+1, c+1
	directions, d := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}, 0
	output := make([][]int, total)
	for x <= total {
		if valid(r, c, rows, cols) {
			output[x-1] = []int{r, c}
			x++
		}
		r += directions[d][0]
		c += directions[d][1]
		if r == rToChange || c == cToChange {
			if r == rToChange {
				rMoveCount++
				rToChange = r - (rMoveCount * directions[d][0])
			}
			if c == cToChange {
				cMoveCount++
				cToChange = c - (cMoveCount * directions[d][1])
			}
			d = (d + 1) % 4
		}
	}
	return output
}
