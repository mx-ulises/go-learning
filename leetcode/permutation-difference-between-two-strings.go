func abs(x int) int {
	if 0 < x {
		return x
	}
	return -x
}

func findPermutationDifference(s string, t string) int {
	n := len(s)
	charDiff := map[byte]int{}
	for i := 0; i < n; i++ {
		charDiff[s[i]] += i
		charDiff[t[i]] -= i
	}
	totalDiff := 0
	for _, diff := range charDiff {
		totalDiff += abs(diff)
	}
	return totalDiff
}
