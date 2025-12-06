package leetcode

func zigzagTraversal(grid [][]int) []int {
	output := []int{}
	for i, row := range grid {
		output = traverseRow(output, row, i)
	}
	return output
}

func traverseRow(output []int, row []int, i int) []int {
	if isOdd(i) {
		return traverseRowOdd(output, row)
	}
	return traverseRowEven(output, row)
}

func isOdd(i int) bool {
	return i&1 == 1
}

func traverseRowOdd(output []int, row []int) []int {
	for i := len(row) - 1; 0 <= i; i-- {
		output = addIfOdd(output, i, row[i])
	}
	return output
}

func traverseRowEven(output []int, row []int) []int {
	for i := 0; i < len(row); i++ {
		output = addIfEven(output, i, row[i])
	}
	return output
}

func addIfOdd(output []int, i int, num int) []int {
	if isOdd(i) {
		output = append(output, num)
	}
	return output
}

func addIfEven(output []int, i int, num int) []int {
	if !isOdd(i) {
		output = append(output, num)
	}
	return output
}
