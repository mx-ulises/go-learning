func maxFrequencyElements(nums []int) int {
	frequencyMap := map[int]int{}
	maxFrequency := 0
	countFrequency := 0
	for _, x := range nums {
		frequencyMap[x]++
		if maxFrequency < frequencyMap[x] {
			maxFrequency = frequencyMap[x]
			countFrequency = 0
		}
		if maxFrequency == frequencyMap[x] {
			countFrequency++
		}
	}
	return maxFrequency * countFrequency
}
