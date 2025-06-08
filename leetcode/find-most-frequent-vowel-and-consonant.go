func maxFreqSum(s string) int {
	vowels := map[int]bool{0: true, 4: true, 8: true, 14: true, 20: true}
	maxVowel, maxConsonant := 0, 1
	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		c := int(s[i]) - 97
		chars[c]++
		if vowels[c] && chars[maxVowel] < chars[c] {
			maxVowel = c
		}
		if !vowels[c] && chars[maxConsonant] < chars[c] {
			maxConsonant = c
		}
	}
	return chars[maxVowel] + chars[maxConsonant]
}
