func bagOfTokensScore(tokens []int, power int) int {
	score, maxScore := 0, 0
	left, right := 0, len(tokens)-1
	sort.Ints(tokens)
	for left <= right {
		if tokens[left] <= power {
			power -= tokens[left]
			score++
			left++
			if maxScore < score {
				maxScore = score
			}
		} else if 1 <= score {
			power += tokens[right]
			score--
			right--
		} else {
			break
		}
	}
	return maxScore
}
