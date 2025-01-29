func findGCD(nums []int) int {
	maximal, minimal := 0, 1001
	for _, num := range nums {
		maximal = max(maximal, num)
		minimal = min(minimal, num)
	}
	for i := minimal; 0 < i; i-- {
		if maximal%i == 0 && minimal%i == 0 {
			return i
		}
	}
	return -1
}
