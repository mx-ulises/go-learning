package leetcode

func getScore(player []int) int {
	score := 0
	for i := 0; i < len(player); i++ {
		if 0 < i && player[i-1] == 10 {
			score += 2 * player[i]
		} else if 1 < i && player[i-2] == 10 {
			score += 2 * player[i]
		} else {
			score += player[i]
		}
	}
	return score
}

func isWinner(player1 []int, player2 []int) int {
	score1, score2 := getScore(player1), getScore(player2)
	if score1 < score2 {
		return 2
	} else if score2 < score1 {
		return 1
	}
	return 0
}
