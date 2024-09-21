func AddNumbers(current, n int, output *[]int) {
	if n < current {
		return
	}
	(*output) = append(*output, current)
	current *= 10
	for i := 0; i <= 9; i++ {
		AddNumbers(current+i, n, output)
	}
}

func lexicalOrder(n int) []int {
	output := []int{}
	for i := 1; i <= 9; i++ {
		AddNumbers(i, n, &output)
	}
	return output
}
