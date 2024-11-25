var pows []int
var graph [][]int

func getState(board [][]int) int {
	state := 0
	for _, row := range board {
		for _, num := range row {
			state = state*10 + num
		}
	}
	return state
}

func getZeroPosition(state int) int {
	zeroPosition := 0
	for state%10 != 0 {
		zeroPosition++
		state /= 10
	}
	return zeroPosition
}

func swap(state, zeroPosition, swapPosition int) int {
	d := (state / pows[swapPosition]) % 10
	return state - d*pows[swapPosition] + d*pows[zeroPosition]
}

func getNeighboors(state int) []int {
	zeroPosition := getZeroPosition(state)
	neightboors := []int{}
	for _, swapPosition := range graph[zeroPosition] {
		neightboors = append(neightboors, swap(state, zeroPosition, swapPosition))
	}
	return neightboors
}

func slidingPuzzle(board [][]int) int {
	pows = []int{1, 10, 100, 1000, 10000, 100000}
	graph = [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	s := [][]int{{getState(board), 0}}
	visited := map[int]bool{}
	for 0 < len(s) {
		state, moves := s[0][0], s[0][1]
		s = s[1:]
		if visited[state] == true {
			continue
		}
		visited[state] = true
		if state == 123450 {
			return moves
		}
		moves++
		neightboors := getNeighboors(state)
		for _, nb := range neightboors {
			s = append(s, []int{nb, moves})
		}
	}
	return -1
}
