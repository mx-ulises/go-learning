func generateKey(num1 int, num2 int, num3 int) int {
	nums := []int{num1, num2, num3}
	key := 0
	for multiplier := 1; multiplier < 10000; multiplier *= 10 {
		minimalDigit := 9
		for i := 0; i < len(nums); i++ {
			minimalDigit = min(minimalDigit, nums[i]%10)
			nums[i] /= 10
		}
		key += minimalDigit * multiplier
	}
	return key
}
