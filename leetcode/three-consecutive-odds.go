func threeConsecutiveOdds(arr []int) bool {
	n := len(arr)
	oddCount := 0
	for i := 0; i < n; i++ {
		if arr[i]&1 == 1 {
			oddCount++
			if oddCount == 3 {
				return true
			}
		} else {
			oddCount = 0
		}
	}
	return false
}
