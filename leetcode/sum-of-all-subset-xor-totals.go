func traverse(nums *[]int, i, n, value int, total *int) {
	if i == n {
		*total += value
	} else {
		j := i + 1
		traverse(nums, j, n, value, total)
		traverse(nums, j, n, value^(*nums)[i], total)
	}
}

func subsetXORSum(nums []int) int {
	total := 0
	traverse(&nums, 0, len(nums), 0, &total)
	return total
}
