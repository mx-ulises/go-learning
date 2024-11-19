var isVowel map[rune]bool = map[rune]bool{
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
}

func sortVowels(s string) string {
	vowels := make([]rune, 0)
	newS := make([]rune, len(s))
	// Collect vowels and string array
	for i, c := range s {
		if isVowel[c] {
			vowels = append(vowels, c)
			newS[i] = '-'
		} else {
			newS[i] = c
		}
	}
	// Sort vowels
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})
	// Replace in new string
	for i, c := range newS {
		if c == '-' {
			newS[i] = vowels[0]
			vowels = vowels[1:]
		}
	}
	return string(newS)
}
