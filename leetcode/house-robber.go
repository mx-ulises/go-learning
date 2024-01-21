func rob(nums []int) int {
	window := [3]int{0, 0, 0}
	for _, x := range nums {
		candidate := window[0]
		if candidate < window[1] {
			candidate = window[1]
		}
		candidate += x
		window[0] = window[1]
		window[1] = window[2]
		window[2] = candidate
	}
	if window[1] < window[2] {
		return window[2]
	}
	return window[1]
}
