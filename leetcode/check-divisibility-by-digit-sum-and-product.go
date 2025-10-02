package leetcode

func checkDivisibility(n int) bool {
	digitSum, digitProduct := 0, 1
	m, digit := n, -1
	for 0 < m {
		digit, m = m%10, m/10
		digitSum, digitProduct = digitSum+digit, digitProduct*digit
	}
	return n%(digitSum+digitProduct) == 0
}
