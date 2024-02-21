func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for 0 < n {
		digit := n % 10
		product *= digit
		sum += digit
		n /= 10
	}
	return product - sum
}
