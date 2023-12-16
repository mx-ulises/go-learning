func isAnagram(s string, t string) bool {
	n := len(s)
	if len(t) != n {
		return false
	}
	anagramMap := make(map[byte]int)
	for i := 0; i < n; i++ {
		anagramMap[s[i]] += 1
		anagramMap[t[i]] -= 1
	}
	for _, v := range anagramMap {
		if v != 0 {
			return false
		}
	}
	return true
}
