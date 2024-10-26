func Test(word *string, pref *string) bool {
	n := len(*pref)
	if len(*word) < n {
		return false
	}
	for i := 0; i < n; i++ {
		if (*word)[i] != (*pref)[i] {
			return false
		}
	}
	return true
}

func prefixCount(words []string, pref string) int {
	count := 0
	for _, word := range words {
		if Test(&word, &pref) == true {
			count++
		}
	}
	return count
}
