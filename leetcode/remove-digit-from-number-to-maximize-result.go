package leetcode

func removeDigit(number string, digit byte) string {
	result, j, removed, count := make([]byte, len(number)-1), 0, false, getInstances(number, digit)
	for i := 0; i < len(number)-1; i++ {
		if number[i] == digit {
			count--
		}
		if !removed && needToBeRemoved(number, digit, i, count) {
			removed = true
		} else {
			result[j] = number[i]
			j++
		}
	}
	if removed {
		result[j] = number[len(number)-1]
	}
	return string(result)
}

func needToBeRemoved(number string, digit byte, i int, count int) bool {
	if number[i] != digit {
		return false
	}
	if count == 0 {
		return true
	}
	if number[i] < number[i+1] {
		return true
	}
	return false
}

func getInstances(number string, digit byte) int {
	count := 0
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			count++
		}
	}
	return count
}
