func areAlmostEqual(s1 string, s2 string) bool {
	n := len(s1)
	diffIndexes := []int{}
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			diffIndexes = append(diffIndexes, i)
		}
	}
	if len(diffIndexes) == 2 {
		i, j := diffIndexes[0], diffIndexes[1]
		return s1[i] == s2[j] && s1[j] == s2[i]
	}
	return len(diffIndexes) == 0
}
