func makeFancyString(s string) string {
	newS := []rune{}
	for _, c := range s {
		n := len(newS)
		if n < 2 {
			newS = append(newS, c)
		} else if newS[n-1] != c || newS[n-2] != c {
			newS = append(newS, c)
		}
	}
	return string(newS)
}
