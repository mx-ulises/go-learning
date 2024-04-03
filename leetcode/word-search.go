func IsValidPosition(board *[][]byte, word *string, i, j, k int) bool {
	if i < 0 || len(*board) <= i {
		return false
	}
	if j < 0 || len((*board)[i]) <= j {
		return false
	}
	return (*board)[i][j] == (*word)[k]
}

func DFS(board *[][]byte, word *string, i, j, k int) bool {
	if k == len(*word) {
		return true
	}
	if IsValidPosition(board, word, i, j, k) {
		(*board)[i][j] = ' '
		if DFS(board, word, i+1, j, k+1) {
			return true
		}
		if DFS(board, word, i-1, j, k+1) {
			return true
		}
		if DFS(board, word, i, j+1, k+1) {
			return true
		}
		if DFS(board, word, i, j-1, k+1) {
			return true
		}
		(*board)[i][j] = (*word)[k]
	}
	return false

}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if DFS(&board, &word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
