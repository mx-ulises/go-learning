package leetcode

func computeKey(word string) string {
	key := make([]byte, len(word))
	offset := word[0] - 'a'
	for i := 0; i < len(word); i++ {
		key[i] = word[i] - offset
	}
	return string(key)
}

func oddString(words []string) string {
	keyCount := map[string]int{}
	keyRep := map[string]string{}
	i := 0
	// Examine the first three words
	for i < 3 {
		key := computeKey(words[i])
		keyCount[key]++
		keyRep[key] = words[i]
		i++
	}
	// Examine the rest until we get two different keys
	for i < len(words) && len(keyCount) == 1 {
		key := computeKey(words[i])
		keyCount[key]++
		keyRep[key] = words[i]
		i++
	}
	// Examine the keys we have and return the one that is unique
	for key, count := range keyCount {
		if count == 1 {
			return keyRep[key]
		}
	}
	return ""
}
