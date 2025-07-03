package leetcode

func getDigits(num int) []int {
	digits := []int{}
	for 0 < num {
		digits = append(digits, num%10)
		num /= 10
	}
	return digits
}

func getReplace(digits []int) int {
	for i := len(digits) - 1; 0 <= i; i-- {
		if digits[i] != 9 {
			return digits[i]
		}
	}
	return -1
}

func replaceDigit(digits []int, old, new int) []int {
	for i := len(digits) - 1; 0 <= i; i-- {
		if digits[i] == old {
			digits[i] = new
		}
	}
	return digits
}

func getNumber(digits []int) int {
	num := 0
	for i := len(digits) - 1; 0 <= i; i-- {
		num = num*10 + digits[i]
	}
	return num
}

func getMaximal(num int) int {
	digits := getDigits(num)
	old := getReplace(digits)
	return getNumber(replaceDigit(digits, old, 9))
}

func getMinimal(num int) int {
	digits := getDigits(num)
	old := digits[len(digits)-1]
	return getNumber(replaceDigit(digits, old, 0))

}

func minMaxDifference(num int) int {
	maximal := getMaximal(num)
	minimal := getMinimal(num)
	return maximal - minimal
}
