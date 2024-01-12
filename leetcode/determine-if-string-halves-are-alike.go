func halvesAreAlike(s string) bool {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	mid := len(s) / 2
	aVowelCount := 0
	bVowelCount := 0
	for i := 0; i < mid; i++ {
		if vowels[s[i]] {
			aVowelCount++
		}
		if vowels[s[i+mid]] {
			bVowelCount++
		}
	}
	return aVowelCount == bVowelCount
}
