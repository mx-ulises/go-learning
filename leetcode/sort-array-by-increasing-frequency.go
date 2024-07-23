func GetNumToFrequencies(nums *[]int) *map[int]int {
	numFrequencies := map[int]int{}
	for _, num := range *nums {
		numFrequencies[num]++
	}
	return &numFrequencies
}

func GetFrequenciesToNum(numFrequencies *map[int]int) *map[int][]int {
	frequenciesNum := map[int][]int{}
	for num, freq := range *numFrequencies {
		frequenciesNum[freq] = append(frequenciesNum[freq], num)
	}
	for freq, _ := range frequenciesNum {
		sort.Ints(frequenciesNum[freq])
	}
	return &frequenciesNum
}

func GetFrequencies(frequenciesNum *map[int][]int) *[]int {
	frequencies := []int{}
	for freq, _ := range *frequenciesNum {
		frequencies = append(frequencies, freq)
	}
	sort.Ints(frequencies)
	return &frequencies
}

func frequencySort(nums []int) []int {
	numFrequencies := GetNumToFrequencies(&nums)
	frequenciesNum := GetFrequenciesToNum(numFrequencies)
	frequencies := GetFrequencies(frequenciesNum)
	output := []int{}
	for _, freq := range *frequencies {
		n := len((*frequenciesNum)[freq]) - 1
		for 0 <= n {
			for i := 0; i < freq; i++ {
				output = append(output, (*frequenciesNum)[freq][n])
			}
			n--
		}
	}
	return output
}
