package leetcode

import "sort"

func inRange(c, start, end rune) bool {
	return start <= c && c <= end
}

func isDigit(c rune) bool {
	return inRange(c, '0', '9')
}

func isChar(c rune) bool {
	return inRange(c, 'a', 'z') || inRange(c, 'A', 'Z')
}

func isValidCode(code string) bool {
	for _, c := range code {
		if isDigit(c) == false && isChar(c) == false && c != '_' {
			return false
		}
	}
	return len(code) != 0
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	validBussinessLines := map[string]bool{
		"electronics": true,
		"grocery":     true,
		"pharmacy":    true,
		"restaurant":  true,
	}
	codesPerBussinessLine := map[string][]string{}
	n := len(code)
	for i := 0; i < n; i++ {
		bl := businessLine[i]
		if isValidCode(code[i]) && validBussinessLines[bl] && isActive[i] {
			codesPerBussinessLine[bl] = append(codesPerBussinessLine[bl], code[i])
		}
	}
	validCoupons := []string{}
	for _, bl := range []string{"electronics", "grocery", "pharmacy", "restaurant"} {
		codes := codesPerBussinessLine[bl]
		sort.Strings(codes)
		validCoupons = append(validCoupons, codes...)
	}
	return validCoupons
}
