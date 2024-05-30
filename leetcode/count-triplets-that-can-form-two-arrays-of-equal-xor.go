func countTriplets(arr []int) int {
	n := len(arr)
	count := 0
	for i := 0; i < n; i++ {
		a := arr[i]
		for j := i + 1; j < n; j++ {
			b := 0
			for k := j; k < n; k++ {
				b ^= arr[k]
				if a == b {
					count++
				}
			}
			a ^= arr[j]
		}
	}
	return count
}
