func doesValidArrayExist(derived []int) bool {
	totalXor := 0
	for _, d := range derived {
		totalXor ^= d
	}
	return totalXor == 0
}
