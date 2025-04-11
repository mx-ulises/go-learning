func isSymmetric(x int) bool {
	bucket := 0
	switch {
	case x < 10:
		return false
	case 10 <= x && x < 100:
		bucket = 1
	case 100 <= x && x < 1000:
		return false
	case 1000 <= x && x < 10000:
		bucket = 2
	default:
		return false
	}
	left, right := 0, 0
	for i := 0; i < bucket; i++ {
		right += x % 10
		x /= 10
	}
	for i := 0; i < bucket; i++ {
		left += x % 10
		x /= 10
	}
	return right == left
}

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for x := low; x <= high; x++ {
		if isSymmetric(x) {
			count++
		}
	}
	return count
}
