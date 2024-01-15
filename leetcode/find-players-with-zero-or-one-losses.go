func findWinners(matches [][]int) [][]int {
	playerLosses := map[int]int{}
	for i := 0; i < len(matches); i++ {
		winner := matches[i][0]
		loser := matches[i][1]
		playerLosses[winner] += 0
		playerLosses[loser] += 1
	}
	answer := [][]int{{}, {}}
	for player, losses := range playerLosses {
		if losses == 0 {
			answer[0] = append(answer[0], player)
		} else if losses == 1 {
			answer[1] = append(answer[1], player)
		}
	}
	slices.Sort(answer[0])
	slices.Sort(answer[1])
	return answer
}
