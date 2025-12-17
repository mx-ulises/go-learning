package leetcode

func maxRepeating(sequence string, word string) int {
	maxRepeatingSubstringLen := 0
	for i := 0; i < len(sequence); i++ {
		candidateRepeatingSubstringLen := getCandidateRepeatingSubstringLen(sequence, i, word)
		maxRepeatingSubstringLen = max(maxRepeatingSubstringLen, candidateRepeatingSubstringLen)
	}
	return maxRepeatingSubstringLen
}

func getCandidateRepeatingSubstringLen(sequence string, start int, word string) int {
	repeatingSubstringLen, j := 0, 0
	for i := start; i < len(sequence); i++ {
		if sequence[i] != word[j] {
			break
		}
		j, repeatingSubstringLen = checkAndUpdateRepetitions(j+1, word, repeatingSubstringLen)
	}
	return repeatingSubstringLen
}

func checkAndUpdateRepetitions(j int, word string, repeatingSubstringLen int) (int, int) {
	if j == len(word) {
		j = 0
		repeatingSubstringLen++
	}
	return j, repeatingSubstringLen
}
