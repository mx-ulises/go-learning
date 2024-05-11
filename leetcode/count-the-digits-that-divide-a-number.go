func countDigits(num int) int {
	count, current := 0, num
	for 0 < current {
		if num%(current%10) == 0 {
			count++
		}
		current /= 10
	}
	return count
}
