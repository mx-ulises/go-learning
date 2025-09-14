package leetcode

import "fmt"

func getDigits(s string, k int) []byte {
	digits := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = s[i]
	}
	return digits
}

func reverse(digits []byte) []byte {
	n := len(digits) / 2
	for i := 0; i < n; i++ {
		j := len(digits) - i - 1
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func addToArray(digits []byte, num int) []byte {
	if num == 0 {
		return []byte{'0'}
	}
	for 0 < num {
		digits = append(digits, byte(num%10)+'0')
		num /= 10
	}
	return digits
}

func digitSum(s string, k int) string {
	digits := getDigits(s, k)
	for k < len(digits) {
		i, n, num := 0, len(digits), 0
		buffer := []byte{}
		for i < n {
			num += int(digits[0] - '0')
			digits = digits[1:]
			i++
			if i%k == 0 {
				buffer = reverse(addToArray(buffer, num))
				digits = append(digits, buffer...)
				buffer = []byte{}
				num = 0
			}
		}
		if i%k != 0 {
			buffer = reverse(addToArray(buffer, num))
			digits = append(digits, buffer...)
		}
		fmt.Println(string(digits))
	}
	return string(digits)
}
