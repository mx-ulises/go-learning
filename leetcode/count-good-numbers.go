func pow(p, x int) int {
	switch {
	case p == 0:
		return 1
	case p == 1:
		return x
	}
	result := pow(p/2, x)
	result *= result
	if p&1 == 1 {
		result *= x
	}
	return result % 1_000_000_007
}

func countGoodNumbers(n int64) int {
	m := int(n)
	even, primes := (m+1)>>1, m>>1
	return pow(even, 5) * pow(primes, 4) % 1_000_000_007
}
