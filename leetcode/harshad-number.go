func sumOfTheDigitsOfHarshadNumber(x int) int {
	digitSum, aux := 0, x
	for 0 < aux {
		digitSum += aux % 10
		aux /= 10
	}
	if (x % digitSum) == 0 {
		return digitSum
	}
	return -1
}
