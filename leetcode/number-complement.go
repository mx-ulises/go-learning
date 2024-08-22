func findComplement(num int) int {
	mask := 1
	for mask <= num {
		mask <<= 1
	}
	mask--
	return mask ^ num
}
