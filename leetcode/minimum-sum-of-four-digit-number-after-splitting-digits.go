func minimumSum(num int) int {
	digits := []int{}
	for 0 < num {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)
	return (digits[0]+digits[1])*10 + (digits[2] + digits[3])
}
