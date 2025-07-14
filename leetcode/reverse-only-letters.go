package leetcode

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func reverseOnlyLetters(s string) string {
	buffer, l, r := make([]byte, len(s)), 0, len(s)-1
	for l <= r {
		for l <= r && !isLetter(s[l]) {
			buffer[l] = s[l]
			l++
		}
		for l <= r && !isLetter(s[r]) {
			buffer[r] = s[r]
			r--
		}
		if l <= r {
			buffer[r], buffer[l] = s[l], s[r]
			l++
			r--
		}
	}
	return string(buffer)
}
