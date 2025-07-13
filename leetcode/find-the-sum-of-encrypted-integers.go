package leetcode

func encrypt(x int) int {
	maximal, digits, output := 0, 0, 0
	for 0 < x {
		maximal = max(maximal, x%10)
		digits++
		x /= 10
	}
	for 0 < digits {
		output = output*10 + maximal
		digits--
	}
	return output
}

func sumOfEncryptedInt(nums []int) int {
	output := 0
	for _, num := range nums {
		output += encrypt(num)
	}
	return output
}
