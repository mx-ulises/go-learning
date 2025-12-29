package leetcode

func isFascinating(n int) bool {
	digits := [10]int{}
	for i := 1; i <= 3; i++ {
		digits = addDigits(digits, n*i)
	}
	return haveAllDigits(digits)
}

func addDigits(digits [10]int, n int) [10]int {
	for 0 < n {
		digits[n%10]++
		n /= 10
	}
	return digits
}

func haveAllDigits(digits [10]int) bool {
	for i := 1; i < 10; i++ {
		if digits[i] != 1 {
			return false
		}
	}
	return digits[0] == 0
}
