func removeTrailingZeros(num string) string {
	if num == "0" {
		return "0"
	}
	i := len(num) - 1
	for 0 <= i && num[i] == '0' {
		i--
	}
	return num[:i+1]
}
