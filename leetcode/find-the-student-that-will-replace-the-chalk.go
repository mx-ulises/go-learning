func Sum(chalk *[]int) int {
	total := 0
	for _, x := range *chalk {
		total += x
	}
	return total
}

func chalkReplacer(chalk []int, k int) int {
	k = k % Sum(&chalk)
	for i, x := range chalk {
		k -= x
		if k < 0 {
			return i
		}
	}
	return 0
}
