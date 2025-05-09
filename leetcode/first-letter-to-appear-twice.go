func repeatedCharacter(s string) byte {
	visited := [26]bool{}
	for i := 0; i < len(s); i++ {
		j := int(s[i] - 'a')
		if visited[j] {
			return s[i]
		}
		visited[j] = true
	}
	return 'a'
}
