func judgeCircle(moves string) bool {
	if len(moves)&1 == 1 {
		return false
	}
	position := 0
	commands := map[rune]int{'U': 1, 'D': -1, 'L': 2, 'R': -2}
	for _, c := range moves {
		position += commands[c]
	}
	return position == 0
}
