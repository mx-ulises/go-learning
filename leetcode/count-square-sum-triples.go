package leetcode

func countTriples(n int) int {
	squares := map[int]bool{}
	for i := 1; i <= n; i++ {
		squares[i*i] = true
	}
	triplets := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if squares[i*i+j*j] {
				triplets += 2
			}
		}
	}
	return triplets
}
