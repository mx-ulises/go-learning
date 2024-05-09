func maximumHappinessSum(happiness []int, k int) int64 {
	var happinessSum, decreasedHappiness int64
	n := len(happiness)
	sort.Ints(happiness)
	for i := 0; i < k; i++ {
		j := n - i - 1
		happinessSum += max(int64(happiness[j])+decreasedHappiness, 0)
		decreasedHappiness--
	}
	return happinessSum
}
