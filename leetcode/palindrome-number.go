func isPalindrome(x int) bool {
	aux := x
	y := 0
	for 0 < aux {
		y *= 10
		y += aux % 10
		aux /= 10
	}
	return x == y
}
