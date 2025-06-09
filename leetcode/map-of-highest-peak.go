func initOutput(n, m int) [][]int {
	output := make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, m)
		for j := 0; j < m; j++ {
			output[i][j] = -1
		}
	}
	return output
}

func initQueue(n, m int, isWater [][]int) [][]int {
	q := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isWater[i][j] == 1 {
				q = append(q, []int{i, j, 0})
			}
		}
	}
	return q
}

func isValid(i, j, n, m int, output [][]int) bool {
	return 0 <= i && i < n && 0 <= j && j < m && output[i][j] == -1
}

func highestPeak(isWater [][]int) [][]int {
	NEIGHBOORS := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	n, m := len(isWater), len(isWater[0])
	output := initOutput(n, m)
	q := initQueue(n, m, isWater)
	for 0 < len(q) {
		i, j, h := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		if !isValid(i, j, n, m, output) {
			continue
		}
		output[i][j] = h
		h++
		for _, neighbor := range NEIGHBOORS {
			q = append(q, []int{i + neighbor[0], j + neighbor[1], h})
		}
	}
	return output
}
