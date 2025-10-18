package leetcode

func pow10(k int) int {
	if k == 0 {
		return 1
	}
	result := pow10(k / 2)
	result *= result
	if k&1 == 1 {
		result *= 10
	}
	return result
}

func divisorSubstrings(num int, k int) int {
	original, divisor := num, pow10(k)
	total := 0
	for divisor < num {
		candidate := num % divisor
		if candidate != 0 && (original%candidate) == 0 {
			total++
		}
		num /= 10
	}
	if (original % num) == 0 {
		total++
	}
	return total
}
