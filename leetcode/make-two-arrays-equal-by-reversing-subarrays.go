func canBeEqual(target []int, arr []int) bool {
	n := len(target)
	numberCount := map[int]int{}
	for i := 0; i < n; i++ {
		numberCount[target[i]]++
		numberCount[arr[i]]--
	}
	for _, count := range numberCount {
		if count != 0 {
			return false
		}
	}
	return true
}
