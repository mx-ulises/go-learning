func isSelfDividing(num int) bool {
	current := num
	for 0 < current {
		d := current % 10
		if d == 0 || (num%d) != 0 {
			return false
		}
		current /= 10
	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	output := []int{}
	for left <= right {
		if isSelfDividing(left) {
			output = append(output, left)
		}
		left++
	}
	return output
}
