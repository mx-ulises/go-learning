package leetcode

func thirdDigit(numCount map[int]int, num int, output []int) []int {
	for i := 0; i < 5; i++ {
		d := i * 2
		if numCount[d] == 0 {
			continue
		}
		output = append(output, num+d)
	}
	return output
}

func secondDigit(numCount map[int]int, num int, output []int) []int {
	for d := 0; d < 10; d++ {
		if numCount[d] == 0 {
			continue
		}
		numCount[d]--
		output = thirdDigit(numCount, num+d*10, output)
		numCount[d]++
	}
	return output
}

func firstDigit(numCount map[int]int) []int {
	output := []int{}
	for d := 1; d < 10; d++ {
		if numCount[d] == 0 {
			continue
		}
		numCount[d]--
		output = secondDigit(numCount, d*100, output)
		numCount[d]++
	}
	return output
}

func findEvenNumbers(digits []int) []int {
	numCount := map[int]int{}
	for _, d := range digits {
		numCount[d]++
	}
	return firstDigit(numCount)
}
