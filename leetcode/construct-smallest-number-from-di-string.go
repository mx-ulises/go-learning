func buildCandidate(current []rune, used map[rune]bool, pattern string, i int) string {
	if i == len(pattern) {
		return string(current)
	}
	p, d := current[i], pattern[i]
	i++
	for x := '1'; x <= '9'; x++ {
		if used[x] || (d == 'I' && x < p) || (d == 'D' && p < x) {
			continue
		}
		used[x] = true
		current = append(current, x)
		candidate := buildCandidate(current, used, pattern, i)
		if candidate != "" {
			return candidate
		}
		current = current[:i]
		used[x] = false
	}
	return ""
}

func smallestNumber(pattern string) string {
	used := make(map[rune]bool)
	for x := '1'; x <= '9'; x++ {
		current := []rune{x}
		used[x] = true
		candidate := buildCandidate(current, used, pattern, 0)
		used[x] = false
		if candidate != "" {
			return candidate
		}
	}
	return ""
}
