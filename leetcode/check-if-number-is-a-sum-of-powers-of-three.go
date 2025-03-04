func check(n, i int, powers, powerSums []int) bool {
	if n < 0 || i < 0 || powerSums[i] < n {
		return false
	}
	if powerSums[i] == n || n == 0 {
		return true
	}
	i--
	output := check(n-powers[i+1], i, powers, powerSums) || check(n, i, powers, powerSums)
	return output
}

func checkPowersOfThree(n int) bool {
	powers := []int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441, 1594323, 4782969}
	powerSums := []int{1, 4, 13, 40, 121, 364, 1093, 3280, 9841, 29524, 88573, 265720, 797161, 2391484, 7174453}
	return check(n, 14, powers, powerSums)
}
