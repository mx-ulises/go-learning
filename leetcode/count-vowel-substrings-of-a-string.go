package leetcode

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func countVowelSubstrings(word string) int {

	count := 0
	for i := 0; i < len(word); i++ {
		diffVowels := 0
		vowelArray := map[byte]bool{}
		for j := i; j < len(word); j++ {
			c := word[j]
			if !isVowel(c) {
				break
			}
			if vowelArray[c] == false {
				vowelArray[c] = true
				diffVowels++
			}
			if diffVowels == 5 {
				count++
			}
		}
	}
	return count
}
