func minSubarray(nums []int, p int) int {
	toSubstract := 0
	for _, num := range nums {
		toSubstract += num
	}
	toSubstract %= p
	n := len(nums)
	latestPosition := map[int]int{0: 0}
	cumulative, minimal := 0, n
	for i := 0; i < n; i++ {
		curPostion := i + 1
		cumulative = (cumulative + nums[i]) % p
		latestPosition[cumulative] = curPostion
		toSearch := cumulative - toSubstract
		if toSearch < 0 {
			toSearch += p
		}
		prevPosition, ok := latestPosition[toSearch]
		if ok {
			minimal = min(minimal, curPostion-prevPosition)
		}
	}
	if minimal == n {
		return -1
	}
	return minimal
}
