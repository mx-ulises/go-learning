func longestMonotonicSubarray(nums []int) int {
	maximal, currentInc, currentDec := 1, 0, 0
	prevNum := nums[0]
	for _, num := range nums {
		if prevNum < num {
			currentInc++
			currentDec = 1
		} else if num < prevNum {
			currentDec++
			currentInc = 1
		} else {
			currentDec = 1
			currentInc = 1
		}
		maximal = max(maximal, currentInc, currentDec)
		prevNum = num
	}
	return maximal
}
