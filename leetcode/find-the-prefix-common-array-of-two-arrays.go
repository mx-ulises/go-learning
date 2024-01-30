func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	sawA := map[int]int{}
	sawB := map[int]int{}
	C := make([]int, n)
	c := 0
	for i := 0; i < n; i++ {
		a := A[i]
		b := B[i]
		sawA[a]++
		sawB[b]++
		if 0 < sawA[b] {
			sawA[b]--
			c++
		}
		if 0 < sawB[a] {
			sawB[a]--
			c++
		}
		if a == b {
			c--
		}
		C[i] = c
	}
	return C
}
