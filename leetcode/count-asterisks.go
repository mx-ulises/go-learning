func countAsterisks(s string) int {
	consider := 1
	count := 0
	for _, c := range s {
		if c == '|' {
			consider ^= 1
		}
		if c == '*' && consider == 1 {
			count += 1
		}
	}
	return count
}
