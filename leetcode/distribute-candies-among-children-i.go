func distributeCandies(n int, limit int) int {
	ways := 0
	for i := 0; i <= min(n, limit); i++ {
		for j := 0; j <= min(n-i, limit); j++ {
			if (n - i - j) <= limit {
				ways++
			}
		}
	}
	return ways
}
