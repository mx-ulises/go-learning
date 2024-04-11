func findMatrix(nums []int) [][]int {
	output := [][]int{}
	numPosition := map[int]int{}
	for _, num := range nums {
		position := numPosition[num]
		if position == len(output) {
			output = append(output, []int{})
		}
		output[position] = append(output[position], num)
		numPosition[num]++
	}
	return output
}
