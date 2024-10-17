func GetDigitSum(num int) int {
	digitSum := 0
	for 0 < num {
		digitSum += num % 10
		num /= 10
	}
	return digitSum
}

func minElement(nums []int) int {
	minimal := 10000
	for _, num := range nums {
		minimal = min(minimal, GetDigitSum(num))
	}
	return minimal
}
