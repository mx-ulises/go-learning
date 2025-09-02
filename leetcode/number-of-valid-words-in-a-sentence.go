package leetcode

import "fmt"

func inRange(c, start, end byte) bool {
	return start <= c && c <= end
}

func isValid(sentence string, start, end int) bool {
	haveHyphen := false
	if start == end {
		return false
	}
	if sentence[start] == '-' || sentence[end-1] == '-' {
		return false
	}
	for i := start; i < end; i++ {
		if inRange(sentence[i], '0', '9') {
			return false
		}
		if sentence[i] == ',' || sentence[i] == '.' || sentence[i] == '!' {
			if i != end-1 {
				return false
			}
		}
		if sentence[i] == '-' {
			if haveHyphen || !inRange(sentence[i-1], 'a', 'z') || !inRange(sentence[i+1], 'a', 'z') {
				return false
			}
			haveHyphen = true
		}
	}
	return true
}

func countValidWords(sentence string) int {
	start, end, n := 0, 0, len(sentence)
	validWordCount := 0
	for end < n {
		if sentence[end] == ' ' {
			if isValid(sentence, start, end) {
				validWordCount++
				fmt.Println(sentence[start : end+1])
			}
			start = end + 1
		}
		end++
	}
	if isValid(sentence, start, end) {
		validWordCount++
	}
	return validWordCount
}
