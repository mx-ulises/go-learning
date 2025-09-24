package leetcode

import "sort"

func getKey(word string) string {
	key := []byte(word)
	sort.Slice(key, func(i, j int) bool {
		return key[i] < key[j]
	})
	return string(key)
}

func removeAnagrams(words []string) []string {
	prev := getKey(words[0])
	output := []string{words[0]}
	for _, word := range words {
		key := getKey(word)
		if prev != key {
			output = append(output, word)
			prev = key
		}
	}
	return output
}
