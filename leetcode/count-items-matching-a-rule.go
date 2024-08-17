func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ruleIndexMap := map[string]int{"type": 0, "color": 1, "name": 2}
	ruleIndex := ruleIndexMap[ruleKey]
	matches := 0
	for _, item := range items {
		if item[ruleIndex] == ruleValue {
			matches++
		}
	}
	return matches
}
