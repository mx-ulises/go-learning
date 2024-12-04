func increment(a byte) byte {
	if a == 'z' {
		return 'a'
	}
	return a + 1
}

func canMakeSubsequence(str1 string, str2 string) bool {
	j := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[j] || increment(str1[i]) == str2[j] {
			j++
		}
		if j == len(str2) {
			return true
		}
	}
	return false
}
