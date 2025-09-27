package leetcode

func getValidHours(a, b byte) int {
	if a == '?' {
		switch b {
		case '?':
			return 24
		case '0', '1', '2', '3':
			return 3
		default:
			return 2
		}
	}
	if b == '?' {
		switch a {
		case '2':
			return 4
		default:
			return 10
		}
	}
	return 1
}

func getValidMinutes(a, b byte) int {
	first, second := 1, 1
	if a == '?' {
		first = 6
	}
	if b == '?' {
		second = 10
	}
	return first * second
}

func countTime(time string) int {
	a, b, c, d := time[0], time[1], time[3], time[4]
	validHours := getValidHours(a, b)
	validMinutes := getValidMinutes(c, d)
	return validHours * validMinutes
}
