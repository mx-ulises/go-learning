func separateDigits(nums []int) []int {
	output, s := []int{}, []int{}
	for _, num := range nums {
		for 0 < num {
			s = append(s, num%10)
			num /= 10
		}
		n := len(s)
		for 0 < n {
			n--
			output = append(output, s[n])
			s = s[:n]
		}
	}
	return output
}
