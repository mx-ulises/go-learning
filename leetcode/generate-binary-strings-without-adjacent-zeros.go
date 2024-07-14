func validStrings(n int) []string {
	prev := []string{"0", "1"}
	for i := 1; i < n; i++ {
		current := []string{}
		for _, s := range prev {
			if s[i-1] == '1' {
				current = append(current, s+"0")
			}
			current = append(current, s+"1")
		}
		prev = current
	}
	return prev
}
