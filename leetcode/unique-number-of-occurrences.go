func uniqueOccurrences(arr []int) bool {
	ocurrenceCount := map[int]int{}
	for i := 0; i < len(arr); i++ {
		ocurrenceCount[arr[i]]++
	}
	countOcurrenceItem := map[int]bool{}
	for _, v := range ocurrenceCount {
		if countOcurrenceItem[v] {
			return false
		}
		countOcurrenceItem[v] = true
	}
	return true
}
