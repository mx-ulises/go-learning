func MyPow(n int, x float64) float64 {
	if n == 0 {
		return 1.0
	}
	sqrt := MyPow(n/2, x)
	extra := 1.0
	if (n & 1) == 1 {
		extra = x
	}
	return sqrt * sqrt * extra
}

func myPow(x float64, n int) float64 {
	output := MyPow(n, x)
	if n < 0 {
		output = 1 / output
	}
	return output
}
