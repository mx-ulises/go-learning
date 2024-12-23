func replaceDigits(s string) string {
	n := len(s)
	output := []byte{}
	isInt := 0
	for i := 0; i < n; i++ {
		if isInt == 1 {
			output = append(output, s[i-1]+s[i]-'0')
		} else {
			output = append(output, s[i])
		}
		isInt ^= 1
	}
	return string(output)
}
