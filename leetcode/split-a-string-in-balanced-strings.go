func balancedStringSplit(s string) int {
	balance, groups := 0, 0
	for _, c := range s {
		switch c {
		case 'R':
			balance++
		case 'L':
			balance--
		}
		if balance == 0 {
			groups++
		}
	}
	return groups
}
