package leetcode

func getKey(word string) int {
	key := 0
	for i := 0; i < len(word); i++ {
		b := 1 << int(word[i]-'a')
		key = key | b
	}
	return key
}

func similarPairs(words []string) int {
	visitedKeys := map[int]int{}
	pairs := 0
	for _, word := range words {
		key := getKey(word)
		pairs += visitedKeys[key]
		visitedKeys[key]++
	}
	return pairs
}
