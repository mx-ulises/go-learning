func findArray(pref []int) []int {
	arr := []int{}
	prev := 0
	for _, current := range pref {
		arr = append(arr, prev^current)
		prev = current
	}
	return arr
}
