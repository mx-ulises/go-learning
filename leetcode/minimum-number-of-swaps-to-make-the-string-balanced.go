func minSwaps(s string) int {
	leftBalance, rightBalance := 0, 0
	left, right := 0, len(s)-1
	swaps := 0
	for left < right {
		for left < right && 0 <= leftBalance {
			if s[left] == '[' {
				leftBalance++
			} else {
				leftBalance--
			}
			left++
		}
		for left <= right && 0 <= rightBalance {
			if s[right] == ']' {
				rightBalance++
			} else {
				rightBalance--
			}
			right--
		}
		if leftBalance < 0 {
			swaps++
			leftBalance = 1
			rightBalance = 1
		}
	}
	return swaps
}
