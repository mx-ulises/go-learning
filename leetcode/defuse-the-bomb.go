func increaseCircular(i, n int) int {
	i++
	if i == n {
		i = 0
	}
	return i
}

func getStartEnd(k, n int) (int, int) {
	start, end := 0, 0
	if k < 0 {
		start = k + n
	}
	if 0 < k {
		start = 1
		end = increaseCircular(k, n)
	}
	return start, end
}

func getCurrent(code []int, start, end int) int {
	n := len(code)
	current := 0
	i := start
	for i != end {
		current += code[i]
		i = increaseCircular(i, n)
	}
	return current
}

func decrypt(code []int, k int) []int {
	n := len(code)
	start, end := getStartEnd(k, n)
	current := getCurrent(code, start, end)
	output := make([]int, n)
	for i := 0; i < n; i++ {
		output[i] = current
		current += code[end] - code[start]
		end = increaseCircular(end, n)
		start = increaseCircular(start, n)
	}
	return output
}
