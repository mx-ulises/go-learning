package leetcode

func valid(num string) bool {
	return num[0] == num[1] && num[1] == num[2]
}

func getMax(num1, num2 string) string {
	if num1 == "" || num1[0] < num2[0] {
		return num2
	}
	return num1
}

func largestGoodInteger(num string) string {
	output := ""
	for i := 3; i <= len(num); i++ {
		candidate := num[i-3 : i]
		if valid(candidate) {
			output = getMax(output, candidate)
		}
	}
	return output
}
