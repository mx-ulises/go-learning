func IsPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func Reverse(s string) string {
	n := len(s)
	output := make([]byte, n)
	for i := 0; i < n; i++ {
		output[i] = s[n-i-1]
	}
	return string(output)
}

func shortestPalindrome(s string) string {
	palindromeSize, n := 0, len(s)
	for i := n; 0 < n; i-- {
		if IsPalindrome(s[:i]) {
			palindromeSize = i
			break
		}
	}
	prefix := Reverse(s[palindromeSize:])
	return prefix + s
}
