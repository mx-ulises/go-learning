package leetcode

func maximizeExpressionOfThree(nums []int) int {
	a, b := getGreatestTwo(nums)
	c := getMinimum(nums)
	return a + b - c
}

func getGreatestTwo(nums []int) (int, int) {
	a, b := -100, -100
	for _, candidate := range nums {
		a, b = checkAndUpdate(a, b, candidate)
	}
	return a, b
}

func checkAndUpdate(a int, b int, candidate int) (int, int) {
	if a <= candidate {
		a, b = candidate, a
	} else {
		b = max(b, candidate)
	}
	return a, b
}

func getMinimum(nums []int) int {
	minimum := nums[0]
	for _, candidate := range nums {
		minimum = min(minimum, candidate)
	}
	return minimum
}
