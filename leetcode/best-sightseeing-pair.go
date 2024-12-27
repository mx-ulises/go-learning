func maxScoreSightseeingPair(values []int) int {
	score, maximal := 0, values[0]
	for i := 1; i < len(values); i++ {
		score = max(score, maximal+values[i]-i)
		maximal = max(maximal, values[i]+i)
	}
	return score
}
