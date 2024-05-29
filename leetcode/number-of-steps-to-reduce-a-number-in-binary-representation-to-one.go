func numSteps(s string) int {
	if s == "10" {
		return 1
	} else if s == "1" {
		return 0
	}
	n := len(s)
	sArray := make([]rune, n)
	for i, c := range s {
		sArray[i] = c
	}
	i := n - 1
	moves := 0
	for 0 <= i {
		aux := sArray[i]
		if sArray[i] == '1' {
			moves++
		}
		for 0 <= i && sArray[i] == aux {
			moves++
			i--
		}
		if 0 <= i {
			sArray[i] = '1'
		}
	}
	return moves
}
