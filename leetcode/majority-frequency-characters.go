package leetcode

func majorityFrequencyGroup(s string) string {
	frequencies := [26]int{}
	for i := 0; i < len(s); i++ {
		frequencies[int(s[i]-'a')]++
	}
	frequencyRepetitions := map[int]int{}
	maximalRepetitions, targetFrequency := -1, -1
	for _, frequency := range frequencies {
		if frequency == 0 {
			continue
		}
		frequencyRepetitions[frequency]++
		if maximalRepetitions < frequencyRepetitions[frequency] {
			maximalRepetitions = frequencyRepetitions[frequency]
			targetFrequency = frequency
		}
		if maximalRepetitions == frequencyRepetitions[frequency] {
			targetFrequency = max(targetFrequency, frequency)
		}
	}
	output := []byte{}
	for i, frequency := range frequencies {
		if targetFrequency == frequency {
			output = append(output, byte(i)+'a')
		}
	}
	return string(output)
}
