func reverseDegree(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		total += int('z'-s[i]+1) * (i + 1)
	}
	return total
}
