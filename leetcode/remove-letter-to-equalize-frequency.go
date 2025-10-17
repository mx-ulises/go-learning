package leetcode

func equalFrequency(word string) bool {
	letterFrequency := map[byte]int{}
	for i := 0; i < len(word); i++ {
		letterFrequency[word[i]]++
	}
	frequencyCount, minFrequency := map[int]int{}, 10000
	for _, frequency := range letterFrequency {
		frequencyCount[frequency]++
		minFrequency = min(minFrequency, frequency)
	}
	if 2 < len(frequencyCount) {
		return false
	}
	if len(frequencyCount) == 2 {
		return (minFrequency == 1 && frequencyCount[1] == 1) || frequencyCount[minFrequency+1] == 1
	}
	return minFrequency == 1 || frequencyCount[minFrequency] == 1

}
