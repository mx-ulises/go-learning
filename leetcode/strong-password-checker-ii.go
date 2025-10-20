package leetcode

func inRange(c, start, end byte) bool {
	return start <= c && c <= end
}

func getSpecial(s string) map[byte]bool {
	special := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		special[s[i]] = true
	}
	return special
}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	specialChars := getSpecial("!@#$%^&*()-+")
	lower, upper, digit, special := false, false, false, false
	for i := 0; i < len(password); i++ {
		c := password[i]
		if 0 < i && c == password[i-1] {
			return false
		}
		lower = lower || inRange(c, 'a', 'z')
		upper = upper || inRange(c, 'A', 'Z')
		digit = digit || inRange(c, '0', '9')
		special = special || specialChars[c]
	}
	return lower && upper && digit && special
}
