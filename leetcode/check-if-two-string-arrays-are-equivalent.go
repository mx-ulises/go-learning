func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i1, j1, n1 := 0, 0, len(word1)
	i2, j2, n2 := 0, 0, len(word2)
	for i1 < n1 && i2 < n2 {
		if word1[i1][j1] == word2[i2][j2] {
			j1++
			j2++
			if j1 == len(word1[i1]) {
				j1 = 0
				i1++
			}
			if j2 == len(word2[i2]) {
				j2 = 0
				i2++
			}
		} else {
			return false
		}
	}
	return i1 == n1 && i2 == n2
}
