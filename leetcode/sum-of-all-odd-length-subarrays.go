func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	totalSum := 0
	for size := 1; size <= n; size += 2 {
		currentSum := 0
		currentSize := 0
		for i := 0; i < n; i++ {
			currentSum += arr[i]
			currentSize++
			if size < currentSize {
				currentSum -= arr[i-size]
				currentSize--
			}
			if size == currentSize {
				totalSum += currentSum
			}
		}
	}
	return totalSum
}
