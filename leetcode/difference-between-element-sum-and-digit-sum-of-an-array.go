func GetDigitSum(num int) int {
	digitSum := 0
	for 0 < num {
		digitSum += num % 10
		num /= 10
	}
	return digitSum
}

func differenceOfSum(nums []int) int {
	sum := 0
	digitSum := 0
	for _, num := range nums {
		sum += num
		digitSum += GetDigitSum(num)
	}
	return sum - digitSum
}
