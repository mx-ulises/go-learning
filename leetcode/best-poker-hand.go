package leetcode

func bestHand(ranks []int, suits []byte) string {
	maxRankCount, rankCount := 0, [14]int{}
	maxSuitCount, suitCount := 0, [4]int{}
	for i := 0; i < 5; i++ {
		rank, suit := ranks[i], int(suits[i]-'a')
		rankCount[rank]++
		maxRankCount = max(maxRankCount, rankCount[rank])
		suitCount[suit]++
		maxSuitCount = max(maxSuitCount, suitCount[suit])
	}
	if 5 == maxSuitCount {
		return "Flush"
	}
	if 3 <= maxRankCount {
		return "Three of a Kind"
	}
	if 2 == maxRankCount {
		return "Pair"
	}
	return "High Card"
}
