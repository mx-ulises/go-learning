func GetMaxCandies(candies *[]int, n int) int {
	maximal := 0
	for i := 0; i < n; i++ {
		maximal = max(maximal, (*candies)[i])
	}
	return maximal
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := GetMaxCandies(&candies, n)
	output := make([]bool, n)
	for i := 0; i < n; i++ {
		if maxCandies <= (candies[i] + extraCandies) {
			output[i] = true
		}
	}
	return output
}
