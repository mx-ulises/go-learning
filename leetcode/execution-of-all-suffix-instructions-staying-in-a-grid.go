func getMoves(n, x, y int, s string, i int, moves map[byte][]int) int {
	m := 0
	for j := i; j < len(s); j++ {
		d := s[j]
		x += moves[d][0]
		y += moves[d][1]
		if x < 0 || x == n || y < 0 || y == n {
			break
		}
		m++
	}
	return m
}

func executeInstructions(n int, startPos []int, s string) []int {
	MOVES := map[byte][]int{'R': {0, 1}, 'L': {0, -1}, 'U': {-1, 0}, 'D': {1, 0}}
	output := []int{}
	for i := 0; i < len(s); i++ {
		output = append(output, getMoves(n, startPos[0], startPos[1], s, i, MOVES))
	}
	return output
}
