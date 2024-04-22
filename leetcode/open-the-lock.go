type State struct {
	Moves int
	Value string
}

func GetNextValue(value string, i int, CONNECTS *map[byte][]byte, j int) string {
	output := []byte{}
	for k := 0; k < 4; k++ {
		c := value[k]
		if k == i {
			c = (*CONNECTS)[c][j]
		}
		output = append(output, c)
	}
	return string(output)
}

func openLock(deadends []string, target string) int {
	CONNECTS := map[byte][]byte{
		'0': {'9', '1'},
		'1': {'0', '2'},
		'2': {'1', '3'},
		'3': {'2', '4'},
		'4': {'3', '5'},
		'5': {'4', '6'},
		'6': {'5', '7'},
		'7': {'6', '8'},
		'8': {'7', '9'},
		'9': {'8', '0'},
	}
	visited := map[string]bool{}
	for _, deadend := range deadends {
		visited[deadend] = true
	}
	queue := []State{State{0, "0000"}}
	for 0 < len(queue) {
		state := queue[0]
		queue = queue[1:]
		if state.Value == target {
			return state.Moves
		}
		if visited[state.Value] {
			continue
		}
		visited[state.Value] = true
		for i := 0; i < 4; i++ {
			queue = append(queue, State{state.Moves + 1, GetNextValue(state.Value, i, &CONNECTS, 0)})
			queue = append(queue, State{state.Moves + 1, GetNextValue(state.Value, i, &CONNECTS, 1)})
		}
	}
	return -1
}
