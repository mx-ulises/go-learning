func BinarySearchSqrt(x int) int {
	start := 0
	end := x
	for start <= end {
		mid := (start + end) / 2
		if x < (mid * mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start - 1
}

func LinearSqrt(x int) int {
	i, j := 0, 1
	current, next := 0, 1
	for !(current <= x && x < next) {
		i = j
		current = next
		j++
		next = j * j
	}
	return i
}

func mySqrt(x int) int {
	return BinarySearchSqrt(x)
}
