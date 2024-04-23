func sumOfSquares(nums []int) int {
	n, output := len(nums), 0
	for i, num := range nums {
		if n%(i+1) == 0 {
			output += (num * num)
		}
	}
	return output
}
