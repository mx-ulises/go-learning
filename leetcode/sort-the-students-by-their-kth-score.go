func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[j][k] < score[i][k]
	})
	return score
}
