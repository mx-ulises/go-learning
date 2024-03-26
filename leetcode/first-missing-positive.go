func firstMissingPositive(nums []int) int {
	n := 100000
	if len(nums) < n {
		n = len(nums)
	}
	numChecker := make([]bool, n+1)
	for _, num := range nums {
		if 0 < num && num <= n {
			numChecker[num-1] = true
		}
	}
	for i, present := range numChecker {
		if present == false {
			return i + 1
		}
	}
	return -1
}
