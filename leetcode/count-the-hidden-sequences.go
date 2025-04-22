func numberOfArrays(differences []int, lower int, upper int) int {
	minimal, maximal, current := 0, 0, 0
	for _, difference := range differences {
		current += difference
		minimal = min(minimal, current)
		maximal = max(maximal, current)
	}
	return max(0, upper-lower-maximal+minimal+1)

}
