func finalPositionOfSnake(n int, commands []string) int {
	COMMANDS := map[string]int{"UP": -n, "RIGHT": 1, "DOWN": n, "LEFT": -1}
	position := 0
	for _, c := range commands {
		position += COMMANDS[c]
	}
	return position
}
