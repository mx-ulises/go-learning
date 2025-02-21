func getInt(num string) int {
	i := 0
	for _, c := range num {
		i <<= 1
		if c == '1' {
			i |= 1
		}
	}
	return i
}

func getString(i, n int) string {
	num := make([]byte, n)
	for j := n - 1; 0 <= j; j-- {
		switch i & 1 {
		case 0:
			num[j] = '0'
		case 1:
			num[j] = '1'
		}
		i >>= 1
	}
	return string(num)
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums[0])
	numSet := make([]bool, n+1)
	firstValid := 0
	for _, num := range nums {
		i := getInt(num)
		if i <= n {
			numSet[i] = true
		}
		for numSet[firstValid] {
			firstValid++
		}
	}
	return getString(firstValid, n)
}
