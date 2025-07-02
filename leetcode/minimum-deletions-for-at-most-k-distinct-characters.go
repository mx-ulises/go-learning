package leetcode

import "sort"

func getCharCount(s string) map[rune]int {
	charCount := map[rune]int{}
	for _, c := range s {
		charCount[c]++
	}
	return charCount
}

func getCounts(charCount map[rune]int) []int {
	counts, i := make([]int, len(charCount)), 0
	for _, count := range charCount {
		counts[i] = count
		i++
	}
	return counts
}

func minDeletion(s string, k int) int {
	charCount := getCharCount(s)
	counts := getCounts(charCount)
	sort.Ints(counts)
	deletions := 0
	for i := 0; i < len(counts)-k; i++ {
		deletions += counts[i]
	}
	return deletions
}
