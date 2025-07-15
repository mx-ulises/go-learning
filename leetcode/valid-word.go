package leetcode

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isVowel(c byte) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U',
		'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isConsonant(c byte) bool {
	return isLetter(c) && !isVowel(c)
}

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	consonant, vowel := false, false

	for i := 0; i < len(word); i++ {
		c := word[i]
		if isConsonant(c) {
			consonant = true
		} else if isVowel(c) {
			vowel = true
		} else if !isDigit(c) {
			return false
		}
	}
	return consonant && vowel
}
