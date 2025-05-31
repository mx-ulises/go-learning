func reverseRows(board [][]int) [][]int {
	n := len(board)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		aux := board[i]
		board[i] = board[j]
		board[j] = aux
	}
	return board
}

func reverseCols(board [][]int) [][]int {
	n := len(board)
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			continue
		}
		for j := 0; j < n/2; j++ {
			k := n - j - 1
			aux := board[i][j]
			board[i][j] = board[i][k]
			board[i][k] = aux
		}
	}
	return board
}

func normalize(board [][]int) [][]int {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if 0 < board[i][j] {
				board[i][j]--
			} else {
				board[i][j] = i*n + j
			}
		}
	}
	return board
}

func getPosition(position int, board [][]int) int {
	n := len(board)
	x, y := position/n, position%n
	return board[x][y]
}

func snakesAndLadders(board [][]int) int {
	board = reverseRows(board)
	board = reverseCols(board)
	board = normalize(board)
	end := (len(board) * len(board)) - 1
	visited := make([]bool, end+1)
	q := [][]int{{0, 0}}
	for 0 < len(q) {
		pos, rolls := q[0][0], q[0][1]
		q = q[1:]
		if visited[pos] {
			continue
		}
		visited[pos] = true
		if pos == end {
			return rolls
		}
		rolls++
		for succ := pos + 1; succ <= min(pos+6, end); succ++ {
			q = append(q, []int{getPosition(succ, board), rolls})
		}
	}
	return -1
}
