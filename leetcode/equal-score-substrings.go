package leetcode

func scoreBalance(s string) bool {
	scoreLeft, scoreRight := 0, getStringScore(s)
	for i := 0; i < len(s); i++ {
		if scoreLeft == scoreRight {
			return true
		}
		scoreLeft, scoreRight = updateScores(scoreLeft, scoreRight, s[i])
	}
	return scoreLeft == scoreRight
}

func getStringScore(s string) int {
	score := 0
	for i := 0; i < len(s); i++ {
		score += getPoints(s[i])
	}
	return score
}

func getPoints(c byte) int {
	return int(c-'a') + 1
}

func updateScores(scoreLeft int, scoreRight int, c byte) (int, int) {
	points := getPoints(c)
	return scoreLeft + points, scoreRight - points
}
