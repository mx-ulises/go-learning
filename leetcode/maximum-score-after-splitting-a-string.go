func maxScore(s string) int {
	leftScore := 0
	if s[0] == '0' {
		leftScore++
	}
	rightScore := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			rightScore++
		}
	}
	maximalScore := leftScore + rightScore
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			leftScore++
		} else {
			rightScore--
		}
		maximalScore = max(maximalScore, leftScore+rightScore)
	}
	return maximalScore
}
