package leetcode

func getLeastFrequentDigit(n int) int {
	digits := [10]int{}
	for 0 < n {
		digits[n%10]++
		n /= 10
	}
	output, minFrequency := -1, 100
	for digit, frequency := range digits {
		if frequency != 0 && frequency < minFrequency {
			minFrequency, output = frequency, digit
		}
	}
	return output
}
